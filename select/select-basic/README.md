## About this module
The select statement in Go is a powerful feature that allows a goroutine to wait on multiple communication operations.
It's primarily used to handle multiple channels concurrently in a clean, organized way.
A select blocks until one of its cases can run, then it executes that case. If multiple operations are ready it chooses
one pseudo-randomly. It's commonly used in conjunction with goroutines and channels to handle asynchronous tasks.


## Instructions
```shell
go mod init github.com/oskiegarcia/select
```
### 2. Creat main.go

### 3. Run the program
```shell
go run .
```


