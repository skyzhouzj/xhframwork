package admin_handler

import (
	"net/http"

	"github.com/skyzhouzj/xhframwork/internal/pkg/code"
	"github.com/skyzhouzj/xhframwork/internal/pkg/core"
	"github.com/skyzhouzj/xhframwork/pkg/errno"
)

type adminMenuViewRequest struct {
	Id string `uri:"id"` // 主键ID
}

type adminMenuViewResponse struct {
	HashID string `json:"hash_id"` // hashID
}

func (h *handler) AdminMenuView() core.HandlerFunc {
	return func(c core.Context) {
		req := new(adminMenuViewRequest)
		if err := c.ShouldBindURI(req); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithErr(err),
			)
			return
		}

		obj := new(adminMenuViewResponse)
		obj.HashID = req.Id

		c.HTML("admin_menu", obj)
	}
}
