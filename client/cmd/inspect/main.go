package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ddelnano/terraform-provider-mikrotik/client"
	"github.com/ddelnano/terraform-provider-mikrotik/client/pkg/inspect"
)

func main() {
	if err := run(); err != nil {
		log.Println(fmt.Errorf("application failed: %w", err))
		os.Exit(1)
	}
}

func run() error {
	var (
		config *inspect.Config = &inspect.Config{}
	)
	flag.StringVar(&config.Root, "root", "/", "The root path to start inspection from, e.g.: /ip/dns2, /system/script")
	flag.IntVar(&config.Depth, "depth", 1, "Depth of inspection: -1 - no limit, 0 - root metadata, 1.. - up to desired level")
	flag.Parse()
	// -----------------------------------
	mikrotik := client.NewClient(client.GetConfigFromEnv())
	mc, err := mikrotik.GetMikrotikClient()
	if err != nil {
		return err
	}

	node, err := inspect.Do(mc, config)
	if err != nil {
		return err
	}

	if err := inspect.WriteJSON(os.Stdout, node); err != nil {
		return err
	}

	return nil
}
