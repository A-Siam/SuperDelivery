import { Rollback } from '../dto/rollback'

export function toRollbackRequestBody(input: any): Rollback {
    return {
        eventName: input['event_name'],
        tranasactionId: input['tranasaction_id'],
    }
}
