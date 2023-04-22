import { Prisma, PrismaClient } from '@prisma/client'
import TransactionLog from '../messages/TranasctionLog'

export async function storeTranasaction(message: TransactionLog) {
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
                    previousEventName: message.previousEventName,
                    order: previousOrder + 1,
                },
            })
        },
        {
            isolationLevel: Prisma.TransactionIsolationLevel.Serializable,
        }
    )
}
