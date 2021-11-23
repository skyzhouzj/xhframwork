package menu_handler

import (
	"net/http"

	"github.com/skyzhouzj/xhframwork/internal/pkg/code"
	"github.com/skyzhouzj/xhframwork/internal/pkg/core"
	"github.com/skyzhouzj/xhframwork/pkg/errno"
)

type updateSortRequest struct {
	Id   string `form:"id"`   // HashId
	Sort int32  `form:"sort"` // 排序
}

type updateSortResponse struct {
	Id int32 `json:"id"` // 主键ID
}

// UpdateSort 更新菜单排序
// @Summary 更新菜单排序
// @Description 更新菜单排序
// @Tags API.menu
// @Accept multipart/form-data
// @Produce json
// @Param id formData string true "Hashid"
// @Param sort formData int true "排序"
// @Success 200 {object} updateSortResponse
// @Failure 400 {object} code.Failure
// @Router /api/menu/sort [patch]
func (h *handler) UpdateSort() core.HandlerFunc {
	return func(c core.Context) {
		req := new(updateSortRequest)
		res := new(updateSortResponse)
		if err := c.ShouldBindForm(req); err != nil {
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

		err = h.menuService.UpdateSort(c, id, req.Sort)
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.MenuUpdateError,
				code.Text(code.MenuUpdateError)).WithErr(err),
			)
			return
		}

		res.Id = id
		c.Payload(res)
	}
}
