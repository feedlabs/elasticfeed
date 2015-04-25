package common

type ElasticfeedConfig struct {
	ElasticfeedDebug       bool              `mapstructure:"elasticfeed_debug"`
	ElasticfeedForce       bool              `mapstructure:"elasticfeed_force"`
}
