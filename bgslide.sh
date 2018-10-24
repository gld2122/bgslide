# !/bin/bash

if [[ $# -eq 0 ]] ; then
  killall bgslide
  nohup go run bgslide &
else
  dir=$(greadlink -f "$1")
  killall bgslide
  nohup go run bgslide -dir $dir &
fi
