package lotto

import (
	"testing"
)

// TestGenerateLottoNumbersOutOfRange 로또 번호가 범위를 벗어나는지 확인한다.
func TestGenerateLottoNumbersOutOfRange(t *testing.T) {
	// Given
	// When
	lottoNumbers := GenerateLottoNumbers()
	// Then
	for _, number := range lottoNumbers {
		if number < 1 || number > 45 {
			t.Errorf("Expected lotto number to be between 1 and 45, got %d", number)
		}
	}
}

// TestGenerateLottoNumbersDuplicates 로또 번호가 중복되는지 확인한다.
func TestGenerateLottoNumbersDuplicates(t *testing.T) {
	// Given
	// When
	lottoNumbers := GenerateLottoNumbers()
	// Then
	numbers := make(map[int]bool)
	for _, number := range lottoNumbers {
		if _, exists := numbers[number]; exists {
			t.Errorf("Expected lotto number to be unique, got %d", number)
		}
		numbers[number] = true
	}
}

// TestGenerateLottoNumbersLength 로또 번호가 6개가 아닌지 확인한다.
func TestGenerateLottoNumbersLength(t *testing.T) {
	// Given
	// When
	lottoNumbers := GenerateLottoNumbers()
	// Then
	if len(lottoNumbers) != 6 {
		t.Errorf("Expected lotto number to be 6, got %d", len(lottoNumbers))
	}
}

// TestGenerateLottoNumbersSorted 로또 번호가 오름차순으로 정렬되어 있는지 확인한다.
func TestGenerateLottoNumbersSorted(t *testing.T) {
	// Given
	// When
	lottoNumbers := GenerateLottoNumbers()
	// Then
	for i := 0; i < len(lottoNumbers)-1; i++ {
		if lottoNumbers[i] > lottoNumbers[i+1] {
			t.Errorf("Expected lotto number to be sorted, got %d", lottoNumbers)
		}
	}
}

// 로또번호를 N번 생성하여 중복되는 번호가 있는지 확인한다.
func TestGenerateLottoNumbersDuplicatesN(t *testing.T) {
	// Given
	// When
	lottoNumbers := make([][]int, 0, 5)
	for i := 0; i < 5; i++ {
		lottoNumbers = append(lottoNumbers, GenerateLottoNumbers())
	}
	// Then
	// 생성된 로또번호가 중복되는지 비교한다.
	for i := 0; i < len(lottoNumbers)-1; i++ {
		for j := i + 1; j < len(lottoNumbers); j++ {
			if lottoNumbers[i][0] == lottoNumbers[j][0] &&
				lottoNumbers[i][1] == lottoNumbers[j][1] &&
				lottoNumbers[i][2] == lottoNumbers[j][2] &&
				lottoNumbers[i][3] == lottoNumbers[j][3] &&
				lottoNumbers[i][4] == lottoNumbers[j][4] &&
				lottoNumbers[i][5] == lottoNumbers[j][5] {
				t.Errorf("Expected lotto number to be unique, got %v, %v", lottoNumbers[i], lottoNumbers[j])
			}
		}
	}
}