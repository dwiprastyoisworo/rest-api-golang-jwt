package config

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

var (
	postgresHost   string
	postgresPort   string
	postgresUser   string
	postgresPass   string
	postgresDbName string
	JwtSecret      string
	JwtExp         int
	JwtAlgorythm   string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	if os.Getenv("POSTGRES_HOST") != "" {
		postgresHost = os.Getenv("POSTGRES_HOST")
	} else {
		panic(errors.New("POSTGRES_HOST is empty"))
	}

	if os.Getenv("POSTGRES_PORT") != "" {
		postgresPort = os.Getenv("POSTGRES_PORT")
	} else {
		panic(errors.New("POSTGRES_PORT is empty"))
	}

	if os.Getenv("POSTGRES_USER") != "" {
		postgresUser = os.Getenv("POSTGRES_USER")
	} else {
		panic(errors.New("POSTGRES_USER is empty"))
	}

	if os.Getenv("POSTGRES_PASSWORD") != "" {
		postgresPass = os.Getenv("POSTGRES_PASSWORD")
	} else {
		panic(errors.New("POSTGRES_PASSWORD is empty"))
	}

	if os.Getenv("POSTGRES_DB_NAME") != "" {
		postgresDbName = os.Getenv("POSTGRES_DB_NAME")
	} else {
		panic(errors.New("POSTGRES_DB_NAME is empty"))
	}

	// ------------- JWT -------------

	if os.Getenv("JWT_SECRET") != "" {
		JwtSecret = os.Getenv("JWT_SECRET")
	} else {
		panic(errors.New("JWT_SECRET is empty"))
	}

	if os.Getenv("JWT_ALGORITHM") != "" {
		JwtAlgorythm = os.Getenv("JWT_ALGORITHM")
	} else {
		panic(errors.New("JWT_ALGORITHM is empty"))
	}

	if os.Getenv("JWT_EXPIRES_IN") != "" {
		JwtExp, err = strconv.Atoi(os.Getenv("JWT_EXPIRES_IN"))
		if err != nil {
			panic(err)
		}

	} else {
		panic(errors.New("JWT_EXPIRES_IN is empty"))
	}
}
