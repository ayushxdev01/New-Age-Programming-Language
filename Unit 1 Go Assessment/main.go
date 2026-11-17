package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var input = bufio.NewReader(os.Stdin)

func readLine(prompt string) string {
	fmt.Print(prompt)
	value, _ := input.ReadString('\n')
	return strings.TrimSpace(value)
}

func readInt(prompt string) (int, error) {
	value := readLine(prompt)
	number, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("please enter a valid whole number")
	}
	return number, nil
}

func readFloat(prompt string) (float64, error) {
	value := readLine(prompt)
	number, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, fmt.Errorf("please enter a valid price")
	}
	return number, nil
}

func addItem(inventory *Inventory) {
	fmt.Println("\n--- Add Item ---")
	id, err := readInt("Enter item ID: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	name := readLine("Enter item name: ")
	quantity, err := readInt("Enter quantity: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	price, err := readFloat("Enter price: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = inventory.AddItem(Item{
		ID:       id,
		Name:     name,
		Quantity: quantity,
		Price:    price,
	})
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Item added successfully.")
}

func updateItem(inventory *Inventory) {
	fmt.Println("\n--- Update Item ---")
	id, err := readInt("Enter the ID of the item to update: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	name := readLine("Enter new item name: ")
	quantity, err := readInt("Enter new quantity: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	price, err := readFloat("Enter new price: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if err := inventory.UpdateItem(id, name, quantity, price); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Item updated successfully.")
}

func removeItem(inventory *Inventory) {
	fmt.Println("\n--- Remove Item ---")
	id, err := readInt("Enter the ID of the item to remove: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if err := inventory.RemoveItem(id); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Item removed successfully.")
}

func displayItems(inventory *Inventory) {
	fmt.Println("\n--- Inventory Items ---")
	items := inventory.GetItems()
	if len(items) == 0 {
		fmt.Println("Inventory is empty.")
		return
	}

	fmt.Printf("%-8s %-20s %-12s %-12s\n", "ID", "Name", "Quantity", "Price")
	for _, item := range items {
		fmt.Printf("%-8d %-20s %-12d ₹%-11.2f\n",
			item.ID, item.Name, item.Quantity, item.Price)
	}
}

func displaySummary(inventory *Inventory) {
	fmt.Println("\n--- Inventory Summary (Concurrent) ---")
	// The summary is calculated by a goroutine and received through a channel.
	result := <-inventory.SummaryAsync()

	fmt.Printf("Number of different items: %d\n", result.ItemCount)
	fmt.Printf("Total quantity: %d\n", result.TotalQuantity)
	fmt.Printf("Total inventory value: ₹%.2f\n", result.TotalValue)
}

func main() {
	inventory := &Inventory{}

	fmt.Println("======================================")
	fmt.Println("       GO INVENTORY MANAGEMENT")
	fmt.Println("======================================")

	for {
		fmt.Println("\n1. Add Item")
		fmt.Println("2. Update Item")
		fmt.Println("3. Remove Item")
		fmt.Println("4. Display Inventory")
		fmt.Println("5. Calculate Summary")
		fmt.Println("6. Exit")

		choice, err := readInt("Enter your choice: ")
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		switch choice {
		case 1:
			addItem(inventory)
		case 2:
			updateItem(inventory)
		case 3:
			removeItem(inventory)
		case 4:
			displayItems(inventory)
		case 5:
			displaySummary(inventory)
		case 6:
			fmt.Println("Thank you for using Inventory Management.")
			return
		default:
			fmt.Println("Invalid choice. Please select 1 to 6.")
		}
	}
}
