# PassMaster

Password generator written in Go.

Produces three options of passwords between 14 and 18 characters in length.

Passwords contain uupercase and lowercase alpha-numeric characters with the addition of `!?()=+$` characters. 

At 14 characters in length, it would take roughly **10,016 years** to brute force this password at 2071.5 MH/s according to [this bruteforce calculator](https://www.proxynova.com/tools/brute-force-calculator/).