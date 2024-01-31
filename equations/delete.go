package equations

import (
	"log/slog"

	"github.com/stsh89/web-math-go/app"
	"github.com/stsh89/web-math-go/providers/notion"

	"github.com/stsh89/math-go/equation"
)

type EquationsDeletor struct {
	Logger *slog.Logger
	Config *app.Config
}

func (l *EquationsDeletor) Delete(id string) string {
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

	return client.ArchivePage(id)
}

func DeleteEquation(term string, logger *slog.Logger, config *app.Config) string {
	deletor := EquationsDeletor{
		Logger: logger,
		Config: config,
	}

	id := SearchEquation(term, logger, config)

	return equation.Delete(id, &deletor, logger)
}
