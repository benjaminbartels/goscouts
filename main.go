package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	LOW_TEMP_F_WARNING  = 0.0
	HIGH_TEMP_F_WARNING = 100.0
	MAX_LOOP            = 5
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < MAX_LOOP; i++ {
		fmt.Print("\nEnter a temperature in Fahrenheit: ")

		var fahrenheit float64
		var celsius float64

		if scanner.Scan() {
			input := scanner.Text()
			var err error
			fahrenheit, err = strconv.ParseFloat(input, 64)
			if err != nil {
				fmt.Println("Data entry error - try again")
			}
			celsius = (fahrenheit - 32) * 5 / 9
		} else {
			fmt.Println("Data entry error - try again")
		}

		fmt.Printf("The temperature in Celsius is: %.2f\n", celsius)

		// Check for high temperature and issue a warning if necessary
		if fahrenheit > HIGH_TEMP_F_WARNING {
			fmt.Print("Remember to hydrate\n")
		}
		// Check for low temperature and issue a warning if necessary
		if fahrenheit < LOW_TEMP_F_WARNING {
			fmt.Print("Remember to pack Long underwear\n")
		}
	}
	os.Exit(-1)
}
