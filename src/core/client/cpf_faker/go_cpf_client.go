package cpffaker

import "github.com/mvrilo/go-cpf"

type GoCPF struct{}

func NewGoCPF() *GoCPF {
	return &GoCPF{}
}

func (g *GoCPF) Generate() string {
	return cpf.Generate()
}
