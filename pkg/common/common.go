package common

// ServerConfig ...
type ServerConfig struct {
	DBConfig         DBConfig         `yaml:"db_config"`
	GRPCConfig       GRPCConfig       `yaml:"grpc_config"`
	PrometheusConfig PrometheusConfig `yaml:"prometheus_config"`
}

// DBConfig ..
type DBConfig struct {
	Duration uint32 `yaml:"duration"`
}

// GRPCConfig ..
type GRPCConfig struct {
	Address string `yaml:"address"`
}

// PrometheusConfig ..
type PrometheusConfig struct {
	Address string `yaml:"address"`
}
