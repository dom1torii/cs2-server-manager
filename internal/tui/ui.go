package tui

import (
	"github.com/rivo/tview"
)

type UI struct {
  App   *tview.Application
  Pages *tview.Pages

  RefreshStartList func()
}

func New() *UI {
	app := tview.NewApplication()
	pages := tview.NewPages()

	return &UI{
		App: app,
		Pages: pages,
	}
}

func (ui *UI) Init() error {
  return ui.App.SetRoot(ui.Pages, true).Run()
}
