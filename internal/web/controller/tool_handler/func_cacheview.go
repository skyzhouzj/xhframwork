package tool_handler

import (
	"github.com/skyzhouzj/xhframwork/internal/pkg/core"
)

func (h *handler) CacheView() core.HandlerFunc {
	return func(c core.Context) {
		c.HTML("tool_cache", nil)
	}
}
