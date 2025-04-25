package problem

import (
	"STUOJ/internal/errors"
	"STUOJ/internal/model"
	"STUOJ/utils"
	"slices"
)

func isPermission(problemMap map[string]any, reqUser model.ReqUser) error {
	// 检查是否是题目的创建者或管理员
	userIds, err := utils.StringToInt64Slice(problemMap["problem_user_id"].(string))
	if err != nil {
		return errors.ErrInternalServer.WithMessage("获取题目修改者id失败")
	}

	if reqUser.Role < 3 && !slices.Contains(userIds, reqUser.Id) { // 非管理员且非创建者
		return errors.ErrUnauthorized.WithMessage("无权限更新此题目")
	}
	return nil
}
