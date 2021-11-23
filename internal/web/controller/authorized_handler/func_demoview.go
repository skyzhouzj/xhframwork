package authorized_handler

import "github.com/skyzhouzj/xhframwork/internal/pkg/core"

func (h *handler) DemoView() core.HandlerFunc {
	return func(c core.Context) {
		c.HTML("authorized_demo", nil)
	}
}
