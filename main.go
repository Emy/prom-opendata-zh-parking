package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/Emy/prom-opendata-zh-parking/internal/handlers"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	handlers.InitializePrometheusHandling()
	logger.Info("main.go: Starting web server on port :4277")
	http.ListenAndServe(":4277", nil)
}
