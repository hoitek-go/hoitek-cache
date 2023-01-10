package hoitekcache

import (
	"github.com/hoitek-go/hoitek-cache/config"
	"github.com/hoitek-go/hoitek-cache/ports"
)

type Driver[T1 ports.DriverType, T2 ports.ConfigType] struct {
	instance T1
	Config   T2
}

type Gach[T1 ports.DriverType, T2 ports.ConfigType] struct {
	Driver Driver[T1, T2]
}

func Connect[T1 ports.DriverType, T2 ports.ConfigType](driver Driver[T1, T2]) *Gach[T1, T2] {
	config.Global = driver.Config
	return &Gach[T1, T2]{
		Driver: driver,
	}
}

func (g *Gach[T1, T2]) Set(key string, value interface{}) error {
	driver := g.GetDriver()
	return driver.Set(key, value)
}

func (g *Gach[T1, T2]) Get(key string) (interface{}, error) {
	driver := g.GetDriver()
	return driver.Get(key)
}

func (g *Gach[T1, T2]) Delete(key string) error {
	driver := g.GetDriver()
	return driver.Delete(key)
}

func (g *Gach[T1, T2]) Exists(key string) (bool, error) {
	driver := g.GetDriver()
	return driver.Exists(key)
}

func (g *Gach[T1, T2]) Expire(key string, seconds int) error {
	driver := g.GetDriver()
	return driver.Expire(key, seconds)
}

func (g *Gach[T1, T2]) TTL(key string) (int, error) {
	driver := g.GetDriver()
	return driver.TTL(key)
}

func (g *Gach[T1, T2]) Flush() error {
	driver := g.GetDriver()
	return driver.Flush()
}

func (g *Gach[T1, T2]) Close() error {
	driver := g.GetDriver()
	return driver.Close()
}

func (g *Gach[T1, T2]) Ping() error {
	driver := g.GetDriver()
	return driver.Ping()
}

func (g *Gach[T1, T2]) GetDriver() T1 {
	return g.Driver.instance
}

func (g *Gach[T1, T2]) GetConfig() T2 {
	return g.Driver.Config
}

func (g *Gach[T1, T2]) SetDriver(driver T1) {
	g.Driver.instance = driver
}

func (g *Gach[T1, T2]) SetConfig(config T2) {
	g.Driver.Config = config
}

func (g *Gach[T1, T2]) GetDriverType() ports.DriverType {
	return g.Driver.instance
}

func (g *Gach[T1, T2]) GetConfigType() ports.ConfigType {
	return g.Driver.Config
}
