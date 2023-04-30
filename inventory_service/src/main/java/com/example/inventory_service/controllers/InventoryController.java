package com.example.inventory_service.controllers;

import java.util.List;

import com.example.inventory_service.dtos.InventoryContentDto;

import com.example.inventory_service.dtos.InventoryRevertRequest;
import com.example.inventory_service.dtos.InventoryUpdateRequest;
import org.apache.coyote.Response;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/inventory")
public class InventoryController {
    @GetMapping
    ResponseEntity<List<InventoryContentDto>> getInventoryContent() {
        final var inventoryContent = InventoryContentDto.builder().numberOfItems(456L).build();
        return null;
    }

    @PostMapping
    ResponseEntity<Void> updateInventory(InventoryUpdateRequest request) {
        return null;
    }

    @PostMapping("/revert")
    ResponseEntity<Void> revert(InventoryRevertRequest request) {
        return null;
    }

}
