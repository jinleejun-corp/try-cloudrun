// Sample run-helloworld is a minimal Cloud Run service.
package main

import (
	"fmt"
	"net/http"
	"os"
	"try/cmd"
)

func main() {
	cmd.Execute()
}

func handler(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello %s!\n", name)
}
