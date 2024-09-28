package main

import (
	"fmt"
	"log"

	"github.com/P3T3R2002/blog_aggreGATOR/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Reading config error: %v", err)
	}
	fmt.Printf("Reading config: %+v\n", cfg)

	err = cfg.SetUser("Peter20022")

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("Reading config error: %v", err)
	}
	fmt.Printf("Reading config again: %+v\n", cfg)
} 

main()