# Mutex for race condition prevention

This exercise demonstrates how to use `sync.Mutex` to protect a shared variable accessed by multiple goroutines. It fixes the data race from the previous exercise by locking and unlocking around the critical section, showing the correct way to synchronize concurrent access in Go.
