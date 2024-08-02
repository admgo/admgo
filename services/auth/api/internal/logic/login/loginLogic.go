package login

import (
	"context"
	"github.com/admgo/admgo/pkg/random"
	"github.com/admgo/admgo/services/user/rpc/user"
	"net/http"
	"strconv"

	"github.com/admgo/admgo/services/auth/api/internal/svc"
	"github.com/admgo/admgo/services/auth/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx            context.Context
	svcCtx         *svc.ServiceContext
	sessionContent map[string]string
	sessionID      string
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger:         logx.WithContext(ctx),
		ctx:            ctx,
		svcCtx:         svcCtx,
		sessionContent: make(map[string]string),
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginRsp, err error) {
	rsp, err := l.svcCtx.UserRPC.FindSingleUserByUsernameAndPassword(l.ctx, &user.FindSingleUserByUsernameAndPasswordRequest{
		UserName: req.UserName,
		Password: req.PassWord,
	})
	if err != nil {
		return nil, err
	}

	l.writeToSessionContent(rsp)
	l.sessionID = random.GenerateSessionID()
	err = l.svcCtx.Redis.HmsetCtx(l.ctx, "/"+l.svcCtx.Config.Name+"/rest/user_token/"+l.sessionID, l.sessionContent)
	if err != nil {
		return nil, err
	}

	return &types.LoginRsp{
		PassWord: rsp.Phone,
	}, nil
}

func (l *LoginLogic) writeToSessionContent(u *user.FindSingleUserByUsernameAndPasswordResponse) {
	l.sessionContent["user_id"] = strconv.FormatUint(uint64(u.UserID), 10)
	l.sessionContent["user_name"] = u.UserName
}

func (l *LoginLogic) WriteUserTokenToHeader(w http.ResponseWriter) {
	w.Header().Set("Set-Cookie", "user_token="+l.sessionID)
}
