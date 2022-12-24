package main

import (
	"context"
	"fmt"
	"log"
	"mindx/external"
	"mindx/internal/router"
	"mindx/pkg/zapx"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	log.Println("starting server...")
	zapx.Info(context.TODO(), "starting server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conf, err := external.GetConfigs("./configs", "dev", "json")
	if err != nil {
		zapx.Fatal(ctx, "failed to get configs.", err)
	}

	ds, err := external.InitDS(conf)
	if err != nil {
		zapx.Fatal(ctx, "failed to initialize data sources.", err)
	}

	r := router.NewRouter(conf, ds)
	engine, err := r.Init()
	if err != nil {
		zapx.Fatal(ctx, "failed to init gin.", err)
	}

	srv := &http.Server{
		Addr:    conf.Host + ":" + conf.Port,
		Handler: engine,
	}

	go func() {
		if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zapx.Fatal(ctx, "failed to initialize server.", err)
		}
	}()

	log.Println(fmt.Sprintf("listening on port: %v", conf.Port))
	zapx.Info(context.TODO(), fmt.Sprintf("listening on port: %v\n.", conf.Port))
	quit := make(chan os.Signal, 1)

	var signalsToIgnore = []os.Signal{os.Interrupt}
	signal.Notify(quit, signalsToIgnore...)

	<-quit

	if err = ds.Close(); err != nil {
		zapx.Fatal(ctx, "a problem occurred gracefully shutting down sources.", err)
	}

	log.Println("shutting down server...")
	zapx.Info(context.TODO(), "shutting down server...")
	if err = srv.Shutdown(ctx); err != nil {
		zapx.Fatal(ctx, "server forced to shutdown.", err)
	}
}
