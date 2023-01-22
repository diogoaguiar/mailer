package configs

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var possibleConfigs = []string{".env", "config/.env"}

// Config is the configuration for the mailer
type Config struct {
	Strategy string
	Interval int
	Sender   string
	Host     string
	Port     int
	Username string
	Password string
}

type value struct {
	key string
	val string
}

func (v *value) string() string {
	return v.val
}

func (v *value) int() int {
	intValue, err := strconv.Atoi(v.val)
	if err != nil {
		log.Fatalf("could not parse int value of \"%s\": %s\n", v.key, err)
	}
	return intValue
}

func (v *value) required() *value {
	if v.val == "" {
		log.Fatalf("missing required environment variable: %s\n", v.key)
	}
	return v
}

func (v *value) optional(defaultValue string) *value {
	if v.val == "" {
		v.val = defaultValue
	}
	return v
}

func valueOf(key string) *value {
	return &value{key: key, val: os.Getenv(key)}
}

// Load loads the configuration from the .env file and the environment
func Load() *Config {
	godotenv.Load(locateConfig())

	c := &Config{}

	c.Strategy = valueOf("STRATEGY").optional("sequential").string()
	c.Interval = valueOf("SEND_INTERVAL").optional("5").int()
	c.Sender = valueOf("SENDER").optional("smtp").string()

	switch c.Sender {
	case "smtp":
		c.Host = valueOf("SMTP_HOST").required().string()
		c.Port = valueOf("SMTP_PORT").required().int()
		c.Username = valueOf("SMTP_USERNAME").required().string()
		c.Password = valueOf("SMTP_PASSWORD").required().string()
	default:
		log.Fatalln("invalid sender:", c.Sender)
	}

	return c
}

func locateConfig() string {
	for _, config := range possibleConfigs {
		if _, err := os.Stat(config); os.IsNotExist(err) {
			continue
		}

		return config
	}

	return ""
}
