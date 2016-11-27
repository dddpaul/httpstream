package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func stream(c echo.Context) error {
	resp := c.Response()
	gone := resp.CloseNotify()
	resp.Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	resp.WriteHeader(http.StatusOK)
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		fmt.Fprintf(resp, "{ \"time\": \"%v\" }\n", time.Now())
		resp.Flush()
		select {
		case <-ticker.C:
		case <-gone:
			break
		}
	}
}

func main() {
	var port string
	flag.StringVar(&port, "port", ":8080", "Port to listen (prepended by colon), i.e. :8080")
	flag.Parse()

	e := echo.New()
	e.GET("/", stream)
	e.Logger.Fatal(e.Start(port))
}
