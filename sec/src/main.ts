import { PrismaClient } from '@prisma/client'
import { initiateListener } from './port/listener'
import TransactionStorageService from './services/transaction-storage'
import TransactionCompensationService from './services/compensation'

async function main() {
    const transactionCompensationService = new TransactionCompensationService()
    const transactionStorageService = new TransactionStorageService(
        transactionCompensationService
    )
    await initiateListener(
        process.env['EVENTS_TOPIC_NAME']!,
        transactionStorageService.storeTranasaction
    )
}

main()
    .catch((e) => {
        console.error(e)
    })
    .finally(() => {
        const prisma = new PrismaClient()
        prisma.$disconnect()
    })
