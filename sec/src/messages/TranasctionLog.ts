export default interface TransactionLog {
    eventName: string
    previousEventName: string
    order: number
    tranasctionId: string
    serviceName: string
}
