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

)

const FLAG_HELP string = "Your Photo Directory"

func main() {

  rand.Seed(time.Now().UnixNano())

  flag.StringVar(&dir, "dir", src.GetDefaultDir() , FLAG_HELP)
  flag.Parse()

  // Change wallpaper when program starts (helps debug)
  oneCycle()

  // Changes wallpaper every 60 seconds
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
