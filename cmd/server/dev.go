package main

import (
	"log"

	"github.com/roscopecoltran/configor"
)

func dumpConfig() {
	if err := configor.Dump(Cfg.Krakend, "krakend", "yaml,json", "./shared/config/dump"); err != nil {
		log.Fatal("ERROR while dumping the config struct:", err.Error())
	}
	if err := configor.Dump(Cfg.Service, "service", "yaml,json", "./shared/config/dump"); err != nil {
		log.Fatal("ERROR while dumping the config struct:", err.Error())
	}

	if err := configor.Dump(Cfg.Providers, "providers", "yaml,json", "./shared/config/dump"); err != nil {
		log.Fatal("ERROR while dumping the config struct:", err.Error())
	}

	if err := configor.Dump(Cfg, "all", "yaml,json", "./shared/config/dump"); err != nil {
		log.Fatal("ERROR while dumping the config struct:", err.Error())
	}
}
