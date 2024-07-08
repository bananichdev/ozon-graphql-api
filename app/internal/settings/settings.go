package settings

import "os"

var (
	Port       = os.Getenv("PORT")
	DBHost     = os.Getenv("DB_HOST")
	DBPort     = os.Getenv("DB_PORT")
	DBUser     = os.Getenv("DB_USER")
	DBPass     = os.Getenv("DB_PASS")
	DBName     = os.Getenv("DB_NAME")
	TestDBName = os.Getenv("TEST_DB_NAME")
	MemoryMode = false
)
