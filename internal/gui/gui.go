package gui

import (
	"fmt"
	"os"
	"os/exec"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"github.com/lucasepe/runpad/internal/config"
)

func Show(all []config.Item, title string) {
	a := app.New()
	w := a.NewWindow(title)

	list := widget.NewList(
		func() int {
			return len(all)
		},
		func() fyne.CanvasObject {
			return widget.NewButton("xxx", nil)
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			el := all[i]
			bt := o.(*widget.Button)
			bt.SetText(el.Label())
			bt.OnTapped = onTappedFunc(bt, el)
		},
	)

	w.SetContent(list)
	w.Resize(fyne.NewSize(180, 240))
	w.ShowAndRun()
}

func onTappedFunc(bt *widget.Button, el config.Item) func() {
	return func() {
		bt.Disable()
		defer bt.Enable()

		cmd := exec.Command(el.Args()[0], el.Args()[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run() // cmd.Start()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
		}
	}
}
