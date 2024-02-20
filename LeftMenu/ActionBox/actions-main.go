package ActionBox

import "github.com/rivo/tview"

func GetActionsBox() tview.Primitive {
	return tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(tview.NewTextView().SetText("Actions"), 0, 0, false).
		SetBorder(true)
}
