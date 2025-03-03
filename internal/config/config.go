package config

var (
    defaultServerConfig = ServerConfig {
        Port: ":8080",
        Addr: "localhost",
    } 
)

type Config struct {
    Server  ServerConfig
}

type ServerConfig struct {
    Port    string
    Addr    string
}

func New() *Config {

    return &Config{
        Server: defaultServerConfig,
    }
}
