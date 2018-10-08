# Word Counter

Simple API that counts the words per sentence and the total number of letters in
the given input.

## Build

To build `wordcounter`, you can just do a normal `go build`.

```
go build
```

## Run

After you run `go build`, a `wordcounter` executable will be written.
`wordcounter` listens on port 9090.

```
$ ./wordcounter
2018/10/08 19:12:04 Listening on http://localhost:9090
```

## Testing API

You can connect to the API using `curl`. There are two test scripts in this
project that can be modified to test the API.

```
$ ls -l *.sh
-rwxrwxr-x. 1 jesusr jesusr 258 Oct  8 10:12 letter_counter.sh*
-rwxrwxr-x. 1 jesusr jesusr 263 Oct  8 10:15 word_per_sentence.sh*
```

Here is an example of the total letter count API. Pass in the `Content-type` and
`Accept` headers and the input data. The input is a simple key:value pair but
the key is currently ignored. The important portion is the sentence string.

```bash
#!/bin/bash -e

req='{"somekey":"This is a sentence. But is this? I KNOW THIS IS!" }'

curl \
    -k \
    -X POST \
    -H "Content-type: application/json" \
    -H "Accept: application/json" \
    -d "$req" \
    "http://localhost:9090/total_letter_count"
```
