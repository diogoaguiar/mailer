package configs

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const configFile = "config/.env"

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
	val string
}

func (v *value) string() string {
	return v.val
}

func (v *value) int() int {
	intValue, err := strconv.Atoi(v.val)
	if err != nil {
		panic(err)
	}
	return intValue
}

func (v *value) required() *value {
	if v.val == "" {
		panic(fmt.Errorf("configuration missing: %s", v.val))
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
	return &value{val: os.Getenv(key)}
}

// Load loads the configuration from the .env file and the environment
func Load() *Config {
	godotenv.Load(configFile)

	c := &Config{}

	c.Strategy = valueOf("STRATEGY").optional("sequential").string()

	c.Interval = valueOf("SEND_INTERVAL").optional("5").int()

	c.Sender = valueOf("SENDER").optional("smtp").string()

	if c.Sender == "smtp" {
		c.Host = valueOf("SMTP_HOST").required().string()

		c.Port = valueOf("SMTP_PORT").required().int()

		c.Username = valueOf("SMTP_USERNAME").required().string()

		c.Password = valueOf("SMTP_PASSWORD").required().string()
	}

	return c
}
