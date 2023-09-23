package glm

var Transform_Ident = Transform{
	Translation: Vec3{X: 0, Y: 0, Z: 0},
	Rotation:    Vec3{X: 0, Y: 0, Z: 0},
	Scale:       Vec3{X: 1, Y: 1, Z: 1}}

type Transform struct {
	Translation Vec3 `yaml:"position,flow"`
	Rotation    Vec3 `yaml:"rotation,flow"`
	Scale       Vec3 `yaml:"scale,flow"`
}

func (t *Transform) SetSum(a, b *Transform) {
	t.Translation.SetSum(&a.Translation, &b.Translation)
	t.Rotation.SetSum(&a.Rotation, &b.Rotation)
	t.Scale.SetMul(&a.Scale, &b.Scale)
}
