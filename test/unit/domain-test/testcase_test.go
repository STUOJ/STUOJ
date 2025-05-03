package domain_test

import (
	"STUOJ/internal/domain/testcase"
	"STUOJ/internal/domain/testcase/valueobject"
	"math/rand"
	"strconv"
	"testing"
)

// 生成随机输入输出
func randomTestInput() string {
	return "输入" + strconv.Itoa(rand.Intn(100000))
}
func randomTestOutput() string {
	return "输出" + strconv.Itoa(rand.Intn(100000))
}

// 测试测试用例创建成功
func TestTestcaseCreate_Success(t *testing.T) {
	tc := testcase.NewTestcase(
		testcase.WithProblemId(1),
		testcase.WithSerial(uint16(rand.Intn(100))),
		testcase.WithTestInput(randomTestInput()),
		testcase.WithTestOutput(randomTestOutput()),
	)
	id, err := tc.Create()
	if err != nil {
		t.Fatalf("创建测试用例失败: %v", err)
	}
	if id == 0 {
		t.Fatalf("创建测试用例返回的ID无效")
	}
}

// 测试测试用例更新成功
func TestTestcaseUpdate_Success(t *testing.T) {
	tc := testcase.NewTestcase(
		testcase.WithProblemId(1),
		testcase.WithSerial(uint16(rand.Intn(100))),
		testcase.WithTestInput(randomTestInput()),
		testcase.WithTestOutput(randomTestOutput()),
	)
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
	tc := testcase.NewTestcase(
		testcase.WithId(99999999),
		testcase.WithProblemId(1),
		testcase.WithSerial(1),
		testcase.WithTestInput(randomTestInput()),
		testcase.WithTestOutput(randomTestOutput()),
	)
	err := tc.Update()
	if err == nil {
		t.Fatalf("更新不存在的测试用例应失败")
	}
}

// 测试测试用例删除成功
func TestTestcaseDelete_Success(t *testing.T) {
	tc := testcase.NewTestcase(
		testcase.WithProblemId(1),
		testcase.WithSerial(uint16(rand.Intn(100))),
		testcase.WithTestInput(randomTestInput()),
		testcase.WithTestOutput(randomTestOutput()),
	)
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
	tc := testcase.NewTestcase(
		testcase.WithId(99999999),
		testcase.WithProblemId(1),
		testcase.WithSerial(1),
		testcase.WithTestInput(randomTestInput()),
		testcase.WithTestOutput(randomTestOutput()),
	)
	err := tc.Delete()
	if err == nil {
		t.Fatalf("删除不存在的测试用例应失败")
	}
}
