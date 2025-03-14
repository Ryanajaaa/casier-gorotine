package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

type Item struct {
	Name  string
	Price float64
}

type Customer struct {
	Name  string
	Items []Item
}

func cashier(id int, customers <-chan Customer, wg *sync.WaitGroup) {
	defer wg.Done()
	for customer := range customers {
		fmt.Printf("Kasir %d melayani %s dengan %d barang\n", id, customer.Name, len(customer.Items))
		processingTime := time.Duration(rand.Intn(3)+1) * time.Second
		time.Sleep(processingTime)
		fmt.Printf("Kasir %d selesai melayani %s\n", id, customer.Name)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	items := []Item{
		{"Buku", 15000}, {"Pensil", 5000}, {"Penghapus", 3000},
		{"Penggaris", 7000}, {"Bolpoin", 8000}, {"Tas", 200000},
	}

	var customers []Customer
	for i := 0; i < 100; i++ {
		name := gofakeit.Name()
		itemCount := rand.Intn(2) + 1 // Setiap pelanggan membawa 1-3 barang
		var customerItems []Item
		for j := 0; j < itemCount; j++ {
			customerItems = append(customerItems, items[rand.Intn(len(items))])
		}
		customers = append(customers, Customer{Name: name, Items: customerItems})
	}

	customerChannel := make(chan Customer, len(customers))
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go cashier(i, customerChannel, &wg)
	}

	for _, customer := range customers {
		customerChannel <- customer
	}
	close(customerChannel)

	wg.Wait()
	fmt.Println("Semua transaksi selesai!")
}
