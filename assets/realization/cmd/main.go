package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"log"
	"os"
	"template/internal/app"
	"template/internal/config"
	logging "template/pkg/logs"
	"template/pkg/printer"
	executor "template/pkg/utils/app"
	"time"
)

func init() {
	parser := argparse.NewParser("SXTP manager service", "Microservice for sxtp manager of controller")

	path := parser.String("c", "config", &argparse.Options{
		Required: false, Help: "Path to configuration file",
	})

	if err := parser.Parse(os.Args); err != nil {
		log.Fatal(parser.Usage(err))
	}

	if *path != "" {
		color.Green("running with config.env file")
		if err := godotenv.Load(*path); err != nil {
			log.Fatalf("Error getting env, not coming through %v", err)
		}
	} else {
		color.Blue("running without config.env file")
	}

	logging.InitLogger()
}

func main() {
	logger := logging.GetLogger()
	logger.Info(printer.GetCyan(fmt.Sprintf("Start service dc-sxtp-manager (%v)", time.Now().Format("02-01-2006 15:04.05"))))
	cfg := config.NewConfig()

	a := app.New(&cfg, logger)

	if err := executor.Run(a); err != nil {
		logger.Fatal(err)
	}
}
