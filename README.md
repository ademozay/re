# re

Just retries tasks in a time manner.

## Install

```
go get -u github.com/ademozay/re
```

## How It Works

`Try`, at first, executes given task with no timer started. If an error occurs, it starts a `Ticker` and a `Timer` and executes the task every time `Ticker` ticks unless `Timer` is finish. If `Timer` finishes, `Try` returns the latest `error` returned from the task.

## Usage

This example tries to connect `:8000`  over `tcp` 6 times in 3 seconds.

```go
package main

import (
	"net"
    
    "github.com/ademozay/re"
)

func main() {
	var conn net.Conn
    
	err := re.Try(func() error {
		c, err := net.Dial("tcp", ":8000")
		if err != nil {
			return err
		}
		
		conn = c
		return nil
	}, time.Millisecond * 500, time.Second * 3)

	if err != nil {
		log.Fatal(err)
	}
    
    // ...
}
```