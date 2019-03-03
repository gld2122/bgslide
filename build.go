package main

// Builds an AppleScript to call a given photo from the background directory.
// Should also work across multiple desktops.
func BuildDir(d string, f string) string {

	str := "tell application \"System Events\"\n"
	str += "set theDesktops to a reference to every desktop\n"
	str += "repeat with x from 1 to (count theDesktops)\n"
	str += "set picture of item x of the theDesktops to " + dfCombine(d, f) + "\n"
	str += "end repeat\n"
	str += "end tell\n"

	return str

}

func BuildFile(f string) string {

	str := "tell application \"System Events\"\n"
	str += "set theDesktops to a reference to every desktop\n"
	str += "repeat with x from 1 to (count theDesktops)\n"
	str += "set picture of item x of the theDesktops to " + fCombine(f) + "\n"
	str += "end repeat\n"
	str += "end tell\n"

	return str

}

func fCombine(f string) string {

	return "\"" + f + "\""

}

func dfCombine(d string, f string) string {

	return "\"" + d + "/" + f + "\""

}
