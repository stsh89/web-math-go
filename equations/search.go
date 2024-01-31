package equations

import (
	"log/slog"

	"github.com/stsh89/web-math-go/app"
	"github.com/stsh89/web-math-go/providers/notion"

	"github.com/stsh89/math-go/equation"
)

type EquationsFinder struct {
	Logger *slog.Logger
	Config *app.Config
}

func (l *EquationsFinder) Find(term string) string {
	if l.Config == nil {
		l.Logger.Error("Missing equation deletor configuration")
		return ""
	}

	if l.Config.NotionConfig == nil {
		l.Logger.Error("Missing Notion configuration")
		return ""
	}

	client := notion.Client{
		Configuration: *l.Config.NotionConfig,
		Logger:        l.Logger,
	}

	return client.SearchPage(term)
}

func SearchEquation(term string, logger *slog.Logger, config *app.Config) string {
	finder := EquationsFinder{
		Logger: logger,
		Config: config,
	}

	return equation.Find(term, &finder, logger)
}
