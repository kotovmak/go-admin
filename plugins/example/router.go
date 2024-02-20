package example

import (
	"github.com/kotovmak/go-admin/context"
	"github.com/kotovmak/go-admin/modules/auth"
	"github.com/kotovmak/go-admin/modules/db"
	"github.com/kotovmak/go-admin/modules/service"
)

func (e *Example) initRouter(prefix string, srv service.List) *context.App {

	app := context.NewApp()
	route := app.Group(prefix)
	route.GET("/example", auth.Middleware(db.GetConnection(srv)), e.TestHandler)

	return app
}
