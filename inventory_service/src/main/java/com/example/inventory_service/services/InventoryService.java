package com.example.inventory_service.services;

import java.util.List;

import com.example.inventory_service.data.entities.Inventory;
import com.example.inventory_service.data.repositories.InventoryRepository;
import com.example.inventory_service.dtos.*;

import com.example.inventory_service.utils.EventEmitter;
import com.fasterxml.jackson.core.JsonProcessingException;
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
	public InventoryContentDto updateInventory(InventoryUpdateRequest updateRequest) throws JsonProcessingException {
		final var inventory = inventoryRepository.findByType(updateRequest.type()).orElseThrow();
		final var oldNumberOfItems = inventory.getNumberOfItems();
		if (updateRequest.requestType() == RequestType.INCREASE) {
			inventory.setNumberOfItems((long) (oldNumberOfItems + updateRequest.amount()));
		} else {
			inventory.setNumberOfItems((long) (oldNumberOfItems - updateRequest.amount()));
		}
		final var result = inventoryRepository.save(inventory);
		final var event = new Event(
				updateRequest.requestType().toString() + "_EVENT_BY_" + updateRequest.amount() + "_ON_" + inventory.getId(),
				serviceName
		);
		eventEmitter.emit(eventsTopicName, event);
		return this.toDto(result);
	}

	private InventoryContentDto toDto(Inventory inventory) {
		return InventoryContentDto
				.builder()
				.numberOfItems(inventory.getNumberOfItems())
				.inventoryName(inventory.getInventoryName())
				.build();
	}

	public InventoryContentDto revert(InventoryRevertRequest revertRequest) {
		if (!revertRequest.service().equals(serviceName)) {
			return null;
		}
		final var revertOperationToken = getInventoryTokens(revertRequest);
		final var inventory = inventoryRepository.findById(revertOperationToken.inventoryId()).orElseThrow();
		final var oldNumberOfItems = inventory.getNumberOfItems();
		if (revertOperationToken.operationType == RequestType.DECREASE) {
			inventory.setNumberOfItems((long) (oldNumberOfItems + revertOperationToken.amount()));
		} else {
			inventory.setNumberOfItems((long) (oldNumberOfItems - revertOperationToken.amount()));
		}
		final var result = inventoryRepository.save(inventory);
		return this.toDto(result);
	}


	private InventoryRevertOperationToken getInventoryTokens(InventoryRevertRequest revertRequest) {
		String[] eventParts = revertRequest.eventName().split("_EVENT_BY_");
		RequestType operationType = RequestType.valueOf(eventParts[0]);
		String eventDescription = eventParts[1];
		String[] eventTokens = eventDescription.split("_");
		Double amount = Double.parseDouble(eventTokens[0]);
		Long inventoryId = Long.parseLong(eventTokens[2]);
		return new InventoryRevertOperationToken(
				operationType,
				amount,
				inventoryId
		);
	}

	record InventoryRevertOperationToken(
		RequestType operationType,
		Double amount,
		Long inventoryId
	) { }
}
