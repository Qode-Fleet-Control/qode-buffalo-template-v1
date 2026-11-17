// A minimal Buffalo app shaped for the fleet: it serves at the root of its own
// hostname, so routes mount directly on the app.
package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gobuffalo/buffalo"
)

func port() string {
	if p := strings.TrimSpace(os.Getenv("PORT")); p != "" {
		return p
	}
	return "8080"
}

func newApp() *buffalo.App {
	app := buffalo.New(buffalo.Options{
		Env:  "production",
		Addr: ":" + port(),
	})
	// KEEP THIS: one log line per request (status + duration) is what the workspace Monitoring tab charts.
	// requestLog (requestlog.go) takes the place of Buffalo's text request logger.
	app.Middleware.Replace(buffalo.RequestLogger, requestLog)

	app.GET("/health", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, buffaloJSON(map[string]string{"status": "ok"}))
	})
	return app
}

func main() {
	log.Printf("buffalo-template listening on :%s", port())
	if err := newApp().Serve(); err != nil {
		log.Fatal(err)
	}
}
