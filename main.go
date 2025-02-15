package main

import (
	"fmt"
	"time"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {
	//TIP Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined or highlighted text
	// to see how GoLand suggests fixing it.

	currentTime := time.Now()
	currentTimeFormat := currentTime.Format("2006-01-02 15:04:05")
	fmt.Printf("Service Started, %s\n", currentTimeFormat)

	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	for i := 1; i <= 5; i++ {
		//TIP You can try debugging your code. We have set one <icon src="AllIcons.Debugger.Db_set_breakpoint"/> breakpoint
		// for you, but you can always add more by pressing <shortcut actionId="ToggleLineBreakpoint"/>. To start your debugging session,
		// right-click your code in the editor and select the <b>Debug</b> option.
		fmt.Println("i =", 100/i)
	}
	fmt.Println("Well, I'm done.")

	//i := 0
	//for {
	//	fmt.Println("current =", i)
	//	i++
	//	time.Sleep(5 * time.Second)
	//}
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
