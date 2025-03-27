package main

import "fmt"

const ebulicaoK = 373.2

func main() {
	ebulicaoC := (ebulicaoK - 273)

	fmt.Printf("O ponto de ebulição da água em graus celsius é de: %g", ebulicaoC)
}
