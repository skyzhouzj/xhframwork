package config_handler

import (
	"github.com/skyzhouzj/xhframwork/configs"
	"github.com/skyzhouzj/xhframwork/internal/pkg/core"
)

func (h *handler) EmailView() core.HandlerFunc {
	return func(c core.Context) {
		c.HTML("config_email", configs.Get())
	}
}
