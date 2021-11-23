package tool_handler

import (
	"github.com/skyzhouzj/xhframwork/configs"
	"github.com/skyzhouzj/xhframwork/internal/pkg/core"
)

func (h *handler) HashIdsView() core.HandlerFunc {
	return func(c core.Context) {
		c.HTML("tool_hashids", configs.Get())
	}
}
