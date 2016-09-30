Try a series of passwords on a disk you posess
==============================================

This is a little script which is capable of looping over the contents of a 
wordlist file sequentially. This is because I unabashedly forgot where the
capital and lower case letters are in the disk encryption password for my
desktop.

How to use it?
--------------

This will change, but for now, you put a file named wordlist.txt in the parent
directory and run the script with no arguments. To generate a wordlist, use
a.out(or compile genwl.go with a name of your choice) and pass it a likely
password as a parameter. It will, very inelegantly, generate every possible
variation of capital and lower-case letters in that password, and eventually
other types of variants(stretches, omissions, general fuzziness) if you ask it
to. Once you have that, you can sort it for uniqueness and pass it to the script,
which will attempt to mount the device with each password successively, until
finally putting the correct password in a folder called "success.txt"
