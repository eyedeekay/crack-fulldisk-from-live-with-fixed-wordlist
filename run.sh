#! /bin/sh
./a.out $1 | sort --dictionary-order --ignore-nonprinting --ignore-leading-blanks --parallel=2 --unique
./unlockdisk.sh
