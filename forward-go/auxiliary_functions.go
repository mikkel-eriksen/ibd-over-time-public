package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func sum(a []int) int {
	s := 0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	return s
}

func sum_float(a []float64) float64 {
	s := 0.0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	return s
}

func sum_vertical(a [][]int, j int) int {
	s := 0
	for i := 0; i < len(a); i++ {
		s += a[i][j]
	}
	return s
}

func mean(a []int) float64 {
	s := sum(a)
	return float64(s) / float64(len(a))
}

func mean_float(a []float64) float64 {
	s := sum_float(a)
	return s / float64(len(a))
}

func mean_non_zero(a []int) float64 {
	s := sum(a)
	n := 0
	for i := 0; i < len(a); i++ {
		if a[i] != 0 {
			n++
		}
	}
	return float64(s) / float64(n)
}

func new_matrix(n, m int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			matrix[i] = append(matrix[i], 0)
		}
	}
	return matrix
}

func min(a []int) int {
	sort.Ints(a)
	return a[0]
}

func max(a []int) int {
	sort.Ints(a)
	return a[len(a)-1]
}

func read_file(filePath string) []int {
	// Read the file
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	var integers []int

	// Split the file contents into individual strings using spaces as the delimiter
	parts := strings.Fields(string(data))

	// Iterate through the strings and convert them to integers
	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			fmt.Printf("Error converting string to integer: %v\n", err)
			continue // Skip the invalid integer and continue with the next one
		}
		integers = append(integers, num)
	}

	return integers
}

func printJSON(data map[string]interface{}) {
	// Convert the map to JSON
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatalf("Failed to encode JSON: %v", err)
	}

	// Print the JSON to stdout
	fmt.Println(string(jsonData))
}
