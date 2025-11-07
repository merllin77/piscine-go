package student

var (
	div int
	mod int
	da  int
	db  int
)

func UltimateDivMod(a *int, b *int) {
	da = *a
	db = *b
	*a = da / db
	*b = da % db
}
