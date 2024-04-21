package button

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

	func Counter()func()int{
		var num int  = 0
		return func()int{
			 num+=1
			return num
		}
	}


	func Button(width float32,height float32) *widget.Button{
        icon,err := os.Open("/Users/asker/Desktop/Text-Editor/assets/Plus.png")
        count := Counter() 
		if err != nil{
			fmt.Println(err)
		}

		r := bufio.NewReader(icon)

		b,err := io.ReadAll(r)

		if err != nil{
			fmt.Println(err)
		}

		btn := widget.NewButtonWithIcon("", fyne.NewStaticResource("icon", b), func() {
			fmt.Println(count())
		})
		btn.Resize(fyne.NewSize(width,height))
		return btn


	}