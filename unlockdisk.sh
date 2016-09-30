#! /bin/bash

WORDLIST=$(cat ../wordlist.txt)
mkdir mntScss
for KEY in $WORDLIST; do
	echo -n "$KEY" | sudo cryptsetup --key-file="-" luksOpen /dev/sda5 theDisk && sudo mount /dev/sda5 mntScss && echo "$KEY" > ../success.txt && break
done
