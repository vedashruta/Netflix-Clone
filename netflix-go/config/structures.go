package config

import (
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	Config Env
	Db     *mongo.Database
	Movies *mongo.Collection
	Users  *mongo.Collection
)

type Env struct {
	Server struct {
		Port    string        `yaml:"port"`
		Timeout time.Duration `yaml:"timeout"`
	} `yaml:"server"`
	Database struct {
		URI     string        `yaml:"uri"`
		Db      string        `yaml:"db"`
		Timeout time.Duration `yaml:"timeout"`
	} `yaml:"database"`
}
