# joli

`joli` joins lines of the given input file.

## Usage

```
~ joli -h
Usage of joli:
  -b int
        buffer size of a line (default 65536)
  -i string
        path to an input file
  -o string
        path to an output file
  -s string
        line separator (default " ")
```

## Examples

```shell
~ echo -e "a\nb\nc" | joli
a b c
```

```shell
~ echo -e "a\nb\nc" > in.txt
~ joli -i in.txt -o out.txt -s ""
~ cat out.txt
abc
```
