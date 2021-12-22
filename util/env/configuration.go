package env

import "os"

var (
	PORT      string
	MONGO_URI string
)

func Load() {
	PORT = os.Getenv("PORT")
	MONGO_URI = os.Getenv("MONGO_URI")
}
