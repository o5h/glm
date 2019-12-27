package glm

var TransformIdent = Transform{
	Location: Vec3{0, 0, 0},
	Scale:    Vec3{1, 1, 1},
	Rotation: Vec3{0, 0, 0}}

type Transform struct {
	Location Vec3
	Scale    Vec3
	Rotation Vec3
}

func (t *Transform) Set(a *Transform) {
	t.Location.Set(&a.Location)
	t.Scale.Set(&a.Scale)
	t.Rotation.Set(&a.Rotation)
}

func (t *Transform) SetSum(a, b *Transform) {
	t.Location.SetSum(&a.Location, &b.Location)
	t.Scale.SetMul(&a.Scale, &b.Scale)
	t.Rotation.SetMul(&a.Rotation, &b.Rotation)
}
