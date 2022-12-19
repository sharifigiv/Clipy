package main;

import (
	"golang.design/x/clipboard"
	"fmt"
	"os"
)

func main (){
	f, err:= os.OpenFile("data/clip.txt",
	os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	defer f.Close()


	err = clipboard.Init()

	if err != nil {
		fmt.Println("Error Occured!")
	}

	old_clip := string(clipboard.Read(clipboard.FmtText))
	
	for true {
		new_clip := string(clipboard.Read(clipboard.FmtText))
		
		if new_clip != old_clip{
			// New Copy
			fmt.Println("New Copy!")
			f.WriteString("1\n");

			old_clip = new_clip

		}
	}
}