package random

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var (
	pwdTable = "abcdefhjmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWYXZ"
	symTable = "!$%&'()*+,-./:;<=>?@[#]^_`{|}~"
	numTable = "23456789"
)

type Random struct {
	Length   int
	PwdTable string
}

func NewRandom(length int, num, symbol bool) *Random {
	p := &Random{
		Length:   length,
		PwdTable: pwdTable,
	}
	if num {
		p.PwdTable = p.PwdTable + numTable
	}
	if symbol {
		p.PwdTable = p.PwdTable + symTable
	}
	return p
}

func (p *Random) New() string {
	pwd := ""
	lenTable := len(p.PwdTable)
	for i := 0; i < p.Length; i++ {
		r := rand.Intn(lenTable)
		pwd += p.PwdTable[r : r+1]
	}
	return pwd
}
