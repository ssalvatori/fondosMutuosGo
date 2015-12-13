package fm

import (
	"fmt"
)

type FmManager struct {
	fm map[string]*Fm
}

func NewFmManager() *FmManager {
	fmt.Println("creating fondosmutuosManager")

	fmm := new(FmManager)
	fmm.fm = make(map[string]*Fm)

	return fmm
}
