package main

import "fmt"

func main() {
	// Hitung rata-rata
	scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}
	lenght := len(scores)
	var total int
	for _, val := range scores {
		total += val
	}
	avarage := float64(total) / float64(lenght)
	fmt.Println(avarage)

	// mencari nilai lebih dari sama dengan 90
	score := [8]int{100, 80, 75, 92, 70, 93, 88, 67}
	var goodScore []int
	for _, val := range score {
		if val >= 90 {
			goodScore = append(goodScore, val)
		}
	}
	fmt.Println(goodScore)

}