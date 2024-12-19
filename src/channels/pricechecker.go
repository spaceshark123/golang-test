package main

import (
	"fmt"
	"time"
	"math/rand"
	"sync"
)

// sale info struct
type SaleInfo struct {
	website string
	price float64
}

// product type (can be TV, laptop, phone)
type ProductType int
const (
	TV ProductType = iota
	Laptop
	Phone
)

// product struct
type Product struct {
	website string
	productType ProductType
}

var websites = []string{"amazon.com", "ebay.com", "walmart.com"}
var prices = map[Product]float64{}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go updatePrices(websites, prices, &wg)
	wg.Wait() // wait for initial prices to be set

	var tvChannel = make(chan SaleInfo) // store price and website
	var laptopChannel = make(chan SaleInfo) // store price and website
	var phoneChannel = make(chan SaleInfo) // store price and website
	for i := range websites {
		go checkTVPrice(websites[i], tvChannel, 70)
		go checkLaptopPrice(websites[i], laptopChannel, 50)
		go checkPhonePrice(websites[i], phoneChannel, 30)
	}
	
	sendNotifications(tvChannel, laptopChannel, phoneChannel) // find the first sale for each product type
}

func sendNotifications(tvChannel chan SaleInfo, laptopChannel chan SaleInfo, phoneChannel chan SaleInfo) {
	foundPhoneSale := false
	foundLaptopSale := false
	foundTVSale := false
	for !foundPhoneSale || !foundLaptopSale || !foundTVSale {
		select {
			case saleInfo := <-tvChannel:
				if foundTVSale {
					continue
				}
				fmt.Printf("TV on sale at %s for $%.2f\n", saleInfo.website, saleInfo.price)
				foundTVSale = true
			case saleInfo := <-laptopChannel:
				if foundLaptopSale {
					continue
				}
				fmt.Printf("Laptop on sale at %s for $%.2f\n", saleInfo.website, saleInfo.price)
				foundLaptopSale = true
			case saleInfo := <-phoneChannel:
				if foundPhoneSale {
					continue
				}
				fmt.Printf("Phone on sale at %s for $%.2f\n", saleInfo.website, saleInfo.price)
				foundPhoneSale = true
		}
	}
}

func checkTVPrice(website string, tvChannel chan SaleInfo, maxPrice float64) {
	for {
		var price = prices[Product{website, TV}]
		if price <= maxPrice {
			tvChannel <- SaleInfo{website, price}
		}
		time.Sleep(1 * time.Second)
	}
}

func checkLaptopPrice(website string, laptopChannel chan SaleInfo, maxPrice float64) {
	for {
		var price = prices[Product{website, Laptop}]
		if price <= maxPrice {
			laptopChannel <- SaleInfo{website, price}
		}
		time.Sleep(1 * time.Second)
	}
}

func checkPhonePrice(website string, phoneChannel chan SaleInfo, maxPrice float64) {
	for {
		var price = prices[Product{website, Phone}]
		if price <= maxPrice {
			phoneChannel <- SaleInfo{website, price}
		}
		time.Sleep(1 * time.Second)
	}
}

func updatePrices(websites []string, prices map[Product]float64, wg *sync.WaitGroup) {
	// keep updating prices to random values to simulate real-time data
	for i := 0;; i++ {
		for i := range websites {
			prices[Product{websites[i], TV}] = rand.Float64() * 100
			prices[Product{websites[i], Laptop}] = rand.Float64() * 100
			prices[Product{websites[i], Phone}] = rand.Float64() * 100
		}
		if i == 0 {
			wg.Done() // done setting initial prices
		}
		time.Sleep(100 * time.Millisecond)
	}
}

