package glm

type Transform3d struct {
	Position Vec3 `yaml:"position,flow"`
	Rotation Vec3 `yaml:"rotation,flow"`
	Scale    Vec3 `yaml:"scale,flow"`
}

func (t *Transform3d) GetPosition() Vec3         { return t.Position }
func (t *Transform3d) GetRotation() Vec3         { return t.Rotation }
func (t *Transform3d) GetScale() Vec3            { return t.Scale }
func (t *Transform3d) SetPosition(position Vec3) { t.Position = position }
func (t *Transform3d) SetRotation(rotation Vec3) { t.Rotation = rotation }
func (t *Transform3d) SetScale(scale Vec3)       { t.Scale = scale }

func (t *Transform3d) SetIdent() {
	t.Position = Vec3{X: 0, Y: 0, Z: 0}
	t.Rotation = Vec3{X: 0, Y: 0, Z: 0}
	t.Scale = Vec3{X: 1, Y: 1, Z: 1}
}

func (t *Transform3d) Set(a *Transform3d) {
	t.Position = a.Position
	t.Rotation = a.Rotation
	t.Scale = a.Scale
}

func (t *Transform3d) SetSum(a, b *Transform3d) {
	t.Position.SetSum(&a.Position, &b.Position)
	t.Rotation.SetSum(&a.Rotation, &b.Rotation)
	t.Scale.SetMul(&a.Scale, &b.Scale)
}
