# Log parser exercise

Write a command that reads a log stream and reports a small summary. Use the following input and output contract as a target while you build the implementation.

Learn: buffered I/O, parsing, errors with context, flags, deterministic output, and focused tests. Keep the first version in the standard library.

## Input contract

Use newline-delimited JSON (JSONL): one JSON object per physical line. All happy-path examples use the same fields and value formats:

```jsonl
{"timestamp":"2026-10-27T09:00:00Z","level":"INFO","message":"service started"}
```

| Field | Required value |
| --- | --- |
| `timestamp` | A string containing a valid UTC timestamp in `YYYY-MM-DDTHH:MM:SSZ` format. |
| `level` | One of the case-sensitive strings `DEBUG`, `INFO`, `WARN`, or `ERROR`. |
| `message` | A nonempty string. Spaces and JSON escapes are allowed. |

Field order does not affect meaning. Ignore additional fields. Missing fields, `null` values, incorrect types, invalid timestamps, and unsupported levels make a record malformed. Blank or whitespace-only lines are malformed too. Do not interpret the message contents as fields.

Read standard input by default; accept `-file PATH` to read a file instead. Both sources must produce the same results for the same bytes. Support LF and CRLF line endings, and process the final record even if it has no trailing newline. Process records incrementally so memory use does not grow with the total stream size. Set a maximum record size of 1 MiB (1,048,576 bytes, excluding the line ending).

## Output contract

At end of input, write these six lines to standard output in exactly this order, including a final newline:

```text
records=0
malformed=0
DEBUG=0
INFO=0
WARN=0
ERROR=0
```

`records` counts valid records only. Each valid record increments exactly one level count. Always include every level, even when its count is zero. Count duplicate records separately; timestamps need not be sorted. Output does not depend on JSON field order or input record order.

Return exit code `0` when input is read successfully and all records are valid. For malformed records within the size limit, skip the record, increment `malformed`, and continue. Return exit code `1` after printing the summary if any records were malformed. For each malformed record, write `line N: malformed record` to standard error, where `N` is its one-based physical line number.

A file-open failure, read failure, or oversized record is fatal: write an explanatory error to standard error, return exit code `1`, and do not print a summary. Include the line number for an oversized record. Stop reading when a fatal error occurs.

## Working with the fixture files

The inputs and expected results below are available as actual files in [`testdata/`](testdata/). Each case has a `.jsonl` input, `.stdout` and `.stderr` expected outputs, and an `.exitcode` file. Empty expected stderr files contain zero bytes. The `empty.jsonl` input is also zero bytes, and `no-final-newline.jsonl` ends immediately after the closing brace.

Once you have implemented the command, run it from the repository root with either a file argument or redirected standard input:

```sh
go run ./cmd/log-parser -file cmd/log-parser/testdata/mixed.jsonl
go run ./cmd/log-parser < cmd/log-parser/testdata/mixed.jsonl
```

To compare the output against the fixture:

```sh
go run ./cmd/log-parser -file cmd/log-parser/testdata/mixed.jsonl > /tmp/log-parser-mixed.stdout
diff -u cmd/log-parser/testdata/mixed.stdout /tmp/log-parser-mixed.stdout
```

For exit-code checks, build and run the binary directly: `go run` can add its own error message and does not preserve every program exit code.