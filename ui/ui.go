package ui

import (
	"github.com/rivo/tview"
)

type UI struct {
	app            *tview.Application
	graphics       *tview.TextView
	characterStats *tview.TextView
	gameText       *tview.TextView
	actions        *tview.TextView
}

func New() (*UI, error) {
	ui := &UI{
		app:            tview.NewApplication(),
		graphics:       tview.NewTextView(),
		characterStats: tview.NewTextView(),
		gameText:       tview.NewTextView(),
		actions:        tview.NewTextView(),
	}
	// Set up the layout
	grid := tview.NewGrid().
		// set 3 rows and 2 columns
		SetRows(0, 0, 3).
		SetColumns(0, 0).
		SetBorders(true).
		AddItem(ui.graphics, 0, 0, 1, 1, 0, 0, false).
		AddItem(ui.characterStats, 0, 1, 1, 1, 0, 0, false).
		AddItem(ui.gameText, 1, 0, 1, 2, 0, 0, false).
		AddItem(ui.actions, 2, 0, 1, 2, 0, 0, true)

	ui.app.SetRoot(grid, true)

	return ui, nil
}

// Add methods to update each section of the UI
func (ui *UI) UpdateGraphics(content string) {
	ui.graphics.SetText(content)
}
func (ui *UI) UpdateCharacterStats(stats string) {
	ui.characterStats.SetText(stats)
}

func (ui *UI) AppendGameText(text string) {
	ui.gameText.SetText(ui.gameText.GetText(true) + "\n" + text)
}

func (ui *UI) UpdateActions(actions string) {
	ui.actions.SetText(actions)
}

func (ui *UI) Run() error {
	return ui.app.Run()
}
