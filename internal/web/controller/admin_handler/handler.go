package admin_handler

import (
	"github.com/skyzhouzj/xhframwork/internal/pkg/cache"
	"github.com/skyzhouzj/xhframwork/internal/pkg/core"
	"github.com/skyzhouzj/xhframwork/internal/pkg/db"

	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	AddView() core.HandlerFunc
	ListView() core.HandlerFunc
	LoginView() core.HandlerFunc
	ModifyPasswordView() core.HandlerFunc
	ModifyInfoView() core.HandlerFunc
	MenuView() core.HandlerFunc
	MenuActionView() core.HandlerFunc
	AdminMenuView() core.HandlerFunc
}

type handler struct {
	db     db.Repo
	logger *zap.Logger
	cache  cache.Repo
}

func New(logger *zap.Logger, db db.Repo, cache cache.Repo) Handler {
	return &handler{
		logger: logger,
		cache:  cache,
		db:     db,
	}
}

func (h *handler) i() {}
