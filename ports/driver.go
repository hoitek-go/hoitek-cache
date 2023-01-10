package ports

import "github.com/hoitek-go/hoitek-cache/drivers"

type DriverType interface {
	drivers.Redis | drivers.MongoDB | any
	Set(key string, value interface{}) error
	Get(key string) (interface{}, error)
	Delete(key string) error
	Exists(key string) (bool, error)
	Expire(key string, seconds int) error
	TTL(key string) (int, error)
	Flush() error
	Close() error
	Ping() error
}
