package main

import "os/user"

// Will look for pictures by default in ~/Pictures/bgslide
func getDefaultDir() string {

	usr, err := user.Current()
	Check(err)

	return string(usr.HomeDir) + "/Pictures/bgslide"

}
