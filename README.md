# bashir
Package bashir calls bash (using `-c`) once per value supplied, and returns command output (if any) as a string slice.

## Why call it that?
bashir is a syllabic abbreviation for "bash, input, repeat." English doesn't have many words with bash in it.

## Why do this at all?
There are times when pulling information from the command line into a Go utility is more expedient than other means. This solves that for me.
