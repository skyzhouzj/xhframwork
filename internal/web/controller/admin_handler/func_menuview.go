package admin_handler

import "github.com/skyzhouzj/xhframwork/internal/pkg/core"

func (h *handler) MenuView() core.HandlerFunc {
	return func(c core.Context) {
		c.HTML("menu_view", nil)
	}
}
