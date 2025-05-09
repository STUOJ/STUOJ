package dao_test

import (
	"STUOJ/internal/infrastructure/persistence/repository/dao"
	"testing"
)

func TestStringToStatus(t *testing.T) {
	str := "2,4"
	res, err := dao.StringToBlogStatusSlice(str)
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}
