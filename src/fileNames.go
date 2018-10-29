package src

import "os"

// Make a slice containing the names of each file in dir
func GetNames(dir string) []string {

  if ( checkIsDir( dir ) ) {

	  f, err := os.Open(dir)
	  Check(err)

	  list, err := f.Readdirnames(0)
	  Check(err)

	  return list

  } else {

    return []string{dir}

  }

}

func checkIsDir( dir string ) bool {

  f , err := os.Stat( dir )
  Check(err)

  return f.IsDir()

}
