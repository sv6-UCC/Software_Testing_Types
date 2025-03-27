package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"strconv"
)

func generateBVC(minL, maxL []int, numVars int) {
	var testCases [][]string

	headers := []string{"Test Case Id"}
	for i := 0; i < numVars; i++ {
		headers = append(headers, string('A'+i))
	}
	headers = append(headers, "Expected Output")

	testCases = append(testCases, headers)

	testID := 1
	for i := 0; i < numVars; i++ {
		values := []int{minL[i], minL[i] + 1, maxL[i] - 1, maxL[i], int(math.Ceil(float64(minL[i]+maxL[i]) / 2))}
		for _, v := range values {
			row := []string{strconv.Itoa(testID)}
			for j := 0; j < numVars; j++ {
				if j == i {
					row = append(row, strconv.Itoa(v))
				} else {
					row = append(row, strconv.Itoa((minL[j]+maxL[j])/2))
				}
			}
			expectedOutput := computeExpectedOutput(row[1:])
			row = append(row, expectedOutput)
			testCases = append(testCases, row)
			testID++
		}
	}

	file, err := os.Create("BVC.csv")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, row := range testCases {
		err := writer.Write(row)
		if err != nil {
			fmt.Println("Error writing to CSV:", err)
			return
		}
	}

	fmt.Println("BVC test cases generated and saved in BVC.csv")
}

func computeExpectedOutput(inputs []string) string {

	sum := 0
	for _, val := range inputs {
		num, _ := strconv.Atoi(val)
		sum += num
	}
	return strconv.Itoa(sum)
}

func main() {
	var minL, maxL []int
	var numVars int

	fmt.Print("How Many Parameters? ")
	fmt.Scan(&numVars)

	for i := 0; i < numVars; i++ {
		var min, max int
		fmt.Printf("Enter min & max for %dth Parameter: ", i+1)
		fmt.Scan(&min, &max)
		minL = append(minL, min)
		maxL = append(maxL, max)
	}

	generateBVC(minL, maxL, numVars)
}
