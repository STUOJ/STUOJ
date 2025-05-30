package problem

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/pkg/errors"
	"STUOJ/pkg/utils"
	"slices"
)

func isPermission(problemMap map[string]any, reqUser request.ReqUser) error {
	// 检查是否是题目的创建者或管理员
	userIds, err := utils.StringToInt64Slice(string(problemMap["problem_user_id"].([]uint8)))
	if err != nil {
		return errors.ErrInternalServer.WithMessage("获取题目修改者id失败")
	}

	if reqUser.Role < 3 && !slices.Contains(userIds, reqUser.Id) { // 非管理员且非创建者
		return errors.ErrUnauthorized.WithMessage("无权限更新此题目")
	}
	return nil
}
