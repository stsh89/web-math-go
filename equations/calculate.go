package equations

import (
	"fmt"
	"log/slog"

	"github.com/stsh89/math-go/equation"
	"github.com/stsh89/web-math-go/app"
)

func Calculate(term string, args map[string]float64, logger *slog.Logger, config *app.Config) []map[string]string {
	CreateEquation(term, logger, config)

	points := equation.CalculateRange(term, args, logger)

	var mapPoints []map[string]string

	for _, point := range points {
		mapPoints = append(mapPoints, map[string]string{"x": floatToString(point.X), "y": floatToString(point.Y)})
	}

	return mapPoints
}

func floatToString(value float64) string {
	return fmt.Sprintf("%0.2f", value)
}
