package com.example.inventory_service.dtos;

import java.util.List;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.Builder;
import lombok.Value;

@Builder
public record InventoryContentDto(
        @JsonProperty("id")
        Long id,
        @JsonProperty("inventory_name")
        String inventoryName,
        @JsonProperty("number_of_items")
        Long numberOfItems
) {
}

