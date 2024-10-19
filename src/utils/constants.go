package utils

type EnvironmentTypes struct {
	Dev  bool
	Prod bool
}

type Environments struct {
	Port string
	Host string
	Type EnvironmentTypes
}

type ServerConstants struct {
}

var Constants ServerConstants
var Env Environments
