package bootstrap

import (
	"github.com/realwebdev/blog/pkg/config"
	"github.com/realwebdev/blog/pkg/database"
	"github.com/realwebdev/blog/pkg/html"
	"github.com/realwebdev/blog/pkg/routing"
	"github.com/realwebdev/blog/pkg/static"
)

func Serve() {
	config.Set()

	database.Connect()

	routing.Init()

	static.LoadStatic(routing.GetRouter())

	html.LoadHTML(routing.GetRouter())

	routing.RegisterRoutes()

	routing.Serve()

}
