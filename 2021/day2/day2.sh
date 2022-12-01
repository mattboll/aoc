#!/bin/sh 

aim=0
horizontal=0
depth=0

while read LINE; do
  n=`echo $LINE | cut -d' ' -f2`
  echo $LINE | grep "down" > /dev/null && aim=$(($aim + $n))
  echo $LINE | grep "up" > /dev/null && aim=$(( $aim - $n ))
  echo $LINE | grep "forward" > /dev/null && horizontal=$(( $horizontal + $n )) && depth=$(( $depth + $aim * $n ))
done < input.txt
echo depth : $depth - horizontal : $horizontal - depth*horizontal $(($depth * $horizontal))
