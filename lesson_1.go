package main

import "fmt"

// updateQuantity updates the quantity of an item and returns the new quantity
// and whether the item is out of stock (quantity = 0)
func updateQuantity(quantity *int, change int) (int, bool) {
	// TODO: Update the quantity pointer and return the new value and out-of-stock status
	*quantity += change
	if *quantity < 0 {
		*quantity = 0
	}
	return *quantity, *quantity == 0
}

// calculateValue calculates the total value of an item based on quantity and price
// It also updates the mostValuableItem if this item is more valuable
func calculateValue(itemName string, quantity int, price float64, mostValuableItem *string, highestValue *float64) float64 {
	// TODO: Calculate the total value, update mostValuableItem if needed, and return the total value
	totalValue := float64(quantity) * price
	if totalValue > *highestValue {
		*highestValue = totalValue
		*mostValuableItem = itemName
	}
	return totalValue
}

// displayInventory prints the inventory details
// It takes pointers to allow the function to show real-time data
func displayInventory(apples, oranges, bananas *int, applePrice, orangePrice, bananaPrice float64) {
	fmt.Printf("Apples: %d (Value: $%.2f)\n", *apples, float64(*apples)*applePrice)
	fmt.Printf("Oranges: %d (Value: $%.2f)\n", *oranges, float64(*oranges)*orangePrice)
	fmt.Printf("Bananas: %d (Value: $%.2f)\n", *bananas, float64(*bananas)*bananaPrice)
}

func main() {
	// Initialize inventory
	apples := 10
	oranges := 15
	bananas := 8

	// Prices
	applePrice := 0.5  // $0.50 each
	orangePrice := 0.7 // $0.70 each
	bananaPrice := 0.3 // $0.30 each

	// Track the most valuable item
	var mostValuableItem string
	var highestValue float64

	// Display initial inventory
	fmt.Println("Initial Inventory:")
	displayInventory(&apples, &oranges, &bananas, applePrice, orangePrice, bananaPrice)

	// Calculate initial values and find most valuable item
	appleValue := calculateValue("Apples", apples, applePrice, &mostValuableItem, &highestValue)
	orangeValue := calculateValue("Oranges", oranges, orangePrice, &mostValuableItem, &highestValue)
	bananaValue := calculateValue("Bananas", bananas, bananaPrice, &mostValuableItem, &highestValue)

	fmt.Printf("Total inventory value: $%.2f\n", appleValue+orangeValue+bananaValue)
	fmt.Printf("Most valuable item: %s\n\n", mostValuableItem)

	// Simulate some sales
	fmt.Println("Processing sales...")
	apples, applesOutOfStock := updateQuantity(&apples, -4)     // Sell 4 apples
	oranges, orangesOutOfStock := updateQuantity(&oranges, -8)  // Sell 8 oranges
	bananas, bananasOutOfStock := updateQuantity(&bananas, -10) // Try to sell 10 bananas (more than we have)

	// Check if any items are out of stock
	if applesOutOfStock {
		fmt.Println("Apples are out of stock!")
	}
	if orangesOutOfStock {
		fmt.Println("Oranges are out of stock!")
	}
	if bananasOutOfStock {
		fmt.Println("Bananas are out of stock!")
	}

	// Display updated inventory
	fmt.Println("\nUpdated Inventory:")
	displayInventory(&apples, &oranges, &bananas, applePrice, orangePrice, bananaPrice)

	// Reset most valuable tracking for recalculation
	mostValuableItem = ""
	highestValue = 0

	// Recalculate values
	appleValue = calculateValue("Apples", apples, applePrice, &mostValuableItem, &highestValue)
	orangeValue = calculateValue("Oranges", oranges, orangePrice, &mostValuableItem, &highestValue)
	bananaValue = calculateValue("Bananas", bananas, bananaPrice, &mostValuableItem, &highestValue)

	fmt.Printf("Total inventory value: $%.2f\n", appleValue+orangeValue+bananaValue)
	fmt.Printf("Most valuable item: %s\n\n", mostValuableItem)

	// Restock items
	fmt.Println("Restocking...")
	updateQuantity(&apples, 5)   // Add 5 apples
	updateQuantity(&oranges, 10) // Add 10 oranges
	updateQuantity(&bananas, 12) // Add 12 bananas

	// Display final inventory
	fmt.Println("\nFinal Inventory:")
	displayInventory(&apples, &oranges, &bananas, applePrice, orangePrice, bananaPrice)

	// Reset most valuable tracking for final calculation
	mostValuableItem = ""
	highestValue = 0

	// Final value calculation
	appleValue = calculateValue("Apples", apples, applePrice, &mostValuableItem, &highestValue)
	orangeValue = calculateValue("Oranges", oranges, orangePrice, &mostValuableItem, &highestValue)
	bananaValue = calculateValue("Bananas", bananas, bananaPrice, &mostValuableItem, &highestValue)

	fmt.Printf("Total inventory value: $%.2f\n", appleValue+orangeValue+bananaValue)
	fmt.Printf("Most valuable item: %s\n", mostValuableItem)
}
