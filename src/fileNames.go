package src

import "os"

// Make a slice containing the names of each file in dir
func GetNames(dir string) []string {

	f, err := os.Open(dir)
	Check(err)

	list, err := f.Readdirnames(0)
	Check(err)

	return list

}
