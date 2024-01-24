package problem1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type DoubleDigit struct {
	first  int
	second int
}

func (d DoubleDigit) GetInt() (int, error) {
	first := strconv.Itoa(d.first)
	second := strconv.Itoa(d.second)

	if second == "0" {
		second = first
	}

	str_digit := first + second

	number, err := strconv.Atoi(str_digit)

	if err != nil {
		return 0, err
	}

	return number, nil
}

func P1_trebuchet() {
	fmt.Println("Trebuchet!")

	file, err := os.Open("problem1/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	digits := []DoubleDigit{}
	linesRead := 0

	// Combine the first and last digit, and form a single two-digit number.
	for scanner.Scan() {
		current := DoubleDigit{first: 0, second: 0}
		line := scanner.Text()
		linesRead++

		for i := 0; i < len(line); i++ {
			maybeDigit := string(line[i])
			digit, err := strconv.Atoi(maybeDigit)
			if err != nil {
				continue
			}

			if current.first == 0 {
				current.first = digit
			} else {
				current.second = digit
			}

		}

		digits = append(digits, current)

		fmt.Printf("For %s, the digits are %d and %d \n", line, current.first, current.second)
	}

	sum := 0
	for i := 0; i < len(digits); i++ {
		val, err := digits[i].GetInt()
		if err != nil {
			fmt.Print(err)
		}
		fmt.Printf("Sum (%d) + Val (%d) = (%d) \n", sum, val, sum+val)
		sum = sum + val

	}

	fmt.Print(sum)
	fmt.Printf("Lines Read: %d\n", linesRead)
}
