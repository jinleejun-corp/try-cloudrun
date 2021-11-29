// Sample run-helloworld is a minimal Cloud Run service.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "jinleejun-backend",
	Short:   "jinleejun",
	Example: "go run main.go",
	Run: func(cmd *cobra.Command, args []string) {
		log.Print("starting server...")
		http.HandleFunc("/", handler)

		// Determine port for HTTP service.
		port := os.Getenv("PORT")
		if port == "" {
			port = "8080"
			log.Printf("defaulting to port %s", port)
		}

		// Start HTTP server.
		log.Printf("listening on port %s", port)
		if err := http.ListenAndServe(":"+port, nil); err != nil {
			log.Fatal(err)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	Execute()
}

func handler(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello %s!\n", name)
}
