# Toscale Env
A simple env wrapper developed on top of [spf13/viper](https://github.com/spf13/viper)

## Installation
```shell
go get github.com/Toscale-platform/toscale-env
```

## Usage
Add your application configuration to your .env file in the root of your project:
```dotenv
TOKEN=ABC
PORT=8080
EXCHANGES=binance,bitfinex
DEBUG=true
```

Then in your Go app you can do something like:
```go
package main

import (
    env "github.com/Toscale-platform/toscale-env"
)

func main() {
    token := env.GetString("TOKEN")
    port := env.GetInt("PORT")
    exchanges := env.GetSlice("EXCHANGES")
    debug := env.GetBool("DEBUG")
}
```
