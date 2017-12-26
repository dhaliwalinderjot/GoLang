package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	//const jsonStream = `{"BTC":"1066000.0","ETH":"54999.0","BCH":"214500.0","XRP":"76.28","LTC":"20807.0","MIOTA":230.42,"OMG":920.99,"GNT":42.14}`
	const jsonStream = `{"prices":{"BTC":"1066000.0","ETH":"54999.0","BCH":"214500.0","XRP":"76.28","LTC":"20807.0","MIOTA":230.42,"OMG":920.99,"GNT":42.14},"stats":{"ETH":{"last_traded_price":"54999.0","lowest_ask":"54999.0","highest_bid":"54800.0","min_24hrs":"49125.0","max_24hrs":"57000.0","vol_24hrs":"4232.708"},"BTC":{"last_traded_price":"1066000.0","lowest_ask":"1066001.0","highest_bid":"1066000.0","min_24hrs":"1010000.0","max_24hrs":"1098000.0","vol_24hrs":"243.5544"},"LTC":{"last_traded_price":"20807.0","lowest_ask":"20818.0","highest_bid":"20808.0","min_24hrs":"19700.0","max_24hrs":"21448.0","vol_24hrs":"7497.943"},"XRP":{"last_traded_price":"76.28","lowest_ask":"76.28","highest_bid":"76.1","min_24hrs":"70.7","max_24hrs":"77.0","vol_24hrs":"3071304.8"},"BCH":{"last_traded_price":"214500.0","lowest_ask":"214500.0","highest_bid":"214499.0","min_24hrs":"205000.0","max_24hrs":"222100.0","vol_24hrs":"920.877"}}}`
	
	type Message struct {
		prices struct{
			 BTC, ETH, BCH, LTC, XRP string
		}
	}
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		var m Message
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Price of Bitcoin: %s\nPrice of Ether: %s\nPrice of Bitcoin Cash: %s\nPrice of Ripple: %s\nPrice of LiteCoin: %s\n", m.prices.BTC)
		//fmt.Printf("Price of Bitcoin: %s\nPrice of Ether: %s\nPrice of Bitcoin Cash: %s\nPrice of Ripple: %s\nPrice of LiteCoin: %s\n", m.prices.BTC, m.ETH, m.BCH, m.XRP, m.LTC)
	}
}
