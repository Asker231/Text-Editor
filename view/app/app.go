package app

import (

	"texteditor/view/components/button"


	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

	func App(width float32 ,height float32){
		// isTaped := false
		 application := app.New()
		 view := application.NewWindow("TextEditor")
		 view.Resize(fyne.NewSize(width,height))
		 ct := container.NewWithoutLayout(button.Button(100,100))
		 ct.Resize(fyne.NewSize(1200,1200))
		 hstack := container.NewHBox(layout.NewSpacer(),ct,layout.NewSpacer())
		 vstack := container.NewVBox(layout.NewSpacer(),hstack,layout.NewSpacer())
		 view.SetContent(vstack)
		 view.ShowAndRun()
	}