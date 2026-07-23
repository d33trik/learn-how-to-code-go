# Race conditions and the race detector

This exercise demonstrates a data race on a shared `int` variable accessed by multiple goroutines. It uses `runtime.Gosched` to make the race more visible and shows how to detect it with the `-race` flag.
