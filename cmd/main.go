// main.go
package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/devlikebear/aureum-cornu/pkg/lotto"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
    app := pocketbase.New()

    // 로또번호 조회하는 API 추가
    app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
        e.Router.AddRoute(echo.Route{
            Method:  http.MethodGet,
            Path:    "/api/lottos",
            Handler: lottoHandler,
        })
        return nil  
    })


    if err := app.Start(); err != nil {
        log.Fatal(err)
    }
}

func lottoHandler(c echo.Context) error {
    // 생성할 로또번호 개수를 파라미터로 받는다.
    count := c.QueryParam("count")
    n, _  := strconv.Atoi(count)

    number := lotto.GenerateLottoNumbers(n)
    return c.JSON(http.StatusOK, map[string]interface{}{
        "numbers": number,
    })
}