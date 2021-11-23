package install_handler

import (
	"runtime"

	"github.com/skyzhouzj/xhframwork/configs"
	"github.com/skyzhouzj/xhframwork/internal/pkg/core"
)

type viewResponse struct {
	Config    configs.Config
	GoVersion string
}

func (h *handler) View() core.HandlerFunc {
	return func(c core.Context) {
		obj := new(viewResponse)
		obj.Config = configs.Get()
		obj.GoVersion = runtime.Version()
		c.HTML("install_view", obj)
	}
}
