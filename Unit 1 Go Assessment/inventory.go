package main

import (
	"errors"
	"strings"
)

// Item represents one product in the inventory.
type Item struct {
	ID       int
	Name     string
	Quantity int
	Price    float64
}

// Inventory stores items using a slice.
type Inventory struct {
	items []Item
}

// Summary contains the inventory statistics.
type Summary struct {
	ItemCount     int
	TotalQuantity int
	TotalValue    float64
}

// AddItem validates and adds a new item to the inventory.
func (inv *Inventory) AddItem(item Item) error {
	if item.ID <= 0 {
		return errors.New("item ID must be greater than zero")
	}
	if strings.TrimSpace(item.Name) == "" {
		return errors.New("item name cannot be empty")
	}
	if item.Quantity <= 0 {
		return errors.New("quantity must be greater than zero")
	}
	if item.Price < 0 {
		return errors.New("price cannot be negative")
	}

	for _, existing := range inv.items {
		if existing.ID == item.ID {
			return errors.New("an item with this ID already exists")
		}
	}

	item.Name = strings.TrimSpace(item.Name)
	inv.items = append(inv.items, item)
	return nil
}

// UpdateItem changes the name, quantity, and price of an existing item.
func (inv *Inventory) UpdateItem(id int, name string, quantity int, price float64) error {
	if id <= 0 {
		return errors.New("item ID must be greater than zero")
	}
	if strings.TrimSpace(name) == "" {
		return errors.New("item name cannot be empty")
	}
	if quantity <= 0 {
		return errors.New("quantity must be greater than zero")
	}
	if price < 0 {
		return errors.New("price cannot be negative")
	}

	for index := range inv.items {
		if inv.items[index].ID == id {
			inv.items[index].Name = strings.TrimSpace(name)
			inv.items[index].Quantity = quantity
			inv.items[index].Price = price
			return nil
		}
	}
	return errors.New("item not found")
}

// RemoveItem deletes an item using its ID.
func (inv *Inventory) RemoveItem(id int) error {
	if id <= 0 {
		return errors.New("item ID must be greater than zero")
	}

	for index, item := range inv.items {
		if item.ID == id {
			inv.items = append(inv.items[:index], inv.items[index+1:]...)
			return nil
		}
	}
	return errors.New("item not found")
}

// GetItems returns a copy so callers cannot directly change the inventory slice.
func (inv *Inventory) GetItems() []Item {
	itemsCopy := make([]Item, len(inv.items))
	copy(itemsCopy, inv.items)
	return itemsCopy
}

// SummaryAsync calculates statistics in a separate goroutine and sends the result
// through a channel.
func (inv *Inventory) SummaryAsync() <-chan Summary {
	resultChannel := make(chan Summary)
	itemsSnapshot := inv.GetItems()

	go func() {
		defer close(resultChannel)

		var result Summary
		result.ItemCount = len(itemsSnapshot)

		for _, item := range itemsSnapshot {
			result.TotalQuantity += item.Quantity
			result.TotalValue += float64(item.Quantity) * item.Price
		}

		resultChannel <- result
	}()

	return resultChannel
}
