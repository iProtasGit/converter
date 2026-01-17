package main

import "fmt"

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
func main() {
	const (
		UsdToEur = 2
		UsdToRub = 2
		EurToRub = UsdToRub / UsdToEur
	)
	text := inputData()
	fmt.Println(text, EurToRub)
}

func inputData() (text string) {
	_, err := fmt.Scan(&text)
	if err != nil {
		panic(err)
	}
	return text
}

func calculateCurrency(num, curr1, curr2 float64) {}
