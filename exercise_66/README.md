# Closing channels and ranging

This exercise demonstrates how to close a channel and iterate over its values using `range`. A goroutine sends a sequence of integers and then calls `close` to signal no more values will be sent, while the receiver uses `range` to read until the channel is closed.
