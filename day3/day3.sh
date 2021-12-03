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

cp input.txt input.bak
nbChar=$(( $(head -n1 input.txt | wc -c| cut -d' ' -f1) - 2 ))
s1=''
for i in {0..$nbChar}
do
  nb=$(grep "^${s1}1" input.txt | wc -l | cut -d' ' -f1)
  nbRows=$(wc -l input.txt | cut -d' ' -f1)
  if [ $(($nb*2)) -ge $nbRows ]; then
    s1=$(echo $s1'1')
  else
    s1=$(echo $s1'0')
  fi
  sed -i /^$s1/!d input.txt
  [ $(wc -l input.txt | cut -d' ' -f1) -le 1 ] && break
done
r1=`cat input.txt`

cp input.bak input.txt
nbChar=$(( $(head -n1 input.txt | wc -c| cut -d' ' -f1) - 2 ))
s1=''
for i in {0..$nbChar}
do
  nb=$(grep "^${s1}1" input.txt | wc -l | cut -d' ' -f1)
  nbRows=$(wc -l input.txt | cut -d' ' -f1)
  if [ $(($nb*2)) -lt $nbRows ]; then
    s1=$(echo $s1'1')
  else
    s1=$(echo $s1'0')
  fi
  sed -i /^$s1/!d input.txt
  [ $(wc -l input.txt | cut -d' ' -f1) -le 1 ] && break
done
r2=`cat input.txt`

echo $(( 2#$r1 * 2#$r2 ))

mv input.bak input.txt
