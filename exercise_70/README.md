# Wrapping errors with fmt.Errorf

This exercise demonstrates how to return descriptive errors from a function using `fmt.Errorf`. A `toJSON` helper wraps `json.Marshal`, returning a `([]byte, error)` pair and adding context to the error message before passing it back to the caller.
