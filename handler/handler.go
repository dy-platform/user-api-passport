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
		Code:                 req.Code,
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
		Gender:               "",
		AvatarUrl:            "dayan.xxx.xxx",
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
		Code: uint32(base.CODE_SUCESS),
	}

	return nil
}

func (h *Handler) MobileSignIn(ctx context.Context, req *api.MobileSignInReq, rsp *api.MobileSignInResp) error {

	return nil
}

func (h *Handler) EmailSignIn(ctx context.Context, req *api.EmailSignInReq, rsp *api.EmailSignInResp) error {

	return nil
}

func (h *Handler) UserNameSignIn(ctx context.Context, req *api.UserNameSignInReq, rsp *api.UserNameSignInResp) error {

	return nil
}

func (h *Handler) TokenSignIn(ctx context.Context, req *api.TokenSignInReq, rsp *api.TokenSignInResp) error {

	return nil
}

func (h *Handler) SignOut(ctx context.Context, req *api.SignOutReq, rsp *api.SignOutResp) error {

	return nil
}

func (h *Handler) ChangePassword(ctx context.Context, req *api.ChangePasswordReq, rsp *api.ChangePasswordResp) error {

	return nil
}


func (h *Handler) ResetPassword(ctx context.Context, req *api.ResetPasswordReq, rsp *api.ResetPasswordResp) error {

	return nil
}

func (h *Handler) BindMobile(ctx context.Context, req *api.BindMobileReq, rsp *api.BindMobileResp) error {

	return nil
}

func (h *Handler) UnbindMobile(ctx context.Context, req *api.UnbindMobileReq, rsp *api.UnbindMobileResp) error {

	return nil
}
