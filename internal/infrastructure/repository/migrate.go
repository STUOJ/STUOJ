package repository

import (
	entity2 "STUOJ/internal/infrastructure/repository/entity"
)

func autoMigrate() error {
	err := Db.AutoMigrate(
		&entity2.Testcase{},
		&entity2.Problem{},
		&entity2.Blog{},
		&entity2.User{},
		&entity2.Comment{},
		&entity2.Judgement{},
		&entity2.Language{},
		&entity2.History{},
		&entity2.ProblemTag{},
		&entity2.Solution{},
		&entity2.Submission{},
		&entity2.Tag{},
		&entity2.Collection{},
		&entity2.CollectionUser{},
		&entity2.CollectionProblem{},
		&entity2.Contest{},
		&entity2.Team{},
		&entity2.TeamUser{},
		&entity2.TeamSubmission{},
	)
	if err != nil {
		return err
	}
	return nil
}
