# PassMaster

Simple Password generator written in Go.

Produces three options of passwords between 14 and 18 characters in length.

Default passwords (no flags) contain uupercase and lowercase alpha-numeric characters with the addition of `!?()=+$` characters.

**TODO: write explanation of flags & help**

At 14 characters in length, it would take roughly **10,016 years** to brute force this password at 2071.5 MH/s according to [this bruteforce calculator](https://www.proxynova.com/tools/brute-force-calculator/).


## TODO:
- Take flags for letter (`-a`), number (`-n`), and/or special characters(`-!`) and create passwords with only those characters
- Take flags for uppercase only (`-u`) and lowercase only (`-l`)
- Flag input can be `-an!u` in any order or `-a -u -! -u`
- `-u` and `-l` cannot be combined
- Create help (`-h` or `--help`), cannot be combined with above flags
- Incorrect flag input will write red error message to console and terminate the program