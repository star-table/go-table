package data

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/star-table/go-table/internal/biz"
	"github.com/star-table/go-table/internal/conf"
	"github.com/star-table/go-table/internal/data/facade/vo/usercentervo"
)

type userCenterRepo struct {
	uc  *http.Client
	log *log.Helper
}

func NewUserCenterRepo(conf *conf.Data, r registry.Discovery, logger log.Logger) biz.UserCenterRepo {
	conn, err := getHttpConn(conf.Facade.UsercenterServer, r, logger)
	if err != nil {
		log.Error(err)
		return nil
	}
	return &userCenterRepo{
		uc:  conn,
		log: log.NewHelper(logger),
	}
}

func (u *userCenterRepo) GetMemberSimpleInfo(ctx context.Context, req *usercentervo.GetMemberSimpleInfoReq) (*usercentervo.GetMemberSimpleInfoResp, error) {
	respVo := usercentervo.GetMemberSimpleInfoResp{}
	path := fmt.Sprintf("usercenter/inner/api/v1/user/getMemberSimpleInfo")
	err := u.uc.Invoke(ctx, "POST", path, req, &respVo)
	if err != nil {
		return nil, err
	}
	return &respVo, nil
}
