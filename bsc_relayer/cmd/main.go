package main

import (
	"bsc_relayer/pkg/relay"
	"context"
	"flag"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
)

func main() {
	configFile := flag.String("c", "./configs/config.json", "filename of the config file")
	flag.Parse()

	viper.SetConfigFile(*configFile)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("read config: %v", err)
	}

	var config relay.Config
	err := viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("unmarshal config into struct: %v", err)
	}

	relayer, err := relay.NewRelayer(config)
	if err != nil {
		log.Fatalf("new relayer: %v", err)
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		<-sig
		cancel()
	}()

	if err := relayer.Run(ctx); err != nil {
		log.Fatalf("run relayer: %v", err)
	}
}
