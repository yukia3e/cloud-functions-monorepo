package config

import "os"

func GetEnvironment() string {
	return os.Getenv("APP_ENV")
}

func IsLocal() bool {
	return GetEnvironment() == "local"
}

func GetPort() string {
	return os.Getenv("PORT")
}

func MustGetFunctionEntryPointName() string {
	name := os.Getenv("FUNCTION_ENTRY_POINT_NAME")
	if name == "" {
		panic("FUNCTION_ENTRY_POINT_NAME is not set")
	}

	return name
}
