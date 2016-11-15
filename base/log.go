package base

import (
	"log"
)

func LogBytesIn(in []byte) {
	log.Printf("<IN>    ", in)
}

func LogBytesOut(out []byte) {
	log.Printf("<OUT>  ", out)
}
