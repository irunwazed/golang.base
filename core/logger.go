package core

import "log"

func Logger(name string, data any) {
	log.Println(name)
	log.Println(data)
}
