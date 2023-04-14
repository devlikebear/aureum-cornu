package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func generateLottoNumbers() []int {
	rand.Seed(time.Now().UnixNano())
	numbers := make(map[int]bool)
	
	for len(numbers) < 6 {
		number := rand.Intn(45) + 1
		if _, exists := numbers[number]; !exists {
			numbers[number] = true
		}
	}

	lottoNumbers := make([]int, 0, 6)
	for number := range numbers {
		lottoNumbers = append(lottoNumbers, number)
	}

	sort.Ints(lottoNumbers)
	return lottoNumbers
}

func main() {
	lottoNumbers := generateLottoNumbers()
	fmt.Println("생성된 로또 번호:", lottoNumbers)
}
