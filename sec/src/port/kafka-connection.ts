import { Kafka } from 'kafkajs'
import { CLIENT_ID } from '../constants'

let _kafka: Kafka | undefined = undefined

export function getKafkaClient(): Kafka {
    if (!_kafka) {
        _kafka = new Kafka({
            clientId: CLIENT_ID,
            brokers: [process.env['KAFKA_BROKER']!],
        })
    }
    return _kafka
}
