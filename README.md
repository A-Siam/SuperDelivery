# Super delivery application
super delivery provide creating order -> update inventory -> process payment -> deliver order

## Patterns Used
- Saga pattern
- Event sourcing
- CQRS

**Note** This project *intentionally* doesn't use any external frameworks to deeply learn about the topics mentioned however for production grade application I would use something like **AxonIQ** framework to handle these patterns for me. Also auth is neglected in this application for simplicity

## Architecture
![super_delivery_arch drawio](https://user-images.githubusercontent.com/34966791/235349102-91a86b8b-c8e0-473a-996c-29fc60ca11f4.png)

## System components
### Languages 
- Java
- NodeJs
- GoLang

### Datastores
- PostgreSQL

### Event Streaming
- Kafka

## Building
