package config

type Configuration struct {
	Env       string
	Databases []Database
	Redis     struct {
		ClusterMode     bool     `yaml:"cluster_mode"`
		Addresses       []string `yaml:"addresses"`
		Password        string   `yaml:"password"`
		MaxRetries      int      `yaml:"max_retries"`
		PoolSizePerNode int      `yaml:"pool_size_per_node"`
		DB              int      `yaml:"db"`
	}
}

// Database 用來提供連線的資料庫數據
type Database struct {
	Type     string
	Name     string
	Username string
	Password string
	Address  string
	DBName   string
}

type MQ struct {
}

func New() Configuration {
	return Configuration{}
}
