# Pointer receivers and interfaces

This exercise demonstrates how methods with pointer receivers interact with interfaces in Go. It shows that a pointer is required to satisfy an interface when the method set includes pointer receivers, and that Go automatically takes the address of an addressable value when calling such methods directly.
