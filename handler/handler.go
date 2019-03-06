package handler

import (
	"context"
	"github.com/dy-gopkg/kit"
	"github.com/dy-platform/user-api-passport/idl"
	api "github.com/dy-platform/user-api-passport/idl/platform/user/api-passport"
	info "github.com/dy-platform/user-api-passport/idl/platform/user/srv-info"
	passport "github.com/dy-platform/user-api-passport/idl/platform/user/srv-passport"
	"github.com/sirupsen/logrus"
)

type Handler struct {

}

func (h *Handler) SignUp(ctx context.Context, req *api.SignUpReq, rsp *api.SignUpResp) error {
	cl := passport.NewPassportService("platform.user.srv.passport", kit.Client())
	req1 := &passport.SignUpReq{
		DeviceID:             req.DeviceID,
		Name:                 req.Name,
		Mobile:               req.Mobile,
		Email:                req.Email,
		Password:             req.Password,
		MsgCode:              req.MsgCode,
		ActiveType:           0,
		AppID:                req.AppID,
		Version:              req.Version,
		ExtraInfo:            req.ExtraInfo,
	}
	rsp1, err := cl.SignUp(ctx, req1)
	if err != nil {
		logrus.Warnf("platform.user.srv.passport SingUp error:%v", err)
		return err
	}

	req2 := &info.CreateUserReq{
		UserId:               rsp1.UserID,
		NickName:             "xxx", // 随机生成
		Gender:               base.Gender_Unkonw,
		AvatarUrl:            "",
		UserType:             0,
	}

	cl2 := info.NewUserInfoService("platform.user.srv.info", kit.Client())
	_, err = cl2.CreateUser(ctx, req2)
	if err != nil {
		logrus.Warnf("platform.user.srv.info CreateUser error:%v", err)
		return err
	}

	rsp.UserID = rsp1.UserID
	rsp.Mobile = rsp1.Mobile
	rsp.BaseResp = &base.Resp{
		Code: int32(base.CODE_OK),
	}

	return nil
}

func (h *Handler) WeChatSignIn(ctx context.Context, req *api.WeChatSignInReq, rsp *api.WeChatSignInResp) error {
	rsp.BaseResp = &base.Resp{
		Code: int32(base.CODE_OK),
	}
	cl := passport.NewPassportService("platform.user.srv.passport", kit.Client())
	req1 := &passport.WeChatSignInReq{
		DeviceID:             req.DeviceID,
		Code:                 req.Code,
		AppID:                req.AppID,
	}

	rsp1, err := cl.WeChatSignIn(ctx, req1)
	if err != nil {
		logrus.Warnf("platform.user.srv.passport WeChatSignIn error:%v", err)
		return err
	}

	if rsp1.BaseResp.Code != int32(base.CODE_OK) {
		rsp.BaseResp.Code = rsp1.BaseResp.Code
		rsp.BaseResp.Msg = rsp1.BaseResp.Msg
		return nil
	}

	//第一次登陆, 添加用户信息
	if rsp1.IsFirstSignIn {
		cl2 := info.NewUserInfoService("platform.user.srv.info", kit.Client())
		req2 := &info.CreateUserReq{
			UserId:               rsp1.UserID,
			NickName:             "",
			Gender:               base.Gender_Unkonw,
			AvatarUrl:            "",
			UserType:             0,
		}
		_, err = cl2.CreateUser(ctx, req2)
		if err != nil {
			logrus.Warnf("platform.user.srv.info CreateUser error:%v", err)
			return err
		}
	}

	rsp.UserID = rsp1.UserID
	return nil
}