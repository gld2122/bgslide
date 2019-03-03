package main

import (
	"flag"
	"math/rand"
	"os/exec"
	"time"
)

var (
	photoList []string
	dir       string
)

const FLAG_HELP string = "Your Photo Directory"

func main() {

	rand.Seed(time.Now().UnixNano())

	flag.StringVar(&dir, "dir", GetDefaultDir(), FLAG_HELP)
	flag.Parse()

	if CheckIsDir(dir) {

		// Change wallpaper when program starts (helps debug)
		oneCycle()

		// Changes wallpaper every 60 seconds
		for range time.NewTicker(time.Second * 60).C {

			oneCycle()

		}

	} else {

		changeToPhoto()

	}

}

func oneCycle() {

	photoList = GetNames(dir)
	x := rand.Intn(len(photoList))

	// AppleScript to change wallpaper
	exec.Command("osascript", "-e", BuildDir(dir, photoList[x])).Run()

}

func changeToPhoto() {

	exec.Command("osascript", "-e", BuildFile(dir)).Run()

}
