package com.example.inventory_service.config;

import org.apache.kafka.clients.admin.NewTopic;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.kafka.config.TopicBuilder;

@Configuration
public class EventsConfig {
    // no need to define Kafka admin as I'm using the default settings for the broker
    // otherwise it will be necessary to define an admin bean
    @Bean
    public NewTopic topic(
            @Value("${events.topic-name}") String eventsTopicName
    ) {
        return TopicBuilder.name(eventsTopicName)
                .partitions(10)
                .replicas(1)
                .build();
    }
}
