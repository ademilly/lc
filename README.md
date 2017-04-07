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
If given multiple files, will return the aggregated count of lines.

```
    $ lc /path/to/some/file
    Output: line count
    $ cat /path/to/some/file | lc -
    Output: line count
    $ cat /path/to/some/file /path/to/some/other/file
    Output: some file line count + some other file line count
```

## Next steps

- Aggregative line counting supporting arbitrary number of file provided through command line