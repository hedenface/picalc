# picalc

> i wanted a better fraction than 22/7

## go

```
$ make go
go build -o picalcgo picalc.go

$ time ./picalcgo 
pi = 3.141593
(22.000000 / 7.000000) = (3.142857), difference from pi = 0.001264
(99733.000000 / 31746.000000) = (1.000000), difference from pi = 0.000000

real    0m9.975s
user    0m9.982s
sys 0m0.011s
```

## c

```
$ make c
gcc -o picalcc picalc.c

$ time ./picalcc
pi = 3.141593
(22.000000 / 7.000000) = (3.142857), difference from pi = 0.001264
(99733.000000 / 31746.000000) = (1.000000), difference from pi = 0.000000

real    0m18.991s
user    0m18.970s
sys 0m0.007s
```

## python

```
$ time python3 picalc.py 

pi = 3.14159265358979
(22.0 / 7.0) = (3.142857142857143), difference from pi = 0.0012644892673527863
(99733 / 31746) = (1.0), difference from pi = 1.1997148607889585e-08

real    11m48.517s
user    11m47.972s
sys 0m0.102s
```

## better fraction

well, two compiled languages give me a better fraction `99733/31746`, and the interpreted language gives me `(99733/31746)`.

plugging this number into a calculator, this seems only slightly better than `22/7`, so that's cool.

using floating points is nowhere near where i'd thought it would have been in 2024. the future is crazy.
