package menu_handler

import (
	"net/http"

	"github.com/skyzhouzj/xhframwork/internal/api/service/menu_service"
	"github.com/skyzhouzj/xhframwork/internal/pkg/code"
	"github.com/skyzhouzj/xhframwork/internal/pkg/core"
	"github.com/skyzhouzj/xhframwork/pkg/errno"
)

type detailRequest struct {
	Id string `uri:"id"` // HashID
}

type detailResponse struct {
	Id   int32  `json:"id"`   // 主键ID
	Pid  int32  `json:"pid"`  // 父类ID
	Name string `json:"name"` // 菜单名称
	Link string `json:"link"` // 链接地址
	Icon string `json:"icon"` // 图标
}

// Detail 菜单详情
// @Summary 菜单详情
// @Description 菜单详情
// @Tags API.menu
// @Accept json
// @Produce json
// @Param id path string true "hashId"
// @Success 200 {object} detailResponse
// @Failure 400 {object} code.Failure
// @Router /api/menu/{id} [get]
func (h *handler) Detail() core.HandlerFunc {
	return func(c core.Context) {
		req := new(detailRequest)
		res := new(detailResponse)
		if err := c.ShouldBindURI(req); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithErr(err),
			)
			return
		}

		ids, err := h.hashids.HashidsDecode(req.Id)
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.HashIdsDecodeError,
				code.Text(code.HashIdsDecodeError)).WithErr(err),
			)
			return
		}

		id := int32(ids[0])

		searchOneData := new(menu_service.SearchOneData)
		searchOneData.Id = id

		info, err := h.menuService.Detail(c, searchOneData)
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.MenuDetailError,
				code.Text(code.MenuDetailError)).WithErr(err),
			)
			return
		}

		res.Id = info.Id
		res.Pid = info.Pid
		res.Name = info.Name
		res.Link = info.Link
		res.Icon = info.Icon
		c.Payload(res)
	}
}
