package student

func Swap(a *int, b *int) {
	tmpa := *a
	tmpb := *b
	*a = tmpb
	*b = tmpa
}
