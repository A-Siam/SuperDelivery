import { PrismaClient } from '@prisma/client'
import { initiateListener } from './port/listener'
import { storeTranasaction } from './services/transaction-storage'

async function main() {
    await initiateListener(process.env['EVENTS_TOPIC_NAME']!, storeTranasaction)
}

main()
    .catch((e) => {
        console.error(e)
    })
    .finally(() => {
        const prisma = new PrismaClient()
        prisma.$disconnect()
    })
