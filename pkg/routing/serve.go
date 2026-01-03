package routing

import (
	"fmt"
	"log"

	"github.com/realwebdev/blog/pkg/config"
)

func Serve() {
	r := GetRouter()

	configs := config.Get()

	if err := r.Run(fmt.Sprintf("%s:%s", configs.Server.Host, configs.Server.Port)); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
