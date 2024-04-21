package app

import (
	"texteditor/view/components/button"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

	func App(width float32 ,height float32){
		 application := app.New()
		 view := application.NewWindow("TextEditor")
		 view.Resize(fyne.NewSize(width,height))
         view.SetContent(button.Button(50,50))
		 view.ShowAndRun()
	}