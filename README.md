# Line counter

lc - line counter command line utilitary

## Setup

If you have go (golang.org) on your system:
```
    $ go get github.com/ademilly/lc
```

Else... yeah no, I'm not bothering serving up binaries for every arch :D

## Usage

Returns line count of file content provided through absolute path via command line argument or through stdin using the `-` placeholder.

```
    $ lc /path/to/some/file
    Output: line count
    $ cat /path/to/some/file | lc -
    Output: line count
```

## Next steps

- Aggregative line counting supporting arbitrary number of file provided through command line