package equations

import (
	"log/slog"

	"github.com/stsh89/web-math-go/app"
	"github.com/stsh89/web-math-go/providers/notion"

	"github.com/stsh89/math-go/equation"
)

type EquationsCreator struct {
	Logger *slog.Logger
	Config *app.Config
}

func (l *EquationsCreator) Save(term string) string {
	if l.Config == nil {
		l.Logger.Error("Missing equation creator configuration")
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

	return client.CreatePage(term)
}

func CreateEquation(term string, logger *slog.Logger, config *app.Config) string {
	creator := EquationsCreator{
		Logger: logger,
		Config: config,
	}

	return equation.Save(term, &creator, logger)
}
