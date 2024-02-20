package main

import (
	"github.com/rivo/tview"
	//"github.com/gdamore/tcell/v2"
	"EasyNET/LeftMenu"
)

func main() {

	app := tview.NewApplication()

	mainFlex := tview.NewFlex().SetDirection(tview.FlexRowCSS)

	mainFlex.AddItem(LeftMenu.GetLeftBox(), 0, 1, false).
		AddItem(tview.NewBox().SetBorder(true), 0, 2, false)

	if err := app.SetRoot(mainFlex, true).Run(); err != nil {
		panic(err)
	}
}
