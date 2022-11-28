package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

const (
	appID    = "dev-common-tool.lynwee.hou"
	HomePage = "https://devcommontool.vercel.app"
)

func IsJsonString(str string) error {
	var i interface{}
	return json.Unmarshal([]byte(str), &i)
}

func makeMenu(app fyne.App, window fyne.Window) {
	help := fyne.NewMenu("Help",
		fyne.NewMenuItem("Documentation", func() {
			u, _ := url.Parse(fmt.Sprintf("%s/doc", HomePage))
			app.OpenURL(u)
		}),
		fyne.NewMenuItem("Support", func() {
			u, _ := url.Parse(fmt.Sprintf("%s/suppport", HomePage))
			app.OpenURL(u)
		}),
		fyne.NewMenuItem("About", func() {
			u, _ := url.Parse(fmt.Sprintf("%s/about", HomePage))
			app.OpenURL(u)
		}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("Sponsor", func() {
			u, _ := url.Parse(fmt.Sprintf("%s/sponsor", HomePage))
			app.OpenURL(u)
		}),
	)
	file := fyne.NewMenu("File",
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("Exit", func() {
			log.Printf("bye!")
			app.Quit()
		}),
	)
	main := fyne.NewMainMenu(file, help)
	window.SetMainMenu(main)
	window.SetMaster()
}

func SetSystemTray(app fyne.App, mainWindow fyne.Window, inputText, formatted *widget.Entry) {
	if desk, ok := app.(desktop.App); ok {
		unixTimestamp := fyne.NewMenuItem("Unix Timestamp", func() {})
		unixTimestamp.Action = func() {
			mainWindow.Show()
		}
		jsonOperation := fyne.NewMenuItem("Json", func() {})
		jsonOperation.Action = func() {
			NewPrettyJsonWindow(app, inputText, formatted).Show()
		}
		menu := fyne.NewMenu("Open", unixTimestamp, jsonOperation)
		desk.SetSystemTrayMenu(menu)
	}
}

func main() {
	app := app.NewWithID(appID)
	setTheme(app)
	// hiddenWindow := app.NewWindow("dct")
	// hiddenWindow.Hide()
	// hiddenWindow.SetMaster()

	window := app.NewWindow("Dev Common Tool")
	window.Resize(fyne.Size{
		Width:  400,
		Height: 200,
	})
	makeMenu(app, window)

	formatted := widget.NewMultiLineEntry()
	formatted.TextStyle.Bold = true

	inputText := widget.NewEntry()
	inputText.MultiLine = true
	inputText.OnChanged = func(s string) {
		inputText.SetText(s)
	}
	inputText.Size()
	inputText.Wrapping = fyne.TextTruncate
	inputText.Validator = IsJsonString
	inputText.Show()
	inputText.OnChanged = func(s string) {
		var i interface{}
		err := json.Unmarshal([]byte(s), &i)
		if err != nil {
			fmt.Println("not a json string")
			// 修改样式
		} else {
			data, err := json.MarshalIndent(i, "", "\t")
			if err != nil {
				fmt.Println("not a json string")
				// 异常情况
			} else {
				var formattedStr string = string(data)
				formatted.SetText(formattedStr)
				formatted.Refresh()
			}
		}
	}

	window.SetContent(container.New(
		layout.NewVBoxLayout(),
		widget.NewButtonWithIcon("Unix Timestamp", theme.SettingsIcon(), func() {
			clipboard := fyne.CurrentApp().Driver().AllWindows()[0].Clipboard()
			clipboard.SetContent(fmt.Sprintf("%d", time.Now().Unix()))
		}),
		container.New(layout.NewGridLayoutWithColumns(2),
			widget.NewButton("-1h", func() {
				clipboard := fyne.CurrentApp().Driver().AllWindows()[0].Clipboard()
				clipboard.SetContent(fmt.Sprintf("%d", time.Now().Unix()-86400/24))
			}),
			widget.NewButton("+1h", func() {
				clipboard := fyne.CurrentApp().Driver().AllWindows()[0].Clipboard()
				clipboard.SetContent(fmt.Sprintf("%d", time.Now().Unix()+86400/24))
			}),
			widget.NewButton("-2h", func() {
				clipboard := fyne.CurrentApp().Driver().AllWindows()[0].Clipboard()
				clipboard.SetContent(fmt.Sprintf("%d", time.Now().Unix()-3600*2))
			}),
			widget.NewButton("+2h", func() {
				clipboard := fyne.CurrentApp().Driver().AllWindows()[0].Clipboard()
				clipboard.SetContent(fmt.Sprintf("%d", time.Now().Unix()+3600*2))
			}),
			widget.NewButton("-3h", func() {
				clipboard := fyne.CurrentApp().Driver().AllWindows()[0].Clipboard()
				clipboard.SetContent(fmt.Sprintf("%d", time.Now().Unix()-3600*3))
			}),
			widget.NewButton("+3h", func() {
				clipboard := fyne.CurrentApp().Driver().AllWindows()[0].Clipboard()
				clipboard.SetContent(fmt.Sprintf("%d", time.Now().Unix()+3600*3))
			}),
			widget.NewButton("-6h", func() {
				clipboard := fyne.CurrentApp().Driver().AllWindows()[0].Clipboard()
				clipboard.SetContent(fmt.Sprintf("%d", time.Now().Unix()-86400/4))
			}),
			widget.NewButton("+6h", func() {
				clipboard := fyne.CurrentApp().Driver().AllWindows()[0].Clipboard()
				clipboard.SetContent(fmt.Sprintf("%d", time.Now().Unix()+86400/4))
			}),
			widget.NewButton("-12h", func() {
				clipboard := fyne.CurrentApp().Driver().AllWindows()[0].Clipboard()
				clipboard.SetContent(fmt.Sprintf("%d", time.Now().Unix()-86400/2))
			}),
			widget.NewButton("+12h", func() {
				clipboard := fyne.CurrentApp().Driver().AllWindows()[0].Clipboard()
				clipboard.SetContent(fmt.Sprintf("%d", time.Now().Unix()+86400/2))
			}),
			widget.NewButton("-1d", func() {
				clipboard := fyne.CurrentApp().Driver().AllWindows()[0].Clipboard()
				clipboard.SetContent(fmt.Sprintf("%d", time.Now().Unix()-86400))
			}),
			widget.NewButton("+1d", func() {
				clipboard := fyne.CurrentApp().Driver().AllWindows()[0].Clipboard()
				clipboard.SetContent(fmt.Sprintf("%d", time.Now().Unix()+86400))
			}),
			widget.NewButton("-7d", func() {
				clipboard := fyne.CurrentApp().Driver().AllWindows()[0].Clipboard()
				clipboard.SetContent(fmt.Sprintf("%d", time.Now().Unix()-86400*7))
			}),
			widget.NewButton("+7d", func() {
				clipboard := fyne.CurrentApp().Driver().AllWindows()[0].Clipboard()
				clipboard.SetContent(fmt.Sprintf("%d", time.Now().Unix()+86400*7))
			})),
		// widget.NewSeparator(),
		widget.NewButtonWithIcon("Pretty JSON", theme.DocumentIcon(), func() {
			NewPrettyJsonWindow(app, inputText, formatted).Show()
		}),
	))
	SetSystemTray(app, window, inputText, formatted)
	window.SetCloseIntercept(func() {
		window.Hide()
	})
	app.Run()
}

func NewPrettyJsonWindow(app fyne.App, inputText, formatted *widget.Entry) fyne.Window {
	prettyJsonWindow := app.NewWindow("Pretty JSON")
	prettyJsonWindow.Resize(fyne.Size{
		Width:  800,
		Height: 600,
	})
	prettyJsonWindow.SetContent(
		container.NewBorder(
			nil,
			container.NewGridWithColumns(2,
				widget.NewButtonWithIcon("Format", theme.ViewRefreshIcon(), func() {
					fmt.Println("click format button")
					str := inputText.Text
					var i interface{}
					err := json.Unmarshal([]byte(str), &i)
					if err != nil {
						fmt.Println("not a json string")
						// 修改样式
					} else {
						data, err := json.MarshalIndent(i, "", "\t")
						if err != nil {
							fmt.Println("not a json string")
							// 异常情况
						} else {
							var formattedStr string = string(data)
							inputText.SetText(formattedStr)
							inputText.Refresh()
						}
					}
				}),
				widget.NewButtonWithIcon("Copy", theme.ContentCopyIcon(), func() {
					clipboard := fyne.CurrentApp().Driver().AllWindows()[0].Clipboard()
					clipboard.SetContent(formatted.Text)
				}),
			),
			nil,
			nil,
			container.New(layout.NewGridLayoutWithColumns(2), inputText, formatted),
		),
	)
	return prettyJsonWindow
}
