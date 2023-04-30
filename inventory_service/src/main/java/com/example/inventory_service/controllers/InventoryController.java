package com.example.inventory_service.controllers;

import java.util.List;

import com.example.inventory_service.dtos.InventoryContentDto;

import com.example.inventory_service.dtos.InventoryRevertRequest;
import com.example.inventory_service.dtos.InventoryUpdateRequest;
import com.example.inventory_service.services.InventoryService;
import com.fasterxml.jackson.core.JsonProcessingException;
import org.apache.coyote.Response;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/inventory")
public record InventoryController(
        InventoryService service
) {

    @GetMapping
    ResponseEntity<List<InventoryContentDto>> getInventoryContent() {
        final var inventoryContent = service.getInventoryContent();
        return ResponseEntity.ok(inventoryContent);
    }

    @PostMapping
    ResponseEntity<InventoryContentDto> updateInventory(InventoryUpdateRequest request) throws JsonProcessingException {
        final var response = service.updateInventory(request);
        return ResponseEntity.ok(response);
    }

    @PostMapping("/revert")
    ResponseEntity<InventoryContentDto> revert(InventoryRevertRequest request) {
        final var response = service.revert(request);
        return ResponseEntity.ok(response);
    }

}
