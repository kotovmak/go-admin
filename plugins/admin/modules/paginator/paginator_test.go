package paginator

import (
	"testing"

	_ "github.com/GoAdminGroup/themes/sword"
	"github.com/kotovmak/go-admin/modules/config"
	"github.com/kotovmak/go-admin/plugins/admin/modules/parameter"
)

func TestGet(t *testing.T) {
	config.Initialize(&config.Config{Theme: "sword"})
	Get(Config{
		Size:         105,
		Param:        parameter.BaseParam(),
		PageSizeList: []string{"10", "20", "50", "100"},
	})
}
