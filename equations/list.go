package equations

import (
	"log/slog"

	"github.com/stsh89/math-go/equation"
)

type EquationsLister struct{}

func (l EquationsLister) List() []string {
	return []string{
		"y=sinx",
		"y=ab+x",
	}
}

func ListEquations(logger *slog.Logger) []string {
	return equation.List(EquationsLister{}, logger)
}
