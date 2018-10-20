package src

import (
	"os/user"
)

func GetDefaultDir() string {

return addDir( getHome() )

}

func getHome() string {

	usr, err := user.Current()
  Check( err )

	return string(usr.HomeDir)

}

func addDir( s string ) string {

  return s + "/Pictures/bgslide"

}
