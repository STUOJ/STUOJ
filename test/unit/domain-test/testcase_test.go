package domain_test

import (
	"STUOJ/internal/domain/testcase"
	"STUOJ/internal/domain/testcase/valueobject"
	"math/rand"
	"strconv"
	"testing"
)

// 生成随机输入输出
func randomTestInput() valueobject.TestInput {
	return valueobject.NewTestInput("输入" + strconv.Itoa(rand.Intn(100000)))
}
func randomTestOutput() valueobject.TestOutput {
	return valueobject.NewTestOutput("输出" + strconv.Itoa(rand.Intn(100000)))
}

// 测试测试用例创建成功
func TestTestcaseCreate_Success(t *testing.T) {
	tc := &testcase.Testcase{
		ProblemId:  1,
		Serial:     uint16(rand.Intn(100)),
		TestInput:  randomTestInput(),
		TestOutput: randomTestOutput(),
	}
	id, err := tc.Create()
	if err != nil {
		t.Fatalf("创建测试用例失败: %v", err)
	}
	if id == 0 {
		t.Fatalf("创建测试用例返回的ID无效")
	}
}

// 测试输入为空时创建失败
func TestTestcaseCreate_EmptyInput(t *testing.T) {
	tc := &testcase.Testcase{
		ProblemId:  1,
		Serial:     1,
		TestInput:  valueobject.NewTestInput(""),
		TestOutput: randomTestOutput(),
	}
	_, err := tc.Create()
	if err == nil {
		t.Fatalf("输入为空时应创建失败")
	}
}

// 测试测试用例更新成功
func TestTestcaseUpdate_Success(t *testing.T) {
	tc := &testcase.Testcase{
		ProblemId:  1,
		Serial:     uint16(rand.Intn(100)),
		TestInput:  randomTestInput(),
		TestOutput: randomTestOutput(),
	}
	id, err := tc.Create()
	if err != nil {
		t.Fatalf("创建测试用例失败: %v", err)
	}
	tc.Id = id
	tc.TestOutput = valueobject.NewTestOutput("更新后的输出")
	err = tc.Update()
	if err != nil {
		t.Fatalf("更新测试用例失败: %v", err)
	}
}

// 测试更新不存在的测试用例
func TestTestcaseUpdate_NotFound(t *testing.T) {
	tc := &testcase.Testcase{
		Id:         99999999,
		ProblemId:  1,
		Serial:     1,
		TestInput:  randomTestInput(),
		TestOutput: randomTestOutput(),
	}
	err := tc.Update()
	if err == nil {
		t.Fatalf("更新不存在的测试用例应失败")
	}
}

// 测试测试用例删除成功
func TestTestcaseDelete_Success(t *testing.T) {
	tc := &testcase.Testcase{
		ProblemId:  1,
		Serial:     uint16(rand.Intn(100)),
		TestInput:  randomTestInput(),
		TestOutput: randomTestOutput(),
	}
	id, err := tc.Create()
	if err != nil {
		t.Fatalf("创建测试用例失败: %v", err)
	}
	tc.Id = id
	err = tc.Delete()
	if err != nil {
		t.Fatalf("删除测试用例失败: %v", err)
	}
}

// 测试删除不存在的测试用例
func TestTestcaseDelete_NotFound(t *testing.T) {
	tc := &testcase.Testcase{
		Id:         99999999,
		ProblemId:  1,
		Serial:     1,
		TestInput:  randomTestInput(),
		TestOutput: randomTestOutput(),
	}
	err := tc.Delete()
	if err == nil {
		t.Fatalf("删除不存在的测试用例应失败")
	}
}
