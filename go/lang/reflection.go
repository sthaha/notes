package main

import (
	"fmt"
	"os"
	"strconv"
)

type StringType func() string
type IntType func() int
type StringListType func() []string

func lookupEnv(env, defaultVal string) string {
	if value, ok := os.LookupEnv(env); ok {
		return value
	}
	return defaultVal
}

func String(key, fallback string) StringType {
	return func() string {
		if value, ok := os.LookupEnv(key); ok {
			return value
		}
		return fallback
	}
}

func Int(key string, fallback int) IntType {
	return func() int {

		value, ok := os.LookupEnv(key)
		if !ok {
			return fallback
		}

		if i, err := strconv.Atoi(value); err == nil {
			return i
		}
		return fallback
	}
}

type config struct {
	bar StringType
	age IntType
}

func newConfig() *config {
	return &config{
		bar: String("BAR", "foobar"),
		age: Int("AGE", 10),
	}
}

func main() {
	c := newConfig()
	fmt.Println(c.bar())
	fmt.Println(c.age())
}
