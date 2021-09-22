package main

import (
	"fmt"
	"math"
)

func recrusiveOne(n float64) float64 {
	if n == 0 {
		return n
	}
	current := 5 * math.Pow(10, n-1)
	fmt.Printf("%.0f + ", current)
	return current + recrusiveOne(n-1)
}

func recrusiveTwo(x, n float64) float64 {
	if n == 0 {
		return n
	}
	numerator := math.Pow(x, n)
	denominator := math.Pow(n, 2)
	fmt.Printf("(%.0f ^ %.0f / %.0f) + ", x, n, denominator)
	return numerator/denominator + recrusiveTwo(x, n-1)
}

func findCapital(text []byte) int {
	if len(text) == 0 {
		return 0
	}

	var count int = 0
	first := text[0]
	if first >= 65 && first <= 90 {
		count = 1
	}

	return count + findCapital(text[1:])
}

func main() {
	var n float64
	fmt.Print("Masukan n deret: ")
	fmt.Scanln(&n)
	fmt.Println("\nHasil: ", recrusiveOne(n))

	var x float64
	fmt.Print("Masukan x: ")
	fmt.Scanln(&x)
	fmt.Print("Masukan n deret: ")
	fmt.Scanln(&n)
	fmt.Println("\nHasil: ", recrusiveTwo(x, n))

	var word []byte
	fmt.Print("Masukan kata: ")
	fmt.Scanln(&word)
	fmt.Println("Hasil: ", findCapital(word))
}
