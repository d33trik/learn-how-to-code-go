# Streaming JSON with NewEncoder

This exercise defines a `user` struct and uses `json.NewEncoder` to stream a slice of structs directly to `os.Stdout`. It demonstrates an alternative to `json.Marshal` for writing JSON output without storing it in a variable first.
