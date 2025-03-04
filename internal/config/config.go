package config

var (
    defaultServerConfig = ServerConfig {
        Port: ":8080",
        Addr: "192.168.1.71",
    } 

    defaultDatabaseConfig = DatabaseConfig {
        Port:       "5432",
        User:       "ek_admin",
        Name:       "ek_db",
        Password:   "1234",
        Driver:     "postgres",
    }
)

type Config struct {
    Server      ServerConfig
    Database    DatabaseConfig
}

type ServerConfig struct {
    Port    string
    Addr    string
}

type DatabaseConfig struct {
    Name        string
    Password    string
    User        string
    Port        string
    Driver      string
}

func New() *Config {

    return &Config{
        Server: defaultServerConfig,
    }
}
