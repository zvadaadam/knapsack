#!/bin/bash

# LOAD OPTIONS
while getopts o:a:n:N:m:W:C:k:d:- opt
do
	case "$opt" in
	o)
	    optimized="$OPTARG" # -a (algorithm name)
	;;

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
echo "optimized: $optimized"
echo "algorithm: $algo"
echo "num items: $num_items"
echo "num instances: $num_instances"
echo "ratio: $capacity_weights_ratio"
echo "max weight: $max_weight"
echo "max price: $max_price"
echo "exponent: $exponent"
echo "type: $type"
echo "__________________________________"

rm -rf input/input-*
rm -rf instance_data.csv


if [ $optimized = "weight" ]
then
    counter=0
    printf -- "________WEIGHT________"
    if [ $algo == "heuristic" ]
    then
        printf -- "weight,duration,error\n" >> instance_data.csv
    else
        printf -- "weight,duration\n" >> instance_data.csv
    fi
    for i in `seq 1000 1000 30000`; do
        ./knapgen/knapgen -n $num_items -N $num_instances -m $capacity_weights_ratio -W $i -C $max_price -k $exponent -d $type > "input/input-${i}"
        #./knapgen/knapgen -n 23 -N 15 -m 0.2 -W $i -C 500 -k 1 -d 0 > "input/input-${i}"

        echo "Running Knapsack on input/input-${i}"

        printf -- "${i},$(./main -algorithm ${algo} < input/input-${i})\n" >> instance_data.csv

        $(( counter++ ))

    done
fi

if [ $optimized = "price" ]
then
    counter=0
    printf -- "________PRICE________"
    if [ $algo == "heuristic" ]
    then
        printf -- "price,duration,error\n" >> instance_data.csv
    else
        printf -- "price,duration\n" >> instance_data.csv
    fi
    for i in `seq 1000 1000 30000`; do
        ./knapgen/knapgen -n $num_items -N $num_instances -m $capacity_weights_ratio -W $max_weight -C $i -k $exponent -d $type > "input/input-${i}"

        echo "Running Knapsack on input/input-${i}"

        printf -- "${i},$(./main -algorithm ${algo} < input/input-${i})\n" >> instance_data.csv

        $(( counter++ ))

    done
fi

if [ $optimized = "ratio" ]
then
    counter=0
    printf -- "________RATIO________"
    if [ $algo == "heuristic" ]
    then
        printf -- "ratio,duration,error\n" >> instance_data.csv
    else
        printf -- "ratio,duration\n" >> instance_data.csv
    fi
    for i in `seq 0.1 0.1 1`; do
        ./knapgen/knapgen -n $num_items -N $num_instances -m $i -W $max_weight -C $max_price -k $exponent -d $type > "input/input-${i}"

        echo "Running Knapsack on input/input-${i}"

        printf -- "${i},$(./main -algorithm ${algo} < input/input-${i})\n" >> instance_data.csv

        $(( counter++ ))

    done
fi

if [ $optimized = "exponent" ]
then
    counter=0
    printf -- "________EXPONENT________"
    if [ $algo == "heuristic" ]
    then
        printf -- "exponent,duration,error\n" >> instance_data.csv
    else
        printf -- "exponent,duration\n" >> instance_data.csv
    fi
    for i in `seq 0 0.5 15`; do
        ./knapgen/knapgen -n $num_items -N $num_instances -m $capacity_weights_ratio -W $max_weight -C $max_price -k $i -d $type > "input/input-${i}"

        echo "Running Knapsack on input/input-${i}"

        printf -- "${i},$(./main -algorithm ${algo} < input/input-${i})\n" >> instance_data.csv

        $(( counter++ ))

    done
fi
