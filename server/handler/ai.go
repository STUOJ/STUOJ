package handler

import (
	"STUOJ/external/neko"
	"github.com/gin-gonic/gin"
)

func ChatAssistant(c *gin.Context) {
	neko.Forward(c, "/chat/assistant")
}

func ParseProblem(c *gin.Context) {
	neko.Forward(c, "/problem/parse")
}

func TranslateProblem(c *gin.Context) {
	neko.Forward(c, "/problem/translate")
}

func GenerateProblem(c *gin.Context) {
	neko.Forward(c, "/problem/generate")
}

func GenerateTestcase(c *gin.Context) {
	neko.Forward(c, "/testcase/generate")
}

func GenerateSolution(c *gin.Context) {
	neko.Forward(c, "/solution/generate")
}

func SubmitVirtualJudge(c *gin.Context) {
	neko.Forward(c, "/judge/submit")
}

func TellJoke(c *gin.Context) {
	neko.Forward(c, "/misc/joke")
}
