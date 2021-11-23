package tool_handler

import (
	"net/http"

	"github.com/skyzhouzj/xhframwork/internal/pkg/cache"
	"github.com/skyzhouzj/xhframwork/internal/pkg/code"
	"github.com/skyzhouzj/xhframwork/internal/pkg/core"
	"github.com/skyzhouzj/xhframwork/pkg/errno"
)

type searchCacheRequest struct {
	RedisKey string `form:"redis_key"` // Redis Key
}

type searchCacheResponse struct {
	Val string `json:"val"` // 查询后的值
	TTL string `json:"ttl"` // 过期时间
}

// SearchCache 查询缓存
// @Summary 查询缓存
// @Description 查询缓存
// @Tags API.tool
// @Accept multipart/form-data
// @Produce json
// @Param redis_key formData string true "Redis Key"
// @Success 200 {object} searchCacheResponse
// @Failure 400 {object} code.Failure
// @Router /api/tool/cache/search [post]
func (h *handler) SearchCache() core.HandlerFunc {
	return func(c core.Context) {
		req := new(searchCacheRequest)
		res := new(searchCacheResponse)
		if err := c.ShouldBindForm(req); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithErr(err),
			)
			return
		}

		if b := h.cache.Exists(req.RedisKey); b != true {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.CacheNotExist,
				code.Text(code.CacheNotExist)),
			)
			return
		}

		val, err := h.cache.Get(req.RedisKey, cache.WithTrace(c.Trace()))
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.CacheGetError,
				code.Text(code.CacheGetError)).WithErr(err),
			)
			return
		}

		ttl, err := h.cache.TTL(req.RedisKey)
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.CacheGetError,
				code.Text(code.CacheGetError)).WithErr(err),
			)
			return
		}

		res.Val = val
		res.TTL = ttl.String()
		c.Payload(res)
	}
}
