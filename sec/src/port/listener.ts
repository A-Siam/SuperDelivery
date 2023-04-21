import TransactionLog from "../data/TranasctionLog";
import { getKafkaClient } from "./kafka-connection";

const consumer = getKafkaClient().consumer({groupId: GROUP_ID})

function initConsumer() {
    
}

export async function initiateListener(topic: string, onMessageRecived: (message: TransactionLog) => void) {
     
}
