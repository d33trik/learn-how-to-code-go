# Comma-ok idiom for channels

This exercise demonstrates how to receive from a channel using the comma-ok idiom. It shows that a receive operation returns the zero value and `false` after the channel is closed, allowing the receiver to detect when no more values will be sent.
