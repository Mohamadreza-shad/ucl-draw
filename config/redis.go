package config

type Redis struct {
	URI    string
	Master Master
}

type Master struct {
	Name string
}

func RedisURI() string {
	if cfg.Redis.URI == "" {
		return "redis://:123456@0.0.0.0:6379"
	}
	return cfg.Redis.URI
}

func RedisMasterName() string {
	return cfg.Redis.Master.Name
}
