package util

import (
	"golang.org/x/exp/rand"
	"time"
)

func GetRandFloat64() float64 {
	var min, max float64
	min, max = 0, 50
	source := rand.NewSource(uint64(time.Now().UnixNano()))
	random := rand.New(source)
	return min + random.Float64()*(max-min)
}

func GetRandomString() string {
	strings := []string{"良好", "恶劣", "不错", "优质", "很差"}
	rand.Seed(uint64(time.Now().UnixNano()))
	index := rand.Intn(len(strings))
	return strings[index]
}
