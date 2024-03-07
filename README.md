# go-unit

`go-unit` provides some useful functions to handle unit-related processing.

## Unit Prefix

- `PrefixParse(s string, prefixMode PrefixMode) (val float64, err error)` can converts:
  - "1024 Ki" -> 1048576
  - "1.2K" -> 1200
  - ...
- `PrefixFormat(val float64, prefixMode PrefixMode, prefixFormatOptionFns ...prefixFormatOptionFn) string` can converts:
  - `999 * 1000 * 1000` -> "999.000 M"
  - `2048` -> "2Ki"
  - `0.568` -> "568m"
  - ...
