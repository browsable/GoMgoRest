package main

import (
	"fmt"
	"GoServer/routes"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

const port string = ":3333"

func main() {

	fmt.Printf("Running at %v\n", port)

	e := echo.New()

	routes.Init(e)

	// e.Run(standard.WithTLS(port, "server/cert/server.crt", "server/cert/server.key"))
	e.Run(standard.New(port))
}
