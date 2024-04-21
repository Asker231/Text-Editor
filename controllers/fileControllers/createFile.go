package filecontrollers

import (
	"fmt"
	"os"
	"strconv"
)


    func Counter()func()int{
     	var num int  = 0	
	    return func()int{
	     	 num+=1
		     return num
     	}
   }
	func CreateFile(){
        count := Counter() 
		num := strconv.Itoa(count())
		_ , err := os.Create("/Users/asker/Desktop/Text-Editor/db/" + num +".txt")
		if err != nil{
			fmt.Println(err)
		}
	}