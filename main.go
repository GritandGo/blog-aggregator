package main

import (
	"fmt"
	"os"

	"github.com/GritandGo/blog-aggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	

	
	err = cfg.SetUser("Ben")
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
		}		


	cfg, err = config.Read()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	
	}
	fmt.Println(cfg)
}


type state struct {
	cfg *config.Config
}