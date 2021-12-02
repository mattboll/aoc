#!/bin/sh 

j=0
k=0

while read LINE; do
  n=`echo $LINE | cut -d' ' -f2`
  echo $LINE | grep "down" > /dev/null && j=$(($j + $n))
  echo $LINE | grep "up" > /dev/null && j=$(( $j - $n )) 
  echo $LINE | grep "forward" > /dev/null && k=$(( $k + $n )) 
done < input.txt
echo j : $j - k : $k - j*k $(($j * $k))
