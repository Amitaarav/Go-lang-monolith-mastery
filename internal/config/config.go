package config // conventionally same, but not mandatory to be name as the folder
import (
	"os"

	"github.com/joho/godotenv"
)

// Config, C capital to make public
type Config struct {
	Port string
	Env  string
}

// Must pattern

func MustLoad() Config {
	// load is must
	// error is normal value, we can receive error in function and return the error
	// 
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == ""{
		panic("PORT is required")
	}

	env := os.Getenv("ENV")
	if env == "" {
		panic("ENV is required")
	}

	return Config{
		Port: port,
		Env: env,
	}
}