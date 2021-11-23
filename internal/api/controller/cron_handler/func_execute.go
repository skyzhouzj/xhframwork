package cron_handler

import (
	"net/http"

	"github.com/skyzhouzj/xhframwork/internal/pkg/code"
	"github.com/skyzhouzj/xhframwork/internal/pkg/core"
	"github.com/skyzhouzj/xhframwork/internal/pkg/validation"
	"github.com/skyzhouzj/xhframwork/pkg/errno"

	"github.com/spf13/cast"
)

type executeRequest struct {
	Id string `uri:"id"` // HashID
}

type executeResponse struct {
	Id int `json:"id"` // ID
}

// Execute 手动执行单条任务
// @Summary 手动执行单条任务
// @Description 手动执行单条任务
// @Tags API.cron
// @Accept json
// @Produce json
// @Param id path string true "hashId"
// @Success 200 {object} detailResponse
// @Failure 400 {object} code.Failure
// @Router /api/cron/:id [patch]
func (h *handler) Execute() core.HandlerFunc {
	return func(ctx core.Context) {
		req := new(executeRequest)
		res := new(executeResponse)
		if err := ctx.ShouldBindURI(req); err != nil {
			ctx.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ParamBindError,
				validation.Error(err)).WithErr(err),
			)
			return
		}

		ids, err := h.hashids.HashidsDecode(req.Id)
		if err != nil {
			ctx.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.HashIdsDecodeError,
				code.Text(code.HashIdsDecodeError)).WithErr(err),
			)
			return
		}

		err = h.cronService.Execute(ctx, cast.ToInt32(ids[0]))
		if err != nil {
			ctx.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.CronExecuteError,
				code.Text(code.CronExecuteError)).WithErr(err),
			)
			return
		}

		res.Id = ids[0]
		ctx.Payload(res)
	}
}
