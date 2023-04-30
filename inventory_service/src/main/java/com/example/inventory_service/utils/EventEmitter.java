package com.example.inventory_service.utils;

import com.example.inventory_service.dtos.Event;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import lombok.RequiredArgsConstructor;
import org.springframework.kafka.core.KafkaTemplate;
import org.springframework.stereotype.Component;

@Component
public record EventEmitter(
        KafkaTemplate<String, String> template,
        ObjectMapper objectMapper
) {
    public void emit(String topic, Event event) throws JsonProcessingException {
        this.template.send(topic, objectMapper.writeValueAsString(event));
    }
}
