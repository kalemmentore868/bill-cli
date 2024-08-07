package main

import (
	"fmt"
	"os"
)

type bill struct {
	name string
	items map[string]float64
	tip float64


}

func newBill(name string) bill {
	b := bill{
		name: name,
		items: map[string]float64{"pie": 5.99, "cake": 3.99, "juice": 4.99},
		tip: 0,
	}

	return b
}

// format bill
func (b *bill) format() string{
	fs := "Bill brakdown: \n"
	var total float64 = 0

	for k,v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v)
		total += v
	}

	fs += fmt.Sprintf("%-25v ...$%0.2f\n", "Tip:", b.tip)

	fs += fmt.Sprintf("%-25v ...$%0.2f", "Total:", total + b.tip)

	return fs
}

func (b *bill) updateTip(tip float64){
	b.tip = tip
}

func (b *bill) addItem(name string, price float64){
	b.items[name] = price
}

func (b *bill) save(){
	data := []byte(b.format())

	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("Bill was saved to the file")
}