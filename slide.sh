# !/bin/bash

if [[ $# -eq 0 ]] ; then
  killall bgslide
  bgslide &
elif [[ "$1" -eq "stop" ]] ; then
	killall bgslide 
	if [[ $? -eq 1 ]] ; then
		printf "no running instances\n"
	else
		printf "bgslide instance stopped\n"
	fi
else
  file=$(greadlink -f "$1")
  killall bgslide
  bgslide -dir $file &
fi
