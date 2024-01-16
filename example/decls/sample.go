package main

type Base struct {
	Name string
}

func (b *Base) PointerReceiver() {
}

func (b Base) ConcreteReceiver() {
}

func existFunc(a, b int) int {
	return a + b
}
