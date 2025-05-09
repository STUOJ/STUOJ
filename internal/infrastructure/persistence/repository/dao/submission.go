package dao

import (
	"STUOJ/internal/infrastructure/persistence/repository"
)

func (_SubmissionStore) SelectACUsers(pid, size int64) ([]map[string]any, error) {
	var res []map[string]any
	tx := repository.Db.Raw("SELECT * FROM (SELECT tbl_submission.create_time,tbl_user.id,tbl_user.username,tbl_user.role,tbl_user.avatar,ROW_NUMBER() OVER (PARTITION BY tbl_submission.user_id ORDER BY tbl_submission.create_time ASC) AS rn FROM `tbl_submission` LEFT JOIN tbl_problem ON tbl_submission.problem_id=tbl_problem.id LEFT JOIN tbl_user ON tbl_submission.user_id=tbl_user.id WHERE `tbl_submission`.`status`=3 AND tbl_submission.problem_id=? AND NOT EXISTS (SELECT 1 FROM tbl_history WHERE tbl_history.user_id=tbl_submission.user_id AND tbl_history.problem_id=tbl_submission.problem_id)) AS ranked_submissions WHERE rn=1 ORDER BY create_time ASC LIMIT ?", pid, size)
	err := tx.Scan(&res).Error
	return res, err
}
