package student

var div int
var mod int
var da int
var db int

func UltimateDivMod(a *int, b *int) {
	da = *a
	db = *b
	div = da / db
	mod = da % db
	*a = div
	*b = mod
}
