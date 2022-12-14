package cmd

import (
	"fmt"
	"log"

	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "run server",
	RunE: func(cmd *cobra.Command, args []string) error {
		e := echo.New()

		port := os.Getenv("PORT")
		if port == "" {
			port = "8080"
		}

		e.GET("/", handler)

		if err := e.Start(fmt.Sprintf(":%s", port)); err != http.ErrServerClosed {
			log.Fatal(err)
		}

		return nil
	},
}

func handler(e echo.Context) error {
	name := os.Getenv("NAME")
	if name == "" {
		name = "World Serve"
	}
	fmt.Fprintf(e.Response().Writer, "Hello %s!\n", name)

	return nil
}
