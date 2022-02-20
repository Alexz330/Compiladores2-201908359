package node

type Nodo struct {
	A string
	B string
}

func NewNodo(x string, y string) Nodo {
	return Nodo{A: x, B: y}
}

