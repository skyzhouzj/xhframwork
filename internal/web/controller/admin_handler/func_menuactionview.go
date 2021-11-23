package admin_handler

import (
	"net/http"

	"github.com/skyzhouzj/xhframwork/internal/pkg/code"
	"github.com/skyzhouzj/xhframwork/internal/pkg/core"
	"github.com/skyzhouzj/xhframwork/pkg/errno"
)

type menuActionViewRequest struct {
	Id string `uri:"id"` // 主键ID
}

type menuActionViewResponse struct {
	HashID string `json:"hash_id"` // hashID
}

func (h *handler) MenuActionView() core.HandlerFunc {
	return func(c core.Context) {
		req := new(menuActionViewRequest)
		if err := c.ShouldBindURI(req); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithErr(err),
			)
			return
		}

		obj := new(menuActionViewResponse)
		obj.HashID = req.Id

		c.HTML("menu_action", obj)
	}
}
