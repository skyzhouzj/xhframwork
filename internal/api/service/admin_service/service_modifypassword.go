package admin_service

import (
	"github.com/skyzhouzj/xhframwork/configs"
	"github.com/skyzhouzj/xhframwork/internal/api/repository/db_repo"
	"github.com/skyzhouzj/xhframwork/internal/api/repository/db_repo/admin_repo"
	"github.com/skyzhouzj/xhframwork/internal/pkg/cache"
	"github.com/skyzhouzj/xhframwork/internal/pkg/core"
	"github.com/skyzhouzj/xhframwork/internal/pkg/password"
)

func (s *service) ModifyPassword(ctx core.Context, id int32, newPassword string) (err error) {
	data := map[string]interface{}{
		"password":     password.GeneratePassword(newPassword),
		"updated_user": ctx.UserName(),
	}

	qb := admin_repo.NewQueryBuilder()
	qb.WhereId(db_repo.EqualPredicate, id)
	err = qb.Updates(s.db.GetDbW().WithContext(ctx.RequestContext()), data)
	if err != nil {
		return err
	}

	s.cache.Del(configs.RedisKeyPrefixLoginUser+password.GenerateLoginToken(id), cache.WithTrace(ctx.Trace()))
	return
}
