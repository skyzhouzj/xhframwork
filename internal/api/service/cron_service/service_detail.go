package cron_service

import (
	"github.com/skyzhouzj/xhframwork/internal/api/repository/db_repo"
	"github.com/skyzhouzj/xhframwork/internal/api/repository/db_repo/cron_task_repo"
	"github.com/skyzhouzj/xhframwork/internal/pkg/core"
)

type SearchOneData struct {
	Id int32 // 任务ID
}

func (s *service) Detail(ctx core.Context, searchOneData *SearchOneData) (info *cron_task_repo.CronTask, err error) {
	qb := cron_task_repo.NewQueryBuilder()

	if searchOneData.Id != 0 {
		qb.WhereId(db_repo.EqualPredicate, searchOneData.Id)
	}

	info, err = qb.QueryOne(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}

	return
}
