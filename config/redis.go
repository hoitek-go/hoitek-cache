package config

import (
	"fmt"
)

type Redis struct {
	Version  string
	Host     string
	Port     int
	Password string
	Database int
}

func (r Redis) GetHost() string {
	return r.Host
}

func (r Redis) GetPort() int {
	return r.Port
}

func (r Redis) GetPassword() string {
	return r.Password
}

func (r Redis) GetDatabase() int {
	return r.Database
}

func (r Redis) GetAddress() string {
	return r.Host + ":" + fmt.Sprintf("%d", r.Port)
}

func (r Redis) GetURI() string {
	return "redis://" + r.GetAddress()
}

func (r Redis) GetDialOptions() []string {
	return []string{
		"redis://" + r.GetAddress(),
		r.GetPassword(),
	}
}

func (r Redis) GetPoolOptions() []string {
	return []string{
		r.GetPassword(),
	}
}

func (r Redis) GetPoolSize() int {
	return 10
}
