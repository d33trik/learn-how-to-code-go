# Select statement for channel multiplexing

This exercise demonstrates how to use the `select` statement to multiplex over multiple channels in Go. A goroutine sends values on one channel and signals completion on another, while the receiver uses `select` to handle either channel until the signal channel is closed.
