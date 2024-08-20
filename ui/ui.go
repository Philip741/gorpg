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

func New() *UI {
	ui := &UI{
		app:            tview.NewApplication(),
		graphics:       tview.NewTextView(),
		characterStats: tview.NewTextView(),
		gameText:       tview.NewTextView(),
		actions:        tview.NewTextView(),
	}
}
