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
  picList []string
  dirFlag *string
)

const FLAG_HELP string = "Your Photo Directory"

func main() {

  dirFlag = flag.String("dir", "/Users/defreitas/Pictures/bgslide", FLAG_HELP)
  flag.Parse()

	picList = getNames(*dirFlag)

	for range time.NewTicker(time.Second * 60).C {

		changePhoto(rand.Intn(len(picList)))

	}

}

func changePhoto(c int) {

	src.Check(ioutil.WriteFile("./script", []byte(src.Build(*dirFlag,picList[c])), 0755))
	src.Check(exec.Command("osascript", "./script").Run())
	src.Check(exec.Command("rm", "./script").Run())

}

func getNames(dirName string) []string {

	f, err := os.Open(dirName)
	src.Check(err)

	list, err := f.Readdirnames(0)
	src.Check(err)

	return list

}
