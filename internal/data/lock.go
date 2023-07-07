package data

import (
	"context"
	"time"

	"github.com/bsm/redislock"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/star-table/go-common/pkg/errors"
	"github.com/star-table/go-table/internal/biz"
	comomPb "github.com/star-table/interface/golang/common/v1"
)

type lockRepo struct {
	data *Data
	log  *log.Helper
}

// NewLockRepo .
func NewLockRepo(data *Data, logger log.Logger) biz.LockRepo {
	return &lockRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (l *lockRepo) TryGetDistributedLock(ctx context.Context, lockKey string, obtainDuration ...time.Duration) (func(), error) {
	locker := redislock.New(l.data.redisCli)
	d := 4 * time.Second
	if len(obtainDuration) > 0 {
		d = obtainDuration[0]
	}
	// Try to obtain lock.
	lock, err := locker.Obtain(ctx, lockKey, d, nil)
	if err == redislock.ErrNotObtained {
		return nil, errors.WithStackLevel(comomPb.ErrorDuplicateOperation("Duplicate Operation"), log.LevelInfo)
	} else if err != nil {
		return nil, errors.Wrapf(err, "[TryGetDistributedLock] lockKey:%s", lockKey)
	}

	return func() { _ = lock.Release(ctx) }, nil
}
