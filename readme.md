across
======
Across automatically compiles your program
for all supported go OS's and architectures!

If I'm missing any OS's or architectures, open
a pull request!

Install
-------
`go get github.com/lordrusk/across`

How?
----
Run `across` in the programs directory.

`-h` Show help menu

`-v` Set version for binaries

`-m` Set custom binary information

'-p' Set the number of parallel compiles

`c` Set the number of parallel compiles to the number of cores (overwrites `-p`)

Examples
--------
[Goscrape](https://github.com/lordrusk/goscrape) uses across for its releases.

FQA - Frequently Questioned Answers
-----------------------------------
- Why does ______ fail to compile?
	I know as much as you. Figure it out.
	If it ends up being an across issue,
	feel free to open a pull request!
