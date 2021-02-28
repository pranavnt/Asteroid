package main

import (
	"os"
)

func createCollection(name string) {
	os.Create("db/collections/" + name + ".properties")
}
