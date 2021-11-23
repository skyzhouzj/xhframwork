package index_handler

import "github.com/skyzhouzj/xhframwork/internal/pkg/core"

func (h *handler) View() core.HandlerFunc {
	return func(c core.Context) {
		c.HTML("index", nil)
	}
}
