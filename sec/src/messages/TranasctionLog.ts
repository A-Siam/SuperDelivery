export default interface TransactionLog {
    eventName: string
    rollbackEventName: string
    order: number
    tranasctionId: string
}
