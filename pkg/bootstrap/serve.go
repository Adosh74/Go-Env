package bootstrap

import (
	"github.com/Adosh74/Go-Env/pkg/config"
	"github.com/Adosh74/Go-Env/pkg/html"
	"github.com/Adosh74/Go-Env/pkg/routing"
	"github.com/Adosh74/Go-Env/pkg/static"
)

func Serve() {
	config.Set()

	routing.Init()

	static.LoadStatic(routing.GetRouter())

	html.LoadHTML(routing.GetRouter())

	routing.RegisterRoute()

	routing.Serve()
}
