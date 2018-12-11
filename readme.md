# Rando

A pseudo random number generator written in Go.

## Invocation

```bash
$ rando [flags]
```

### For random numbers

| Flag |  Default | Action |
|---|:-:|---|
| -c, --columns |  1 | The number of columns to generate. |
| --max | 100 | The upper limit of the random numbers to generate. |
| --min | 1 | The lower limit of the random numbers to generate. |
| -n,  --number | 1 | The number of random numbers to generate. |
| -s, --sort | false | Sort the random numbers? |
| -t, --trials | 1 | The number of trials to generate. |
| -u, --unique | false | Should the random numbers be unique? |

### For random strings

```bash
$ rando --str=1 [flags]
```

| Flag |  Default | Action |
|---|:-:|---|
| -c, --columns |  1 | The number of columns to generate. |
| -l, --length |  16 | The length of strings to generate. |
| -n,  --number | 1 | The number of random numbers to generate. |
| -str | false | Generate random strings instead of numbers numbers? |
| -t, --trials | 1 | The number of trials to generate. |


### Example 1

Generate twenty random numbers between 1 and 1000, displayed 4-wide.

```bash
$ rando --max=1000 --number=20 --columns=4

   81  321   80  434
  122  865  603  551
  570  527  348   52
  167  463  750  701
  947  313  651  668
```

### Example 2

Generate ten Loto 6/49 qiick picks, sorting each row.

```bash
$ rando -n=6 --max=49 -c=6 -u=true -s=true -t=10

  8 16 21 36 40 45
  4 16 27 35 38 47
  1 16 24 38 42 48
  3 18 24 29 39 43
  7 19 21 24 36 42
  8  9 26 30 36 47
 10 17 32 34 42 45
  5  8 16 29 44 48
  6 15 28 29 36 48
  8  9 25 28 31 43
```

### Example 3

Generate five random strings of length 10.

```bash
$ ./rando --str=1 -l=10 -n=5

oikX_STRqnjCkQ==
czyXWZ2yEVbldA==
or6ZN25pq0XcZg==
JvWqAVDpDsN5Iw==
EOTT4TrinYwt1Q==
```


