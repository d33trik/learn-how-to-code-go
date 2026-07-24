# Buffered channels

This exercise demonstrates how to use buffered channels in Go. It creates a channel with a buffer size of 1 using `make(chan int, 1)`, allowing a send operation to complete without a concurrent receiver, which differs from unbuffered channels that block until the value is received.
