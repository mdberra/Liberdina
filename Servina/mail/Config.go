package acceso
 
import (
	"log"
 
	"github.com//toml"
)
 
type Config struct {
	Server   string
	Port     int
	Email    string
	Password string
}
 
func (c *Config) Read() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}e
}