package problem

import (
	"STUOJ/external/neko"
	"STUOJ/internal/model"
)

func Generate(pi model.NekoProblemInstruction) (model.NekoProblem, error) {
	return neko.GenerateProblem(pi)
}

func Translate(p model.NekoTranslateInstruction) (model.NekoProblem, error) {
	return neko.TranslateProblem(p)
}
