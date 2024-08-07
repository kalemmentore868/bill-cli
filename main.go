package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error){
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	input = strings.TrimSpace(input)
	return input, err
}

func createBill() bill{
	reader := bufio.NewReader(os.Stdin)
	name, _ := getInput("Create a name for your bill: ", reader)
	return newBill(name)
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose option (a - add an item, s - save the bill, t - give a tip): ", reader)

	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item Price: ", reader)

		p, err := strconv.ParseFloat(price, 64)

		if err != nil {
			fmt.Println("The price must be a number")
			promptOptions(b)
		}

		b.addItem(name, p)
		
		fmt.Println("Item added - ", name, price)
		promptOptions(b)
	case "t":
		tip, _ := getInput("Tip: ", reader)


		t, err := strconv.ParseFloat(tip, 64)

		if err != nil {
			fmt.Println("The tip must be a number")
			promptOptions(b)
		}
		
		b.updateTip(t)
		fmt.Println("tip added - ", tip)
		promptOptions(b)

	case "s":
		b.save()
	default:
		fmt.Println("That was not a valid option")
		promptOptions(b)
	}
}


func main(){
	mybill := createBill()
	promptOptions(mybill)

}