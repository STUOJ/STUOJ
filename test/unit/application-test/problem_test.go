package application_test

import (
	"STUOJ/cmd/bootstrap"
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/application/service/problem"
	"STUOJ/internal/infrastructure/persistence/entity"
	"testing"
)

func TestDailyProblem(t *testing.T) {
	bootstrap.InitConfig()
	bootstrap.InitDatabase()

	user := request.ReqUser{Id: 2, Role: entity.RoleUser}

	// First query
	problem1, err := problem.SelectDailyProblem(user)
	if err != nil {
		t.Fatal(err)
	}
	if problem1.Id == 0 {
		t.Fatal("expected a valid problem")
	}

	// Second query with same user
	problem2, err := problem.SelectDailyProblem(user)
	if err != nil {
		t.Fatal(err)
	}
	if problem2.Id == 0 {
		t.Fatal("expected a valid problem")
	}

	// Verify same problem is returned for same user
	if problem1.Id != problem2.Id {
		t.Errorf("expected same problem id, got %d and %d", problem1.Id, problem2.Id)
	}

	// Verify score indicates passed submission (assuming 60 is passing threshold)
	if problem1.UserScore >= 100 {
		t.Errorf("expected passing score (>=60), got %d", problem1.UserScore)
	}
}
