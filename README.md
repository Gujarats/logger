# Logger
A very simple logger for debugging purposes.

## Installation
Make sure you have setup Go in your environtment 

```shell
go get github.com/gujarats/logger
```

## Usage
It is very simple you can just call it : 

```go
logger.Debug("prefix::", "Hello World logger")
```
This will output

```shell
2017-11-14::8:00PM::prefix :: Hello World logger
```

That's it. it would print the log message with specific format time. And save it to a `logger,log` file
