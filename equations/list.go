package equations

import (
	"log/slog"

	"github.com/stsh89/web-math-go/app"
	"github.com/stsh89/web-math-go/providers/notion"

	"github.com/stsh89/math-go/equation"
)

type EquationsLister struct {
	Logger *slog.Logger
	Config *app.Config
}

func (l *EquationsLister) List() []string {
	if l.Config == nil {
		l.Logger.Error("Missing equation lister configuration")
		return []string{}
	}

	if l.Config.NotionConfig == nil {
		l.Logger.Error("Missing Notion configuration")
		return []string{}
	}

	client := notion.Client{
		Configuration: *l.Config.NotionConfig,
		Logger:        l.Logger,
	}

	subItems := client.ListSubitems()

	var listing []string

	for _, item := range subItems {
		listing = append(listing, item.Properties.Name)
	}

	return listing
}

func ListEquations(logger *slog.Logger, config *app.Config) []string {
	lister := EquationsLister{
		Logger: logger,
		Config: config,
	}

	return equation.List(&lister, logger)
}
