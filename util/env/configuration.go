package env

import "os"

var (
	PORT      string
	MONGO_URI string
	DB_NAME   string
)

func Load() {
	PORT = os.Getenv("PORT")
	MONGO_URI = os.Getenv("MONGO_URI")
	DB_NAME = os.Getenv("DB_NAME")
}
