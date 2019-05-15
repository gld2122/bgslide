package main

import (
	"flag"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

func main() {

	var dir string
	const FLAG_HELP string = "Your Photo Directory"

	flag.StringVar(&dir, "dir", getDefaultDir(), FLAG_HELP)
	flag.Parse()

	rand.Seed(time.Now().UnixNano())

	if checkDir(dir) {

		// Changes wallpaper every 60 seconds
		for range time.NewTicker(time.Second * 60).C {

			fp, err := os.Open(dir)
			Check(err)

			dirList, err := fp.Readdirnames(0)
			Check(err)

			x := rand.Intn(len(dirList))

			// AppleScript to change wallpaper
			exec.Command("osascript", "-e", BuildDir(dir, dirList[x])).Run()

		}

	} else {

		exec.Command("osascript", "-e", BuildFile(dir)).Run()

	}

}

func checkDir(dir string) bool {

	fp, err := os.Open(dir)
	Check(err)

	st, err := fp.Stat()
	Check(err)

	result := st.IsDir()

	Check(fp.Close())

	return result

}
