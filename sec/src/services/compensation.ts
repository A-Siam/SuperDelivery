import axios from 'axios'
import services from '../../services_map.json'
import { SAGA_ROLLBACK_API } from '../constants'

export default class TransactionCompensationService {
    async compoensate(
        serviceName: string,
        transactionId: string,
        eventName: string
    ) {
        const serviceDate = services[serviceName]
        try {
            await axios.post(serviceDate + SAGA_ROLLBACK_API, {
                tranasaction_id: transactionId,
                event_name: eventName,
            })
        } catch (e) {

        }
    }
}
