package cron_handler

import "github.com/skyzhouzj/xhframwork/internal/pkg/core"

func (h *handler) AddView() core.HandlerFunc {
	return func(ctx core.Context) {
		ctx.HTML("cron_task_add", nil)
	}
}
