package crawler

import (
	"github.com/ssalvatori/fondosMutuosGo/fondosmutuos"
)

const dataURL = "http://"

type Crawler struct {
	fmManager *fm.FmManager
}

func NewCrawler() *Crawler {

	c := new(Crawler)
	c.fmManager = fm.NewFmManager()

	return c
}
