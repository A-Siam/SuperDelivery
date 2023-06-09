import { GROUP_ID } from '../constants'
import TransactionLog from '../messages/TranasctionLog'
import { getKafkaClient } from './kafka-connection'

const consumer = getKafkaClient().consumer({ groupId: GROUP_ID! })

type OnMessageHandler = (message: TransactionLog) => Promise<void>

async function initConsumer(onMessage: OnMessageHandler) {
    await consumer.run({
        eachMessage: async ({ message }) => {
            onMessage(JSON.parse(message.value?.toString()!))
        },
    })
}

export async function initiateListener(
    topic: string,
    onMessageRecived: OnMessageHandler
) {
    await consumer.connect()
    await consumer.subscribe({ topic, fromBeginning: true })
    initConsumer(onMessageRecived)
}
