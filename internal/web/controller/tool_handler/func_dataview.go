package tool_handler

import (
	"github.com/skyzhouzj/xhframwork/internal/pkg/core"
)

func (h *handler) DataView() core.HandlerFunc {
	return func(c core.Context) {
		c.HTML("tool_data", nil)
	}
}
