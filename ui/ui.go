package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// create UI struct
type UI struct {
	app            *tview.Application
	graphics       *tview.TextView
	characterStats *tview.TextView
	gameText       *tview.TextView
	actions        *tview.Flex
	moveButton     *tview.Button
	attackButton   *tview.Button
	inputChan      chan string
	done           chan struct{}
}

func New() (*UI, error) {
	ui := &UI{
		app:            tview.NewApplication(),
		graphics:       tview.NewTextView(),
		characterStats: tview.NewTextView(),
		gameText:       tview.NewTextView(),
		actions:        tview.NewFlex(),
		inputChan:      make(chan string),
		done:           make(chan struct{}),
	}
	// Create Flex containers for each section
	graphicsFlex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(ui.graphics, 0, 1, false)
	statsFlex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(ui.characterStats, 0, 1, false)
	textFlex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(ui.gameText, 0, 1, false)
	actionsFlex := tview.NewFlex().SetDirection(tview.FlexRow)

	// Set up the layout
	grid := tview.NewGrid().
		// set 3 rows and 2 columns
		SetRows(0, 0, 3).
		SetColumns(0, 0).
		SetBorders(false).
		AddItem(graphicsFlex, 0, 0, 1, 1, 0, 0, false).
		AddItem(statsFlex, 0, 1, 1, 1, 0, 0, false).
		AddItem(textFlex, 1, 0, 1, 2, 0, 0, false).
		AddItem(actionsFlex, 2, 0, 1, 2, 0, 0, true)

	actionsGrid := tview.NewGrid().
		SetRows(0).
		SetColumns(0).
		SetBorders(false).
		SetGap(2, 2)

	actionsFlex.AddItem(actionsGrid, 0, 1, false)
	actionsFlex.SetTitle("Actions").SetBorder(true)

	// Add titles to the sections
	ui.graphics.SetTitle("Graphics").SetBorder(true)
	ui.characterStats.SetTitle("Character Stats").SetBorder(true)
	ui.gameText.SetTitle("Game Text").SetBorder(true)
	ui.actions.SetTitle("Actions").SetBorder(true)

	ui.app.SetRoot(grid, true)

	// Create the move button
	ui.moveButton = tview.NewButton("Move (M)").SetSelectedFunc(func() {
		ui.inputChan <- "move" //send move to the input channel when selected
	})

	// Create the attack button
	ui.attackButton = tview.NewButton("Attack (A)").SetSelectedFunc(func() {
		ui.inputChan <- "attack"
	})
	actionsGrid.AddItem(ui.moveButton, 0, 0, 1, 1, 0, 0, false)
	actionsGrid.AddItem(ui.attackButton, 0, 1, 1, 1, 0, 0, false)

	//ui.actions.AddItem(ui.moveButton, 0, 1, false)
	//ui.actions.AddItem(ui.attackButton, 0, 1, false)

	// Set up input handling
	ui.app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyRune:
			switch event.Rune() {
			case 'm', 'M':
				ui.inputChan <- "move"
				ui.app.SetFocus(ui.moveButton)
			case 'a', 'A':
				ui.inputChan <- "attack"
				ui.app.SetFocus(ui.attackButton)
			case 'q', 'Q':
				ui.inputChan <- "quit"
				close(ui.done)
				ui.app.Stop()
			}
		case tcell.KeyEsc:
			ui.inputChan <- "quit"
			close(ui.done)
			ui.app.Stop()
		}
		return event
	})

	return ui, nil
}

func (ui *UI) Run() error {
	go func() {
		if err := ui.app.Run(); err != nil {
			panic(err)
		}
	}()

	<-ui.done
	return nil
}

func (ui *UI) Stop() {
	close(ui.done)
	ui.app.Stop()
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

//func (ui *UI) Run() error {
//	return ui.app.Run()
//}

func (ui *UI) GetInputChannel() <-chan string {
	return ui.inputChan
}
