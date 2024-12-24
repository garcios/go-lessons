## About this module

The `select` is a powerful control flow mechanism in Go that allows you to work with multiple channels simultaneously. 
In the case of a ticker, which periodically sends a value on a channel, select lets your program wait until either the 
ticker sends a value or another event occurs.

In this example, the Ticker sends a value on its C channel every second, and the select statement waits until a value is
available. When the Ticker ticks, "Tick!" is printed to the console.
Please remember to stop the ticker when you're done with it, to prevent a resource leak. This can be done with 
ticker.Stop(). Within a select statement, you can also handle multiple channels. For instance, you might have another 
stop channel that indicates when you should break the loop.

You can use the time.After function to implement a timeout along with a ticker. This function returns a channel 
that will send the current time after the specified duration.

## Instructions

### 1. Create a module
```shell
go mod init github.com/oskiegarcia/select-ticker
```

### 2. Creat main.go

### 3. Run the program
```shell
go run .
```


