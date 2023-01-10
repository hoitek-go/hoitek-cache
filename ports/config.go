package ports

import "github.com/hoitek-go/hoitek-cache/config"

type ConfigType interface {
	config.Redis | config.MongoDB | any
}
