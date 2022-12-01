#!/bin/sh

rm input2.txt
j=0;k=0;
for i in `cat input.txt`
do
  echo $(( $i + $j + $k )) >> input2.txt
  k=$j
  j=$i
done
j=0;k=0
for i in `tail -n +3 input2.txt`
do
  [ "$i" -gt "$j" ] && k=$(($k+1))
  j=$i
done
echo $(($k-1))
