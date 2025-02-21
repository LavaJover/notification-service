package main

import (
	"fmt"

	"github.com/LavaJover/notification-service/internal/config"
)

func main(){
	cfg := config.MustLoad()

	fmt.Println(cfg)
}