package idgenerate

import (
	"math/rand"
	"time"
)

func GenerateUniqueID() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int()
}