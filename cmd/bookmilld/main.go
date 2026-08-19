package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/TrueBlocks/trueblocks-art/packages/appd"
	"github.com/TrueBlocks/trueblocks-art/packages/cli"
)

var version = "dev"

func main() {
	app := cli.App{
		Name:        "bookmilld",
		Description: "Web server for the bookmill pipeline dashboard.",
		Version:     version,
		Flags: []cli.FlagDef{
			{Name: "addr", Help: "listen address", Default: ":8082"},
			{Name: "apps-config", Help: "path to daemons.json for cross-app nav", Default: appd.DefaultConfigPath()},
		},
		Run: run,
	}
	cli.Exit(app.Main())
}

func run(c *cli.Context) error {
	mux := http.NewServeMux()
	if _, err := appd.RegisterNav(mux, c.String("apps-config")); err != nil {
		return fmt.Errorf("registering nav: %w", err)
	}

	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = fmt.Fprint(w, `<!doctype html>
<html>
<head>
  <meta charset="utf-8">
  <title>Bookmill Dashboard</title>
  <link rel="stylesheet" href="/__tb__/nav.css">
</head>
<body>
  <script src="/__tb__/nav.js"></script>
  <main style="padding:1rem;font-family:system-ui,sans-serif;">
    <h1>Bookmill Dashboard</h1>
    <p>Pipeline status and controls will appear here.</p>
  </main>
</body>
</html>`)
	})
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	})

	addr := c.String("addr")
	srv := &http.Server{
		Addr:              addr,
		Handler:           mux,
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      60 * time.Second,
		IdleTimeout:       120 * time.Second,
	}

	go func() {
		c.Logger.Info("bookmilld listening", "addr", addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			c.Logger.Error("listen error", "err", err.Error())
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	c.Logger.Info("shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return srv.Shutdown(ctx)
}
