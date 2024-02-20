package ConfigBox

import "github.com/rivo/tview"

func GetConfigBox() tview.Primitive {
	return tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(tview.NewTextView().SetText("Config"), 0, 0, false).
		SetBorder(true)
}
