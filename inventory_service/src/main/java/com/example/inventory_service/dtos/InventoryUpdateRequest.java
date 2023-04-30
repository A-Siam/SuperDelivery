package com.example.inventory_service.dtos;

import com.fasterxml.jackson.annotation.JsonProperty;

public record InventoryUpdateRequest(
        @JsonProperty("inventory_type") String type,
        @JsonProperty("request_type") RequestType requestType,
        @JsonProperty("amount") Double amount) {
}
