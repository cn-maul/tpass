package tpass

import (
	"errors"
	"math/rand"
)

const (
	Digitchar   = "0123456789"
	Upperchar   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Lowerchar   = "abcdefghijklmnopqrstuvwxyz"
	Specialchar = "!@#$%^*()"
)

// 生成密码的标志位，包装成结构体传入函数
type Flag struct {
	Length  int
	Digit   bool
	Upper   bool
	Lower   bool
	Special bool
}

func (flag *Flag) Generate() (pass string, err error) {
	if flag.Length == 0 || (!flag.Digit && !flag.Upper && !flag.Lower && !flag.Special) {
		err = errors.New("要求生成的长度为0")
		return "", err
	}

	var charset string
	switch {
	case flag.Digit:
		charset += Digitchar
		fallthrough
	case flag.Upper:
		charset += Upperchar
		fallthrough
	case flag.Lower:
		charset += Lowerchar
		fallthrough
	case flag.Special:
		charset += Specialchar
	}

	password := make([]byte, flag.Length)
	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}

	return string(password), nil
}
