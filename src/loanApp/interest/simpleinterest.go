package interest

func Si(p float32, r float32, t int) float32 {
	i := (p * r * float32(t)) / 100
	return i
}
