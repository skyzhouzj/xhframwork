package authorized_handler

import "github.com/skyzhouzj/xhframwork/internal/pkg/core"

func (h *handler) AddView() core.HandlerFunc {
	return func(c core.Context) {
		c.HTML("authorized_add", nil)
	}
}
