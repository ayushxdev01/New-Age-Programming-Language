package main

import "testing"

func TestAddItem(t *testing.T) {
	inventory := &Inventory{}

	err := inventory.AddItem(Item{
		ID:       1,
		Name:     "Keyboard",
		Quantity: 5,
		Price:    750,
	})
	if err != nil {
		t.Fatalf("expected item to be added, got error: %v", err)
	}

	items := inventory.GetItems()
	if len(items) != 1 {
		t.Fatalf("expected 1 item, got %d", len(items))
	}
	if items[0].Name != "Keyboard" {
		t.Errorf("expected Keyboard, got %s", items[0].Name)
	}
}

func TestAddDuplicateItem(t *testing.T) {
	inventory := &Inventory{}
	item := Item{ID: 1, Name: "Mouse", Quantity: 3, Price: 400}

	if err := inventory.AddItem(item); err != nil {
		t.Fatalf("unexpected error adding first item: %v", err)
	}
	if err := inventory.AddItem(item); err == nil {
		t.Fatal("expected an error when adding a duplicate ID")
	}
}

func TestUpdateItem(t *testing.T) {
	inventory := &Inventory{}
	_ = inventory.AddItem(Item{ID: 2, Name: "Monitor", Quantity: 2, Price: 8000})

	err := inventory.UpdateItem(2, "LED Monitor", 4, 9000)
	if err != nil {
		t.Fatalf("expected item to be updated, got error: %v", err)
	}

	item := inventory.GetItems()[0]
	if item.Name != "LED Monitor" || item.Quantity != 4 || item.Price != 9000 {
		t.Errorf("updated item has unexpected values: %+v", item)
	}
}

func TestRemoveItem(t *testing.T) {
	inventory := &Inventory{}
	_ = inventory.AddItem(Item{ID: 3, Name: "USB Cable", Quantity: 10, Price: 100})

	if err := inventory.RemoveItem(3); err != nil {
		t.Fatalf("expected item to be removed, got error: %v", err)
	}
	if len(inventory.GetItems()) != 0 {
		t.Fatal("expected inventory to be empty after removal")
	}
}

func TestSummaryAsync(t *testing.T) {
	inventory := &Inventory{}
	_ = inventory.AddItem(Item{ID: 1, Name: "Pen", Quantity: 10, Price: 5})
	_ = inventory.AddItem(Item{ID: 2, Name: "Notebook", Quantity: 4, Price: 40})

	summary := <-inventory.SummaryAsync()

	if summary.ItemCount != 2 {
		t.Errorf("expected 2 different items, got %d", summary.ItemCount)
	}
	if summary.TotalQuantity != 14 {
		t.Errorf("expected total quantity 14, got %d", summary.TotalQuantity)
	}
	if summary.TotalValue != 210 {
		t.Errorf("expected total value 210, got %.2f", summary.TotalValue)
	}
}

func TestRejectInvalidItem(t *testing.T) {
	inventory := &Inventory{}
	err := inventory.AddItem(Item{ID: 0, Name: "Invalid", Quantity: 1, Price: 10})
	if err == nil {
		t.Fatal("expected validation error for ID zero")
	}
}
