# t-digest

A go implementation of
[the t-digest algorithm](https://github.com/tdunning/t-digest).

Based on
[spenczar's implementation](https://github.com/spenczar/tdigest/blob/master/tdigest.go).
Over 90% of the lines of code have been modified.

## Performance

Per benchmark, mean insertion time is ~46μs. Subtract `Rand`
benchmark from `TDigest_Add` benchmark to compute insertion time for your
machine.

```bash
go test ./... --test.bench=.
```

## Limitations

While much faster at adding new elements than spenczar's implementation, this
implementation has significant limitations.

- **No tests.** Please don't use this in production until I've added unit tests.
This was just a fun weekend project.
- Only supports adding elements of weight 1.
- Optimized for time-independent distributions.
- No support for merging TDigests.
- Optimized for my machine. It is possible certain choices, such as when to
switch from a binary search to a linear search, will be more optimal with
different thresholds on other machines. You'll have to test and edit these
constants yourself. It's just faster to make them compile time constants
than to leave these as variables.
