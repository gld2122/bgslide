package main

import "os"

// Make a slice containing the names of each file in dir
func GetNames(dir string) []string {

	f, err := os.Open(dir)
	Check(err)

	list, err := f.Readdirnames(0)
	Check(err)

	return list

}

func CheckIsDir(dir string) bool {

	f, err := os.Stat(dir)
	Check(err)

	return f.IsDir()

}
