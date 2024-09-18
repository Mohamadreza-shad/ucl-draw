package config

type Postgres struct {
	URL string
}

func PostgresURL() string {
	if cfg.Postgres.URL == "" {
		return "postgresql://postgres:postgres@localhost:5432/user_db?sslmode=disable"
	}
	return cfg.Postgres.URL
}