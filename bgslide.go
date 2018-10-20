package main

import (
	"io/ioutil"
	"math/rand"
	"os/exec"
	"time"
	"flag"
  "bgslide/src"
)

var (
  photoList []string
  dir string
  home string
)

const FLAG_HELP string = "Your Photo Directory"

func main() {

  flag.StringVar(&dir, "dir", src.GetDefaultDir() , FLAG_HELP)
  flag.Parse()

  // Change photo every 60 seconds
	for range time.NewTicker(time.Second * 60).C {

    // Refreshes background directory on every iteration
	  photoList = src.GetNames(dir)
		x := rand.Intn(len(photoList))
	  ioutil.WriteFile("./script", []byte(src.Build(dir,photoList[x])), 0755)

	  exec.Command("osascript", "./script").Run()
	  exec.Command("rm", "./script").Run()

	}

}
