package db

import (
	"STUOJ/internal/entity"
)

func autoMigrate() error {
	err := Db.AutoMigrate(&entity.Judgement{}, &entity.Language{}, &entity.Problem{}, &entity.History{}, &entity.ProblemTag{}, &entity.Solution{}, &entity.Submission{}, &entity.Tag{}, &entity.Testcase{}, &entity.User{})
	if err != nil {
		return err
	}
	return nil
}
