package button

import (
	"bufio"
	"fmt"
	"io"
	"os"
	filecontrollers "texteditor/controllers/fileControllers"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)


	func Button(width float32,height float32) *widget.Button{
        icon,err := os.Open("/Users/asker/Desktop/Text-Editor/assets/Plus.png")
		if err != nil{
			fmt.Println(err)
		}

		r := bufio.NewReader(icon)

		b,err := io.ReadAll(r)

		if err != nil{
			fmt.Println(err)
		}

		btn := widget.NewButtonWithIcon("", fyne.NewStaticResource("icon", b), func() {
			filecontrollers.CreateFile()

		})
		btn.Resize(fyne.NewSize(width,height))
		return btn


	}