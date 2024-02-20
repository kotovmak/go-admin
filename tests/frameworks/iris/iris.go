package iris

import (
	// add iris adapter
	_ "github.com/kotovmak/go-admin/adapter/iris"
	"github.com/kotovmak/go-admin/modules/config"
	"github.com/kotovmak/go-admin/modules/language"
	"github.com/kotovmak/go-admin/plugins/admin/modules/table"

	// add mysql driver
	_ "github.com/kotovmak/go-admin/modules/db/drivers/mysql"
	// add postgresql driver
	_ "github.com/kotovmak/go-admin/modules/db/drivers/postgres"
	// add sqlite driver
	_ "github.com/kotovmak/go-admin/modules/db/drivers/sqlite"
	// add mssql driver
	_ "github.com/kotovmak/go-admin/modules/db/drivers/mssql"
	// add adminlte ui theme
	"github.com/GoAdminGroup/themes/adminlte"

	"github.com/kotovmak/go-admin/template"
	"github.com/kotovmak/go-admin/template/chartjs"

	"net/http"
	"os"

	"github.com/kataras/iris/v12"
	"github.com/kotovmak/go-admin/engine"
	"github.com/kotovmak/go-admin/plugins/admin"
	"github.com/kotovmak/go-admin/plugins/example"
	"github.com/kotovmak/go-admin/tests/tables"
)

func internalHandler() http.Handler {
	app := iris.New()

	eng := engine.Default()

	adminPlugin := admin.NewAdmin(tables.Generators)
	adminPlugin.AddGenerator("user", tables.GetUserTable)
	examplePlugin := example.NewExample()
	template.AddComp(chartjs.NewChart())

	if err := eng.AddConfigFromJSON(os.Args[len(os.Args)-1]).
		AddPlugins(adminPlugin, examplePlugin).Use(app); err != nil {
		panic(err)
	}

	eng.HTML("GET", "/admin", tables.GetContent)

	if err := app.Build(); err != nil {
		panic(err)
	}

	return app.Router
}

func NewHandler(dbs config.DatabaseList, gens table.GeneratorList) http.Handler {
	app := iris.New()

	eng := engine.Default()

	adminPlugin := admin.NewAdmin(gens)

	examplePlugin := example.NewExample()
	template.AddComp(chartjs.NewChart())

	if err := eng.AddConfig(&config.Config{
		Databases: dbs,
		UrlPrefix: "admin",
		Store: config.Store{
			Path:   "./uploads",
			Prefix: "uploads",
		},
		Language:    language.EN,
		IndexUrl:    "/",
		Debug:       true,
		ColorScheme: adminlte.ColorschemeSkinBlack,
	}).
		AddPlugins(adminPlugin, examplePlugin).Use(app); err != nil {
		panic(err)
	}

	eng.HTML("GET", "/admin", tables.GetContent)

	if err := app.Build(); err != nil {
		panic(err)
	}

	return app.Router
}
