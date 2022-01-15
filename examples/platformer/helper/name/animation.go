package name

func PlayerMovement(v bool) string {
	if v {
		return "run"
	}
	return "jump"
}

func PlayerDirection(v float64) string {
	if v < 0 {
		return "left"
	}
	return "right"
}
