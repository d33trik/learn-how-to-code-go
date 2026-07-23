# Atomic operations for race-free concurrency

This exercise demonstrates how to use `sync/atomic` functions to safely update a shared variable accessed by multiple goroutines. It shows an alternative to `sync.Mutex` for simple operations, using `atomic.AddInt64` to increment a counter without explicit locking.
