package LeftMenu

import (
	"ActionBox"
	"ConfigBox"
	"StatusBox"
	"github.com/rivo/tview"
)

func GetLeftBox() tview.Primitive {
	LeftFlex := tview.NewFlex().SetDirection(tview.FlexColumnCSS).
		AddItem(StatusBox.GetStatusBox(), 0, 1, false).
		AddItem(ActionBox.GetActionsBox(), 0, 3, false).
		AddItem(ConfigBox.GetConfigBox(), 0, 2, false)

	return LeftFlex

}
