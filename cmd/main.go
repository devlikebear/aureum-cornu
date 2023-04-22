package main

import (
	"net/http"
	"strconv"

	"github.com/devlikebear/aureum-cornu/pkg/lotto"
	"github.com/devlikebear/aureum-cornu/ui"
	"github.com/gin-gonic/gin"
)

func main() {
	
	r := gin.Default()

	// frontend 포함
	r.StaticFS("/web", ui.GetFileSystem())

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/web")
	})

	r.GET("/api/lottos", lottoHandler)

	r.Run(":8080")
}

func lottoHandler(c *gin.Context) {
    // 생성할 로또번호 개수를 파라미터로 받는다.
    count := c.Query("count")
    n, _  := strconv.Atoi(count)

    number := lotto.GenerateLottoNumbers(n)


    c.JSON(http.StatusOK, map[string]interface{}{
        "numbers": number,
    })
}