package db

import "STUOJ/internal/entity"

func autoMigrate() error {
	err := Db.AutoMigrate(
		&entity.Testcase{},
		&entity.Problem{},
		&entity.Blog{},
		&entity.User{},
		&entity.Comment{},
		&entity.Judgement{},
		&entity.Language{},
		&entity.History{},
		&entity.ProblemTag{},
		&entity.Solution{},
		&entity.Submission{},
		&entity.Tag{},
		&entity.Collection{},
		&entity.CollectionUser{},
		&entity.CollectionProblem{},
	)
	if err != nil {
		return err
	}
	return nil
}
