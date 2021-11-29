package cmd

import (
	"fmt"
	"log"

	// "jinleejun/app"
	// "jinleejun/db"
	// "jinleejun/routes"
	"net/http"
	"os"

	// "github.com/labstack/echo/v4"
	// "github.com/labstack/echo/v4/middleware"
	// "github.com/labstack/gommon/log"
	"github.com/spf13/cobra"
	// "golang.org/x/net/http2"
	// "gorm.io/gorm"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "run server",
	RunE: func(cmd *cobra.Command, args []string) error {

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
		return nil
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

// func GetNewDB(logger echo.Logger, dbDriver db.DbDriver) (*db.DB, error) {
// 	database := db.DB{
// 		DB:     &gorm.DB{},
// 		Logger: logger,
// 	}
// 	err := database.ConnectDB(dbDriver)
// 	if err != nil {
// 		database.Logger.Errorf("database.ConnectDB error: %+v", err)
// 		return nil, err
// 	}
// 	return &database, nil
// }

func handler(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	if name == "" {
		name = "World Serve"
	}
	fmt.Fprintf(w, "Hello %s!\n", name)
}
