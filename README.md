# go-unit

`go-unit` provides some useful functions to handle unit-related processing.

- `Parse(s string, mode Mode) (val float64, err error)` can converts:
  - "1024Ki" -> 1048576
  - "1.2K" -> 1200
  - ...
- `Format(val float64, mode Mode, formatOptions ...FormatOption) string` can converts:
  - `999 * 1000 * 1000` -> "999.000 M"
  - `2048` -> "2Ki"
  - `0.568` -> "568m"
