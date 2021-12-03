#!/bin/env zsh

nbRows=$(wc -l input.txt | cut -d' ' -f1)
nbChar=$(( $(head -n1 input.txt | wc -c| cut -d' ' -f1) - 2 ))
s1='';s2=''
for i in {0..$nbChar}
do
  k=0
  while read LINE; do
    n=${LINE:$i:1}
    [ "$n" -eq "1" ] && k=$(($k+1))
  done < input.txt
  if [ $(($k*2)) -gt $nbRows ]; then
    s1=$(echo $s1'1') && s2=$(echo $s2'0')
  else
    s1=$(echo $s1'0') && s2=$(echo $s2'1')
  fi
done
echo $(( 2#$s1 * 2#$s2 ))
