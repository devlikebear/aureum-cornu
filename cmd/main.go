package main

import (
	"fmt"

	"github.com/devlikebear/aureum-cornu/pkg/lotto"
)

func main() {
	lottoNumbers := lotto.GenerateLottoNumbers()
	fmt.Println("생성된 로또 번호:", lottoNumbers)
}
