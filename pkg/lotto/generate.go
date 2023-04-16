package lotto

import (
	"math/rand"
	"sort"
	"time"
)

var prevSeed int64

// GenerateLottoNumbers 1부터 45까지의 숫자 중에서 6개를 랜덤으로 생성한다.
func GenerateLottoNumbers(numberOfGenerate int) [][]int {
	lottoNumbers := make([][]int, 0, numberOfGenerate)
	for i := 0; i < numberOfGenerate; i++ {
		lottoNumbers = append(lottoNumbers, generateLottoNumber())
	}
	return lottoNumbers
}

// generateLottoNumber 1부터 45까지의 숫자 중에서 6개를 랜덤으로 생성한다.
func generateLottoNumber() []int {
	// 랜덤 시드를 지역적으로 생성한다.
	seed := time.Now().UnixNano()
	for seed == prevSeed {
		seed = time.Now().UnixNano()
	}
	prevSeed = seed

	source := rand.NewSource(seed)
	rng := rand.New(source)

	numbers := make(map[int]bool)

	for len(numbers) < 6 {
		number := rng.Intn(45) + 1
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