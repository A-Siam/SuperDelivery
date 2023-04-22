# Saga Execution Controller for Super delivery
This is a simplified SEC implmenetation in nodejs for Super Delivery 
## Build and run
### Env
```env
KAFKA_BROKER="localhost:9002"
EVENTS_TOPIC_NAME="transactional-events-topic"
CLIENT_ID="sec"
GROUP_ID="transactional-events-group"

# events db url
DATABASE_URL="postgresql://admin:admin@localhost:5432/transaction_events?schema=public"
```
