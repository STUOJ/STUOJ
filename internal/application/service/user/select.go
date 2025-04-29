package user

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/application/dto/response"
	"STUOJ/internal/domain/user"
	"STUOJ/internal/infrastructure/repository/querycontext"
	"STUOJ/internal/model"
	"STUOJ/pkg/errors"
	"STUOJ/pkg/utils"
	"log"
)

type UserPage struct {
	Users []response.UserData `json:"users"`
	model.Page
}

// SelectById 根据ID查询用户
func SelectById(id int64, reqUser model.ReqUser) (response.UserData, error) {
	var resp response.UserData

	// 查询
	qc := querycontext.UserQueryContext{}
	qc.Id.Add(id)
	qc.Field.SelectId().SelectUsername().SelectRole().SelectEmail().SelectAvatar().SelectSignature().SelectCreateTime().SelectUpdateTime()
	u0, _, err := user.Query.SelectOne(qc)
	if err != nil {
		return resp, err
	}

	resp = domain2Resp(u0)
	return resp, nil
}

// SelectByEmail 根据邮箱查询用户
func SelectByEmail(email string, reqUser model.ReqUser) (response.UserData, error) {
	var resp response.UserData
	qc := querycontext.UserQueryContext{}
	qc.Email.Set(email)
	qc.Field.SelectId().SelectUsername().SelectRole().SelectEmail().SelectAvatar().SelectSignature().SelectCreateTime().SelectUpdateTime()

	// 查询
	dmUser, _, err := user.Query.SelectOne(qc)
	if err != nil {
		return resp, err
	}

	resp = domain2Resp(dmUser)
	return resp, nil
}

// Select 查询所有用户
func Select(params request.QueryUserParams, reqUser model.ReqUser) (UserPage, error) {
	var resp UserPage

	// 查询
	qc := params2Query(params)
	qc.Field.SelectId().SelectUsername().SelectRole().SelectEmail().SelectAvatar().SelectSignature().SelectCreateTime().SelectUpdateTime()
	users, _, err := user.Query.Select(qc)
	if err != nil {
		return resp, err
	}

	for _, u := range users {
		respUser := domain2Resp(u)
		resp.Users = append(resp.Users, respUser)
	}

	resp.Page.Page = qc.Page.Page
	resp.Size = qc.Page.PageSize
	resp.Page.Total, err = Count(params)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// LoginByEmail 根据邮箱验证密码
func LoginByEmail(req request.UserLoginReq, reqUser model.ReqUser) (string, error) {
	qc := querycontext.UserQueryContext{}
	qc.Email.Set(req.Email)
	qc.Field.SelectId().SelectUsername().SelectRole().SelectEmail().SelectAvatar().SelectSignature().SelectCreateTime().SelectUpdateTime()

	// 查询
	u0, _, err := user.Query.SelectOne(qc)
	if err != nil {
		return "", err
	}

	// 验证密码
	err = u0.Password.VerifyHash(req.Password)
	if err != nil {
		log.Println(err)
		return "", errors.ErrUnauthorized.WithMessage("密码错误")
	}

	// 生成token
	token, err := utils.GenerateToken(uint64(u0.Id))
	if err != nil {
		log.Println(err)
		return "", errors.ErrInternalServer.WithMessage("生成Token失败")
	}

	return token, nil
}
