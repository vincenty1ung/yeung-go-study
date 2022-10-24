package main

import (
	"fmt"
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Canvas")
	myCanvas := myWindow.Canvas()

	blue := color.NRGBA{R: 0, G: 0, B: 180, A: 255}
	rect := canvas.NewRectangle(blue)
	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
	rect.FillColor = green
	myCanvas.SetContent(rect)

	go func() {
		time.Sleep(time.Second)
		setContentToCircle(myCanvas)
		// rect.Refresh()
	}()

	myWindow.Resize(fyne.NewSize(100, 100))
	myWindow.ShowAndRun()
	/*myApp := app.New()
	myWindow := myApp.NewWindow("Widget")

	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}

	text1 := canvas.NewText("Hello", green)
	text2 := canvas.NewText("There", green)
	// text2.Move(fyne.NewPos(20, 20))
	content := container.New(layout.NewCenterLayout(), text1, text2)

	myWindow.SetContent(content)

	myWindow.ShowAndRun()*/
}

func setContentToText(c fyne.Canvas) {
	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
	text := canvas.NewText("Text", green)
	text.TextStyle.Bold = true
	c.SetContent(text)
}

func setContentToCircle(c fyne.Canvas) {
	red := color.NRGBA{R: 0xff, G: 0x33, B: 0x33, A: 0xff}
	circle := canvas.NewCircle(color.White)
	circle.StrokeWidth = 4
	circle.StrokeColor = red
	ctrlTab := &desktop.CustomShortcut{KeyName: fyne.KeyF1, Modifier: fyne.KeyModifierShift}
	c.AddShortcut(
		ctrlTab, func(shortcut fyne.Shortcut) {
			fmt.Println("We tapped Ctrl+Tab")
		},
	)
	c.SetContent(circle)
}
