// Homework 2: Object Oriented Programming
// Due February 7, 2017 at 11:59pm
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Feel free to use the main function for testing your functions
	world := struct {
		English string
		Spanish string
		French  string
	}{
		"world",
		"mundo",
		"monde",
	}
	fmt.Printf("Hello, %s/%s/%s!\n", world.English, world.Spanish, world.French)

	fmt.Println("---My Print Statements---")

	var p Price
	fmt.Println(p)

	RegisterItem(Prices, "apple", 50)
	RegisterItem(Prices, "apple", 70)

	var items []string
	items = append(items, "eggs")
	items = append(items, "eggs")
	items = append(items, "milk")
	cart := Cart{items, 2*Prices["eggs"] + Prices["milk"]}
	fmt.Println(cart.hasMilk())
	fmt.Println(cart.HasItem("eggs"))
	fmt.Println(cart.HasItem("bread"))

	fmt.Println(fmt.Sprintf("%v", cart.Items))
	fmt.Println(fmt.Sprintf("%d", cart.TotalPrice))
	cart.AddItem("banana")
	cart.AddItem("apple")
	fmt.Println(fmt.Sprintf("%v", cart.Items))
	fmt.Println(fmt.Sprintf("%d", cart.TotalPrice))

	cart.Checkout()
	cart.Checkout()
	fmt.Println(fmt.Sprintf("%v", cart.Items))
	fmt.Println(fmt.Sprintf("%d", cart.TotalPrice))
	cart.AddItem("chocolate")
	cart.AddItem("chocolate")
	cart.AddItem("chocolate")
	fmt.Println(fmt.Sprintf("%v", cart.Items))
	fmt.Println(fmt.Sprintf("%d", cart.TotalPrice))
}

// Price is the cost of something in US cents.
type Price int64

// String is the string representation of a Price
// These should be represented in US Dollars
// Example: 2595 cents => $25.95
func (p Price) String() string {
	return fmt.Sprintf("$%d.%02d", p/100, p%100)
}

// Prices is a map from an item to its price.
var Prices = map[string]Price{
	"eggs":          219,
	"bread":         199,
	"milk":          295,
	"peanut butter": 445,
	"chocolate":     150,
}

// Warning is a logger for warnings
var Warning = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)

// RegisterItem adds the new item in the prices map.
// If the item is already in the prices map, a warning should be displayed to the user,
// but the value should be overwritten.
// Bonus (1pt) - Use the "log" package to print the error to the user
func RegisterItem(prices map[string]Price, item string, price Price) {
	if oldPrice, ok := prices[item]; ok {
		Warning.Println(fmt.Sprintf("The price of %s is changed from %s to %s.",
			item, oldPrice, price))
	}
	prices[item] = price
}

// Cart is a struct representing a shopping cart of items.
type Cart struct {
	Items      []string
	TotalPrice Price
}

// hasMilk returns whether the shopping cart has "milk".
func (c *Cart) hasMilk() bool {
	return c.HasItem("milk")
}

// HasItem returns whether the shopping cart has the provided item name.
func (c *Cart) HasItem(item string) bool {
	for _, x := range c.Items {
		if x == item {
			return true
		}
	}
	return false
}

// Error is a logger for errors
var Error = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

// AddItem adds the provided item to the cart and update the cart balance.
// If item is not found in the prices map, then do not add it and print an error.
// Bonus (1pt) - Use the "log" package to print the error to the user
func (c *Cart) AddItem(item string) {
	price, ok := Prices[item]
	if !ok {
		Error.Println(fmt.Sprintf("The item is not found in the prices map: %s.", item))
		return
	}
	c.Items = append(c.Items, item)
	c.TotalPrice += price
}

// Info is a logger for infos
var Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

// Checkout displays the final cart balance and clears the cart completely.
func (c *Cart) Checkout() {
	Info.Println(fmt.Sprintf("The final cart balance is %s.", c.TotalPrice))
	c.Items = make([]string, 0)
	c.TotalPrice = 0
}
