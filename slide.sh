# !/bin/bash

if [[ $# -eq 0 ]] ; then
  killall bgslide 2>/dev/null
  bgslide &
elif [[ "$1" -eq "stop" ]] ; then
	killall bgslide 2>/dev/null
	if [[ $? -eq 1 ]] ; then
		printf "no running instances\n"
	else
		printf "bgslide instance stopped\n"
	fi
else
  file=$(greadlink -f "$1")
  killall bgslide 2>/dev/null
  bgslide -dir $file &
fi
