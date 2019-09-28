package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"math/rand"
	"strconv"
	"time"
)

func RandSeed() string {
	rand.Seed(time.Now().Unix())
	return Md5(strconv.Itoa(rand.Intn(1000000) + rand.Intn(100000)))
}

func Md5(text string) string {
	h := md5.New()
	io.WriteString(h, text)
	return (fmt.Sprintf("%x", h.Sum(nil)))
}
