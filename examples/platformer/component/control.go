package component

type Control struct {
	JumpSpeed float64 // Initial jump speed.
	LowSpeed  float64 // Falling speed after reaching the apex.
	FallSpeed float64 // Falling speed before reaching the apex.
	MoveSpeed float64 // Horizontal travel speed.
}

func NewControl(jumpSpeed, lowSpeed, fallSpeed, moveSpeed float64) Control {
	return Control{
		JumpSpeed: jumpSpeed,
		LowSpeed:  lowSpeed,
		FallSpeed: fallSpeed,
		MoveSpeed: moveSpeed,
	}
}
