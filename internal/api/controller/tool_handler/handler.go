package tool_handler

import (
	"github.com/skyzhouzj/xhframwork/configs"
	"github.com/skyzhouzj/xhframwork/internal/pkg/cache"
	"github.com/skyzhouzj/xhframwork/internal/pkg/core"
	"github.com/skyzhouzj/xhframwork/internal/pkg/db"
	"github.com/skyzhouzj/xhframwork/pkg/hash"

	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// HashIdsEncode HashIds 加密
	// @Tags API.tool
	// @Router /api/tool/hashids/encode/{id} [get]
	HashIdsEncode() core.HandlerFunc

	// HashIdsDecode HashIds 解密
	// @Tags API.tool
	// @Router /api/tool/hashids/decode/{id} [get]
	HashIdsDecode() core.HandlerFunc

	// SearchCache 查询缓存
	// @Tags API.tool
	// @Router /api/tool/cache/search [post]
	SearchCache() core.HandlerFunc

	// ClearCache 清空缓存
	// @Tags API.tool
	// @Router /api/tool/cache/clear [patch]
	ClearCache() core.HandlerFunc

	// Dbs 查询 DB
	// @Tags API.tool
	// @Router /api/tool/data/dbs [get]
	Dbs() core.HandlerFunc

	// Tables 查询 Table
	// @Tags API.tool
	// @Router /api/tool/data/tables [post]
	Tables() core.HandlerFunc

	// SearchMySQL 执行 SQL 语句
	// @Tags API.tool
	// @Router /api/tool/data/mysql [post]
	SearchMySQL() core.HandlerFunc
}

type handler struct {
	logger  *zap.Logger
	db      db.Repo
	cache   cache.Repo
	hashids hash.Hash
}

func New(logger *zap.Logger, db db.Repo, cache cache.Repo) Handler {
	return &handler{
		logger:  logger,
		db:      db,
		cache:   cache,
		hashids: hash.New(configs.Get().HashIds.Secret, configs.Get().HashIds.Length),
	}
}

func (h *handler) i() {}