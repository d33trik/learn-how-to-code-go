# Custom sorting with sort.Interface

This exercise defines a `User` struct and a `Users` slice type that implements `Len` and `Swap` from `sort.Interface`. It demonstrates how to create custom sorting orders by embedding `Users` in `ByAge` and `ByLast` types that each provide a different `Less` method.
