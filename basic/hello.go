package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func setCursor(on bool) {
	// use escape code to turn cursor on or off
	if on {
		fmt.Print("\x1B[?25h")
	} else {
		fmt.Print("\x1B[?25l")
	}
}

func setFont(name string) {
	// use setfont command to change the font
	// run `ls /usr/share/consolefonts` in a terminal to get a list of available fonts
	exec.Command("setfont", name).Run()
}

func main() {
	// redirect log to stderr so that it prints messages in VS Code
	// 由此可见标准错误输出会从vsc的output中输出，这是怎么做到的呢
	log.SetOutput(os.Stderr)

	// configure the console
	setFont("Lat15-Terminus32x16")
	setCursor(false)

	// print some messages
	fmt.Println("Hello EV3!")
	time.Sleep(time.Second * 10)
	log.Println("Hello VS Code!")

	// give some time to look at the screen before the program exits
	time.Sleep(time.Second * 5)
}
