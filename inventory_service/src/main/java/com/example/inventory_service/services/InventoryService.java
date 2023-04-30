package com.example.inventory_service.services;

import java.util.List;

import com.example.inventory_service.data.entities.Inventory;
import com.example.inventory_service.data.repositories.InventoryRepository;
import com.example.inventory_service.dtos.Event;
import com.example.inventory_service.dtos.InventoryContentDto;

import com.example.inventory_service.dtos.InventoryUpdateRequest;
import com.example.inventory_service.dtos.RequestType;
import com.example.inventory_service.utils.EventEmitter;
import com.fasterxml.jackson.core.JsonProcessingException;
import lombok.RequiredArgsConstructor;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

@Service
public record InventoryService (
		InventoryRepository inventoryRepository,
		EventEmitter eventEmitter,
		@Value("${events.topic-name}") String eventsTopicName,
		@Value("${events.service-name}") String serviceName
) {

	public List<InventoryContentDto> getInventoryContent() {
	   final var inventory = inventoryRepository.findAll();
	   return inventory
			   .stream()
			   .map(this::toDto)
			   .toList();
	}

	@Transactional
	public void updateInventory(InventoryUpdateRequest updateRequest) throws JsonProcessingException {
		final var inventory = inventoryRepository.findByType(updateRequest.type()).orElseThrow();
		final var oldNumberOfItems = inventory.getNumberOfItems();
		if (updateRequest.requestType() == RequestType.INCREASE) {
			inventory.setNumberOfItems((long) (oldNumberOfItems + updateRequest.amount()));
		} else {
			inventory.setNumberOfItems((long) (oldNumberOfItems - updateRequest.amount()));
		}
		inventoryRepository.save(inventory);
		final var event = new Event(updateRequest.requestType().toString() + "_EVENT", serviceName);
		eventEmitter.emit(eventsTopicName, event);
	}

	private InventoryContentDto toDto(Inventory inventory) {
		return InventoryContentDto
				.builder()
				.numberOfItems(inventory.getNumberOfItems())
				.inventoryName(inventory.getInventoryName())
				.build();
	}
}
