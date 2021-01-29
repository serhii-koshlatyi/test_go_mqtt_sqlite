package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/sergeykoshlatuu/test_go_mqtt_sqlite/pkg/config"
	"github.com/sergeykoshlatuu/test_go_mqtt_sqlite/pkg/mosquitto"
	"github.com/sergeykoshlatuu/test_go_mqtt_sqlite/pkg/processor"
	"github.com/sergeykoshlatuu/test_go_mqtt_sqlite/pkg/store/sqlitestore"
)

func main() {
	var configPath string

	flag.StringVar(&configPath, "config", "./config/config.toml", "Path to the configuration file")

	flag.Parse()

	cfg, err := config.ReadConfig(configPath)
	if err != nil {
		panic(err)
	}

	broker, err := mosquitto.NewMosquittoBroker(cfg.Mosquitto)
	if err != nil {
		panic(err)
	}

	broker.Subscribe()

	store, err := sqlitestore.NewStore(cfg.SQLite)

	processor := processor.NewStateProcessor(store)

	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		sign := <-quitCh
		fmt.Printf("Received signal: %v. Running graceful shutdown...\n", sign)
		broker.Unsubscribe()
	}()

	processor.Run(broker.ResultChan())
}
