package com.example.inventory_service.data.entities;

import java.util.List;

import jakarta.persistence.*;
import lombok.Getter;
import lombok.Setter;

@Entity
@Getter
@Setter
@Table(
        indexes = {
                @Index(name = "inventory_type_idx", columnList = "type")
        }
)
public class Inventory {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    private String type;
    private String inventoryName;
    private Long numberOfItems;
}
