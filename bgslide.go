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

  // Change background photo when program starts (helps me debug)
  oneCycle()

  // Changes background photo every 60 seconds
	for range time.NewTicker( time.Second * 60 ).C {

    oneCycle()

	}

}

func oneCycle() {

  photoList = src.GetNames( dir )
  x := rand.Intn( len( photoList ) )
  slc := []byte( src.Build( dir , photoList[x] ) )

  ioutil.WriteFile( "./script" , slc  , 0755 )

  exec.Command("osascript", "./script").Run()
  exec.Command("rm", "./script").Run()

}
