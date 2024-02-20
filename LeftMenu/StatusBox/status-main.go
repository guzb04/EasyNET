package StatusBox

import "github.com/rivo/tview"

func GetStatusBox() tview.Primitive {
	statusBox := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(tview.NewTextView().SetText("Status"), 0, 0, false).
		SetBorder(true)
	return statusBox
}
