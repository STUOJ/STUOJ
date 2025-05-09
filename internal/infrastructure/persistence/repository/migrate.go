package repository

import (
	entity "STUOJ/internal/infrastructure/persistence/entity"
)

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
		&entity.Contest{},
		&entity.ContestProblem{},
		&entity.ContestUser{},
		&entity.Team{},
		&entity.TeamUser{},
		&entity.TeamSubmission{},
	)
	if err != nil {
		return err
	}
	return nil
}
