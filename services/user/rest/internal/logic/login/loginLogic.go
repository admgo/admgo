package login

import (
	"context"
	"github.com/admgo/admgo/pkg/errorx"
	xparse "github.com/admgo/admgo/pkg/parse"
	"github.com/admgo/admgo/pkg/random"
	"github.com/admgo/admgo/services/user/rpc/user"
	"github.com/zeromicro/go-zero/core/logc"
	"net/http"
	"strconv"
	"time"

	"github.com/admgo/admgo/services/user/rest/internal/svc"
	"github.com/admgo/admgo/services/user/rest/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx            context.Context
	svcCtx         *svc.ServiceContext
	sessionContent map[string]string
	sessionID      string
	w              http.ResponseWriter
	r              *http.Request
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext, rw http.ResponseWriter, rr *http.Request) *LoginLogic {
	return &LoginLogic{
		Logger:         logx.WithContext(ctx),
		ctx:            ctx,
		svcCtx:         svcCtx,
		sessionContent: make(map[string]string),
		w:              rw,
		r:              rr,
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
	// 写入session到缓存
	err = l.svcCtx.Redis.HmsetCtx(l.ctx, "/"+l.svcCtx.Config.Name+"/rest/session/"+l.sessionID, l.sessionContent)
	if err != nil {
		logc.Error(l.ctx, err)
		return nil, errorx.ServerErrForRedis
	}
	err = l.svcCtx.Redis.ExpireCtx(l.ctx, "/"+l.svcCtx.Config.Name+"/rest/session/"+l.sessionID, types.SessionExpireTime)
	if err != nil {
		logc.Error(l.ctx, err)
		return nil, errorx.ServerErrForRedis
	}
	// 写入该用户所有的session到缓存
	_, err = l.svcCtx.Redis.SaddCtx(l.ctx, "/"+l.svcCtx.Config.Name+"/rest/sessions/"+strconv.Itoa(int(rsp.UserID)), l.sessionID)
	if err != nil {
		logc.Error(l.ctx, err)
		return nil, errorx.ServerErrForRedis
	}
	err = l.svcCtx.Redis.ExpireCtx(l.ctx, "/"+l.svcCtx.Config.Name+"/rest/sessions/"+strconv.Itoa(int(rsp.UserID)), types.SessionExpireTime)
	if err != nil {
		logc.Error(l.ctx, err)
		return nil, errorx.ServerErrForRedis
	}

	l.writeUserTokenToHeader(l.w)

	return &types.LoginRsp{}, nil
}

func (l *LoginLogic) writeToSessionContent(u *user.FindSingleUserByUsernameAndPasswordResponse) {
	l.sessionContent["user_id"] = strconv.Itoa(int(u.UserID))
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05")
	l.sessionContent["login_at"] = formattedTime
	l.sessionContent["host"] = l.r.RequestURI
	ua := xparse.Parse(l.r.Header.Get("User-Agent"))
	l.sessionContent["os_version"] = ua.OSVersion
	//l.sessionContent["os_version"] = ua.IsLinux()
	l.sessionContent["os"] = ua.OS
}

func (l *LoginLogic) writeUserTokenToHeader(w http.ResponseWriter) {
	cookie := http.Cookie{
		Name:     "access_token",
		Value:    l.sessionID,
		HttpOnly: false,
		Path:     "/",
		MaxAge:   types.SessionExpireTime,
	}
	w.Header().Set("Set-Cookie", cookie.String())
}
