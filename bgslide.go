package main

import (
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"time"
	"my"
)

var picList []string

func main() {

	picList = getNames("/Users/defreitas/Pictures/bgslide")

	for range time.NewTicker(time.Second * 60).C {

		changePhoto(rand.Intn(len(picList)))

	}

}

func changePhoto(c int) {

	slc := []byte(build(picList[c]))
	my.Check(ioutil.WriteFile("./script", slc, 0755))

	my.Check(exec.Command("osascript", "./script").Run())
	my.Check(exec.Command("rm", "./script").Run())

}

func build(s string) string {

	str := "tell application \"System Events\"\n"
	str += "set theDesktops to a reference to every desktop\n"
	str += "repeat with x from 1 to (count theDesktops)\n"
	str += "set picture of item x of the theDesktops to \"~/Pictures/bgslide/" + s + "\"\n"
	str += "end repeat\n"
	str += "end tell\n"

	return str

}

func getNames(dirName string) []string {

	f, err := os.Open(dirName)
	my.Check(err)

	list, err := f.Readdirnames(0)
	my.Check(err)

	return list

}
