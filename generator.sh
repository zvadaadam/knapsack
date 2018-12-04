#!/bin/bash

# LOAD OPTIONS
while getopts a:n:N:m:W:C:k:d:- opt
do
	case "$opt" in
	a)
	    algo="$OPTARG" # -a (algorithm name)
	;;

	n)
		num_items="$OPTARG" # -n (number of items)
	;;

	N)
		num_instances="$OPTARG" # -N (number of instances)
	;;

	m)
		capacity_weights_ratio="$OPTARG" # -m (ratio between knapsack capacity and sum of weights)
	;;

	W)
		max_weight="$OPTARG" # -W (max weight)
	;;

	C)
		max_price="$OPTARG" # -C (max price)
	;;

	k)
	    exponent="$OPTARG" # -k (exponent k)
	;;
	d)
		type="$OPTARG" # -d (-1 = more small items, 0 = equlibrium, 1 = more big items)

	;;
	esac
done
shift "$(( OPTIND - 1 ))"


echo "__________________________________"
echo "algorithm: $algo"
echo "num items: $num_items"
echo "num instances: $num_instances"
echo "ratio: $capacity_weights_ratio"
echo "max weight: $max_weight"
echo "max price: $max_price"
echo "exponent: $exponent"
echo "type: $type"
echo "__________________________________"

rm -rf input-*
rm -rf data.csv


counter=0
printf -- "instance,duration\n" >> data.csv

for i in `seq 7000 500 15000`; do
    ./knapgen/knapgen -n $num_items -N $num_instances -m $capacity_weights_ratio -W $i -C $max_price -k $exponent -d $type > "input/input-${i}"
    #./knapgen/knapgen -n 23 -N 15 -m 0.2 -W $i -C 500 -k 1 -d 0 > "input/input-${i}"

    echo "Running Knapsack on input/input-${i}"

    printf -- "${counter},$(./main -algorithm ${algo} < input/input-${i})\n" >> data.csv

    $(( counter++ ))

done
