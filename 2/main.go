package main

//Declaração de escopo global
var (
	b bool    = true
	c int     = 10
	d string  = "minha string"
	e float64 = 1.2
)

func main() {
	a := "já cria a variável e atribui"
	println(a)

	a = "modifica o valor de a"
	println(a)

	println(b)
	println(c)
	println(d)
	println(e)
}
