#!/bin/bash

# Did I use a AI to write this script? Yes.
# Why? Because it's bash.

wordList=("apple" "banana" "cherry" "date" "elderberry" "fig" "grape" "honeydew" "kiwi" "lemon" "mango" "nectarine" "orange" "papaya" "quince" "raspberry" "strawberry" "tangerine" "ugli" "watermelon")

for i in {1..20}; do
    randomWord=${wordList[$RANDOM % ${#wordList[@]}]}

    echo "Round #$i"
    ./leeter -Input "$randomWord"
    echo "------------------------------"
done
