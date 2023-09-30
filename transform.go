package glm

var Transform_Ident = Transform{
	Rotation:    Quat{X: 0, Y: 0, Z: 0, W: 1},
	Translation: Vec3{X: 0, Y: 0, Z: 0},
	Scale:       Vec3{X: 1, Y: 1, Z: 1}}

type Transform struct {
	Rotation    Quat
	Translation Vec3
	Scale       Vec3
}
