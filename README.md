# bgslide
Go script to customize your desktop photo slideshow on macOS.

## To Use

1. While inside repository, run `go build; go install`.
2. Place all background photos into one folder. Default if no flag is
   specified is `~/Pictures/bgslide`.
3. Run the following command `nohup go run bgslide -dir <folder-path> &`.

## Dependencies

* Up-to-date Go installation (available at https://golang.org/)
