package idgenerate

import (
    "math/rand"
    "time"
)

func GenerateUniqueID() int {
    rand.Seed(time.Now().UnixNano())
    return rand.Int()
}