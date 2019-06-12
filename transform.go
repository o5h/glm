package glm

var TransformIdent = Transform{
	Position: Vec3{0, 0, 0},
	Scale:    Vec3{1, 1, 1},
	Rotation: Vec3{0, 0, 0}}

type Transform struct {
	Position Vec3 `json:"pos"`
	Scale    Vec3 `json:"scale"`
	Rotation Vec3 `json:"rot"`
}

func (t *Transform) SetSum(a, b *Transform) {
	t.Position.SetSum(&a.Position, &b.Position)
	t.Scale.SetMul(&a.Scale, &b.Scale)
	t.Rotation.SetMul(&a.Rotation, &b.Rotation)
}
