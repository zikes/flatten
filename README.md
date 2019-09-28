# Flatten

Flattens an arbitrarily nested array of `Type` into a flat array of `Type`,
e.g. `[[1,2,[3]],4]` -> `[1,2,3,4]`.

Currently only the `int` type is supported via the `flatten.Int` function.
Additional types may be supported in the future.

# Example Usage

**`func Int([]interface{}) ([]int, error)`**

```go
input := []interface{}{
  []interface{}{
    1,
    2,
    []interface{}{
      3,
    },
  },
  4,
}

ints := flatten.Int(input)
fmt.Println(ints) // [1 2 3 4]
```
