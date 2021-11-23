package cron_handler

import (
	"net/http"

	"github.com/skyzhouzj/xhframwork/internal/pkg/code"
	"github.com/skyzhouzj/xhframwork/internal/pkg/core"
	"github.com/skyzhouzj/xhframwork/pkg/errno"
)

type editViewRequest struct {
	Id string `uri:"id"` // 主键ID
}

type editViewResponse struct {
	HashID string `json:"hash_id"` // hashID
}

func (h *handler) EditView() core.HandlerFunc {
	return func(ctx core.Context) {
		req := new(editViewRequest)
		if err := ctx.ShouldBindURI(req); err != nil {
			ctx.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithErr(err),
			)
			return
		}

		obj := new(editViewResponse)
		obj.HashID = req.Id
		ctx.HTML("cron_task_edit", obj)
	}
}
