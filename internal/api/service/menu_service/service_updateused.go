package menu_service

import (
	"github.com/skyzhouzj/xhframwork/internal/api/repository/db_repo"
	"github.com/skyzhouzj/xhframwork/internal/api/repository/db_repo/menu_repo"
	"github.com/skyzhouzj/xhframwork/internal/pkg/core"
)

func (s *service) UpdateUsed(ctx core.Context, id int32, used int32) (err error) {
	data := map[string]interface{}{
		"is_used":      used,
		"updated_user": ctx.UserName(),
	}

	qb := menu_repo.NewQueryBuilder()
	qb.WhereId(db_repo.EqualPredicate, id)
	err = qb.Updates(s.db.GetDbW().WithContext(ctx.RequestContext()), data)
	if err != nil {
		return err
	}

	return
}
