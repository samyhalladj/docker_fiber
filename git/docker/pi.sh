#!/bin/sh

START=1
END=1000

TOTAL=0
INSIDE=0

while [[ $START -le $END ]]
do
    x=$RANDOM
    x=`echo "$x / 32767"| bc -l`

    y=$RANDOM
    y=`echo $y / 32767 | bc -l`

    xxyy=`echo "($x * $x + $y * $y) <= 1 " | bc -l`

    if [[ $xxyy -eq 1 ]]; then
        INSIDE=$((INSIDE+1))
    fi

    TOTAL=$((TOTAL+1))

    START=$(($START + 1))
done

echo "$INSIDE*4/$TOTAL" | bc -l
