# Tarkov API

A quick and dirty Golang interface for the Escape from Tarkov API.

## Installation

You can install it via `go get` like so:

`$ go get -u github.com/mbStavola/tarkov-api`

Or you can simply reference the library in code:

```go
import "github.com/mbStavola/tarkov-api"
```

## Usage

Everything is provided off the `TarkovAPI` struct:

```go
import "github.com/mbStavola/tarkov-api"

// ...

api := tarkov.NewTarkovAPI("this-is-your-session-id")
```

It might be a bit tricky to get your session ID, but if you can use Wireshark to analyze traffic you'd find it as `PHPSESSID` in one of the requests.

Each individual API is split off of the `TarkovAPI` struct like so:

- Locations => `api.Locations()`
- Economy => `api.Economy()`
- Items => `api.Items()`

and so on. Within these scoped APIs, you can fetch the specific data you need:

- Ammunition types => `api.Items().AmmunitionTypes()`
- Ammunition => `api.Items.Ammunition("9x18")`

Putting this all together, here is how you'd query location information: 

```go
import "github.com/mbStavola/tarkov-api"

func main() {
    // Create an API instance with your session ID
    api := tarkov.NewTarkovAPI("this-is-your-session-id")

    // Access the Locations API and search locations by name
    searchedLocations, err := api.Locations().LocationsByName("Interchange")

    // ...
}
```

You can find various examples in the [examples](./examples) directory.