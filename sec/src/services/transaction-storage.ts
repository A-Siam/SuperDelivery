import { Prisma, PrismaClient, ActionType } from '@prisma/client'

import TransactionLog from '../messages/TranasctionLog'
import { Rollback } from '../dto/rollback'
import TransactionCompensationService from './compensation'

export default class TransactionStorageService {
    constructor(private compensationService: TransactionCompensationService) {}

    async rollback(message: Rollback) {
        const prisma = new PrismaClient()
        await prisma.$transaction(async (txn) => {
            const transactionLogRepo = txn.transactionLog
            const currentTxnLog = await transactionLogRepo.findFirst({
                where: {
                    transactionId: message.tranasactionId,
                    currentEventName: message.eventName,
                },
            })
            if (currentTxnLog == null) {
                throw new Error('Invalid tranasction id or event name')
            }
            const previousTransactions = await transactionLogRepo.findMany({
                orderBy: {
                    order: 'desc',
                },
                where: {
                    order: {
                        lt: currentTxnLog.order,
                    },
                },
            })
            previousTransactions.forEach((pTxn) => {
                this.compensationService.compoensate(
                    pTxn.serviceName,
                    pTxn.transactionId,
                    pTxn.currentEventName
                )
            })
        })
    }

    async storeTranasaction(message: TransactionLog) {
        const prisma = new PrismaClient()
        await prisma.$transaction(
            async (txn) => {
                const tranasctionLogRepo = txn.transactionLog
                const previousTransaction = await tranasctionLogRepo.findFirst({
                    orderBy: {
                        order: 'desc',
                    },
                    where: {
                        transactionId: message.tranasctionId,
                    },
                })
                const previousOrder =
                    previousTransaction != null ? previousTransaction.order : 0

                tranasctionLogRepo.create({
                    data: {
                        transactionId: message.tranasctionId,
                        currentEventName: message.eventName,
                        previousEventName:
                            previousTransaction?.currentEventName,
                        order: previousOrder + 1,
                        actionType: ActionType.TRANSACTION,
                        serviceName: message.serviceName,
                    },
                })
            },
            {
                isolationLevel: Prisma.TransactionIsolationLevel.Serializable,
            }
        )
    }
}
