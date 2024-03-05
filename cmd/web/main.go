package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {
	port := flag.String("port", ":3000", "Server Port")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &application{
		logger: logger,
	}

	logger.Info("starting server", slog.String("port", *port))

	err := http.ListenAndServe(*port, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
