# Hoitek-Cache

![Build Status](https://travis-ci.org/nock/nock.svg)
![Coverage Status](http://img.shields.io/badge/coverage-100%25-brightgreen.svg)
![contributions welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)

Hoitek-Cache is a simple multi-driver cache library for Golang. It is designed to be simple to use and easy to integrate into your project.

## Features 🔥
- Multi-driver support
- Easy to use
- Complete API Reference
- 100% Test Coverage

## Installation

You can install Hoitek-Cache using composer:

```bash
go get github.com/hoitek-go/hoitek-cache
```

## Usage

### Basic Usage

```go
package main

import (
	hoitekcache "github.com/hoitek-go/hoitek-cache"
	"github.com/hoitek-go/hoitek-cache/config"
	"github.com/hoitek-go/hoitek-cache/drivers"
	"github.com/hoitek-go/hoitek-cache/versions"
	"log"
)

func main() {
	// Create a new cache instance
	cache := hoitekcache.Connect(hoitekcache.Driver[drivers.Redis, config.Redis]{
		Config: config.Redis{
			Version:  versions.REDIS_V7,
			Host:     "localhost",
			Port:     6379,
			Password: "",
			Database: 0,
		},
	})
	defer cache.Close()

	// Set a value
	cache.Set("foo", "bar")

	// Get a value
	value, err := cache.Get("foo")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(value)

	// Delete a value
	err = cache.Delete("foo")
	if err != nil {
		log.Fatal(err)
	}
}
```

### Redis Driver Support

```go
cache := hoitekcache.Connect(hoitekcache.Driver[drivers.Redis, config.Redis]{
    Config: config.Redis{
        Version:  versions.REDIS_V7,
        Host:     "localhost",
        Port:     6379,
        Password: "",
        Database: 0,
    },
})
```

### MongoDB Driver Support

```go
cache := hoitekcache.Connect(hoitekcache.Driver[drivers.MongoDB, config.MongoDB]{
    Config: config.MongoDB{
       Host: "localhost",
       Port:     27017,
       Username: "root",
       Password: "111111",
       Database: "cachedb",
    },
})
```

### API Reference

```go
Connect[T1 ports.DriverType, T2 ports.ConfigType](driver Driver[T1, T2]) *Gach[T1, T2]
Set(key string, value interface{}) error
Get(key string) (interface{}, error)
Delete(key string) error
Exists(key string) (bool, error)
Expire(key string, seconds int) error
TTL(key string) (int, error)
Flush() error
Close() error
Ping() error
```

## Run Tests

~~~bash  
  make test
~~~

## Export Test Coverage

~~~bash  
  make testcov
~~~

## Tech Stack
**Server:** Golang

## Licence
[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)   

