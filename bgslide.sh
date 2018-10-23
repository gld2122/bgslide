# !/bin/bash

if [ $# -eq 0 ] ; then
  killall bgslide
  nohup go run bgslide &
else
  killall bgslide
  nohup go run bgslide -dir "$1" &
fi
