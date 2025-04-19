package user

import (
	"STUOJ/internal/app/dto/request"
	"STUOJ/internal/app/dto/response"
	"STUOJ/internal/db/querycontext"
	"STUOJ/internal/domain/user"
	"STUOJ/internal/errors"
	"STUOJ/internal/model"
	"STUOJ/utils"
	"log"
)

type UserPage struct {
	Users []response.UserData `json:"users"`
	model.Page
}

// SelectById 根据ID查询用户
func SelectById(id uint64, reqUser model.ReqUser) (response.UserData, error) {
	var resp response.UserData
	qc := querycontext.UserQueryContext{}
	qc.Id.Add(id)
	qc.Field.SelectId().SelectUsername().SelectRole().SelectEmail().SelectAvatar().SelectSignature().SelectCreateTime().SelectUpdateTime()

	// 查询
	dmUser, _, err := user.Query.SelectOne(qc)
	if err != nil {
		return resp, err
	}

	resp = domain2Resp(dmUser)
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

// 查询所有用户
func Select(params request.QueryUserParams, reqUser model.ReqUser) (UserPage, error) {
	var resp UserPage
	qc := params2Query(params)
	qc.Field.SelectId().SelectUsername().SelectRole().SelectEmail().SelectAvatar().SelectSignature().SelectCreateTime().SelectUpdateTime()

	// 查询
	users, _, err := user.Query.Select(qc)
	if err != nil {
		return UserPage{}, err
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

// 根据邮箱验证密码
func LoginByEmail(req request.UserLoginReq, reqUser model.ReqUser) (string, error) {
	qc := querycontext.UserQueryContext{}
	qc.Email.Set(req.Email)
	qc.Field.SelectId().SelectUsername().SelectRole().SelectEmail().SelectAvatar().SelectSignature().SelectCreateTime().SelectUpdateTime()

	// 查询
	dmUser, _, err := user.Query.SelectOne(qc)
	if err != nil {
		return "", err
	}

	// 验证密码
	err = dmUser.Password.VerifyHash(req.Password)
	if err != nil {
		log.Println(err)
		return "", errors.ErrUnauthorized.WithMessage("密码错误")
	}

	// 生成token
	token, err := utils.GenerateToken(dmUser.Id)
	if err != nil {
		log.Println(err)
		return "", errors.ErrInternalServer.WithMessage("生成Token失败")
	}

	return token, nil
}
