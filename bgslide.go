package main

import (
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"time"
	"flag"
  "bgslide/src"
)

var (
  photoList []string
  dir string
)

const FLAG_HELP string = "Your Photo Directory"

func main() {

  flag.StringVar(&dir, "dir", "/Users/defreitas/Pictures/bgslide", FLAG_HELP)
  flag.Parse()

  // Change every 60 seconds
	for range time.NewTicker(time.Second * 60).C {

    // Refresh photo dir every iteration
	  photoList = getNames(dir)
		x := rand.Intn(len(photoList))
	  ioutil.WriteFile("./script", []byte(src.Build(dir,photoList[x])), 0755)
	  exec.Command("osascript", "./script").Run()
	  exec.Command("rm", "./script").Run()

	}

}

// Make a slice containing the names of each file in dir
func getNames(dir string) []string {

	f, err := os.Open(dir)
	src.Check(err)

	list, err := f.Readdirnames(0)
	src.Check(err)

	return list

}
