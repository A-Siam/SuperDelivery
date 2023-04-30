package com.example.inventory_service.dtos;

import com.fasterxml.jackson.annotation.JsonProperty;

public record InventoryRevertRequest(
        @JsonProperty("event_name") String eventName,
        @JsonProperty("service") String service
) {
}
