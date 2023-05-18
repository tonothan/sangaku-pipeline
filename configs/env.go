package configs

import "os"

func EnvMongoURI() string {
	return os.Getenv("MONGO_URI")
}

func EnvMongoCollection() string {
	return os.Getenv("MONGO_COLLECTION")
}

func EnvImageStorePath() string {
	return os.Getenv("IMAGE_STORE_PATH")
}
