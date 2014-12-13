#!/bin/sh
SERVICE='main.go'

while sleep 1;
do
	if ps ax | grep -v grep | grep $SERVICE > /dev/null
	then
		echo "$SERVICE service running, everything is fine"
	else
		go run main.go
		echo "$SERVICE is not running"
		echo "$SERVICE is not running!" | mail -s "$SERVICE down" root
	fi
done
