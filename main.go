package main

import (
	"net/http"
	"os"
	"log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
        fs := http.FileServer(http.Dir("./assets"))
        http.Handle("/home", fs)
        log.Println("Listening on :8080 ...")
        err := http.ListenAndServe(":8080", nil)
        if err != nil {
           log.Fatal(err)
}
	// e := echo.New()
         
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())
         
	// e.GET("/", func(c echo.Context) error {
	// 	return c.HTML(http.StatusOK, "Look at this Docker GO server! v0.0.4")
	// })

	// e.GET("/ping", func(c echo.Context) error {
	// 	return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	// })
        
	 httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}
        // e.Logger.Fatal(e.Start(":" + httpPort))
}
