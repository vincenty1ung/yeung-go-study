package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

// build : fyne package -os ios  -profile C3CBHTX5N7
func main() {

	/*off := make(chan struct{})
	defer tidyUp(off)
	myApp := app.New()
	myWindow := myApp.NewWindow("定时器")

	myWindow.SetContent(widget.NewLabel("Hello"))
	// 设置窗口分辨率
	myWindow.Resize(fyne.NewSize(435, 256))

	// 这设置为主窗口
	myWindow.SetMaster()

	myWindow.SetContent(
		widget.NewButton(
			"Open new", func() {
				w3 := myApp.NewWindow("Third")
				// w3.SetContent(widget.NewLabel("Third"))
				w3.Resize(fyne.NewSize(435, 256))

				clock := widget.NewLabel("")
				w3off := make(chan struct{})
				w3.SetOnClosed(
					func() {
						close(w3off)
					},
				)
				w3.SetContent(clock)
				formatted := time.Now().Format("Time: 03:04:05")
				clock.SetText(formatted)

				go func() {
					defer fmt.Println("g1 end")
					ticker := time.NewTicker(time.Second)
					for {
						select {
						case <-ticker.C:
							updateTime(clock)
						case <-w3off:
							return
						}
					}

				}()
				w3.Show()

			},
		),
	)

	// myWindow.SetContent(widget.NewLabel("Third"))
	myWindow.Show()
	myApp.Run()*/

	off := make(chan struct{})
	defer tidyUp(off)
	myApp := app.New()
	𓂺 := myApp.NewWindow("𓂺")
	newLabel := widget.NewLabel("Hello")
	𓂺.SetContent(newLabel)
	// 设置窗口分辨率
	// myWindow.Resize(fyne.NewSize(435, 256))

	// 这设置为主窗口

	go func() {
		for _𓂺 := range time.NewTicker(1 * time.Second).C {
			_ = _𓂺
			newLabel.SetText(fmt.Sprintf("𓂺==>%v", time.Now()))
		}
	}()
	// myWindow.SetContent(widget.NewLabel("Third"))
	𓂺.Show()
	myApp.Run()
}

func updateTime(clock *widget.Label) {
	formatted := time.Now().Format("Time: 03:04:05")
	clock.SetText(formatted)
	fmt.Println(11)
}

func tidyUp(off chan struct{}) {
	close(off)
	fmt.Println("Exited")
}
