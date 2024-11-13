package main

import (
	"fmt"
	ewallet "unit-test/app"
)

func main() {
	fmt.Println(ewallet.Run([]string{"deposit", "deposit", "deposit", "withdraw", "withdraw", "withdraw"}))
}
