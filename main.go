package main

import (
	"errors"
	"fmt"
)

func main() {
	for {
		fmt.Print("Welcome to the iProtas change! \nFastest and privacy currencies converter on the Golang🩵\n\n")
		starterCurrency := inputStarterCur()
		amount := inputAmount()
		mainCurrency := inputMainCurrency(starterCurrency)
		fmt.Println(amount, starterCurrency, mainCurrency)
		result, err := calculateCurrency(amount, starterCurrency, mainCurrency)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%.02f %s converted to %s = %.02f\n\n", amount, starterCurrency, mainCurrency, result)

		fmt.Println("Thanks for choice us! See you later! Enter 'exit' to quit.")
		var choice string
		_, err = fmt.Scanln(&choice)
		if err != nil {
			panic(err)
		}
		if choice == "exit" {
			break
		}

	}
}

func inputStarterCur() (currency string) {
	fmt.Print("What is the currency you want to convert?\nAvailable: USD, EUR, RUB\n\nPlease enter currency: ")

	for {
		_, err := fmt.Scan(&currency)
		if err != nil {
			fmt.Printf("error reading input: %s. Please try again: ", err)
		} else if currency != "USD" && currency != "EUR" && currency != "RUB" {
			fmt.Print("unavailable currency. Please try again: ")
		} else {
			break
		}
	}

	return currency
}

func inputAmount() (amount float64) {
	fmt.Print("\nPlease enter amount: ")

	for {
		_, err := fmt.Scan(&amount)
		if err != nil {
			fmt.Printf("error reading input: %s. Please try again: ", err)
		} else {
			break
		}
	}

	return amount
}

func inputMainCurrency(starterCurrency string) (mainCurrency string) {
	fmt.Print("\n\nChoose a currency you want to convert.")

	switch starterCurrency {
	case "USD":
		fmt.Println("\nAvailable currencies to convert: EUR, RUB")
	case "EUR":
		fmt.Println("Available currencies to convert: USD, RUB")
	case "RUB":
		fmt.Println("Available currencies to convert: EUR, USD")
	}

	fmt.Print("Please enter currency: ")

	for {
		_, err := fmt.Scan(&mainCurrency)
		if err != nil {
			fmt.Printf("error reading input: %s. Please try again: ", err)
		} else if mainCurrency != "USD" && mainCurrency != "EUR" && mainCurrency != "RUB" {
			fmt.Print("unavailable currency. Please try again: ")
		} else if mainCurrency == starterCurrency {
			fmt.Printf("Currency %s not converted to %s. Please try again: ", mainCurrency, starterCurrency)
		} else {
			break
		}
	}

	return mainCurrency
}

func calculateCurrency(amount float64, starterCurr, mainCurr string) (float64, error) {
	const (
		UsdToEur = 2
		UsdToRub = 2
		EurToRub = 1
	)

	if (starterCurr == "EUR" && mainCurr == "USD") || (starterCurr == "USD" && mainCurr == "EUR") {
		return amount * UsdToEur, nil
	}
	if (starterCurr == "USD" && mainCurr == "RUB") || (starterCurr == "RUB" && mainCurr == "USD") {
		return amount * UsdToRub, nil
	}
	if (starterCurr == "RUB" && mainCurr == "EUR") || (starterCurr == "EUR" && mainCurr == "RUB") {
		return amount * EurToRub, nil
	}
	return 0, errors.New("something went wrong. please contact the tech support team")
}
