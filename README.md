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
logger.Debug("prefix :: ", "Hello World logger")
logger.Debug("prefix2 :: ", "testing")
logger.Debug("prefix3 :: ", "foo bar")
```
This will output

[![Example Output](resource/logger.png)]()

That's it. it would print the log message with specific format time. And save it to a `logger,log` file
