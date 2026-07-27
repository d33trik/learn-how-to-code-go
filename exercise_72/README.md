# Custom error types with struct methods

This exercise demonstrates how to create custom error types in Go by defining a struct that implements the `error` interface. A `coordinatesError` struct holds the invalid values and an underlying error, and its `Error()` method returns a formatted message describing the problem.
