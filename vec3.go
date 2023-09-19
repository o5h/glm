package glm

type Vec3 struct {
	X float32
	Y float32
	Z float32
}

func (v *Vec3) SetXYZ(x, y, z float32) {
	v.X = x
	v.Y = y
	v.Z = z
}

func (v *Vec3) SetSum(a, b *Vec3) {
	v.X = a.X + b.X
	v.Y = a.Y + b.Y
	v.Z = a.Z + b.Z
}

func (v *Vec3) SetSub(a, b *Vec3) {
	v.X = a.X - b.X
	v.Y = a.Y - b.Y
	v.Z = a.Z - b.Z
}

func (v *Vec3) Add(k Vec3) {
	v.X += k.X
	v.Y += k.Y
	v.Z += k.Z
}

func (v *Vec3) Sub(k Vec3) {
	v.X -= k.X
	v.Y -= k.Y
	v.Z -= k.Z
}

func (v *Vec3) SetMul(a, b *Vec3) {
	v.X = a.X * b.X
	v.Y = a.Y * b.Y
	v.Z = a.Z * b.Z
}

func (v *Vec3) Scale(factor float32) {
	v.X *= factor
	v.Y *= factor
	v.Z *= factor
}

func (v *Vec3) DistanceTo(k *Vec3) float32 {
	return LengthXYZ(v.X-k.X, v.Y-k.Y, v.Z-k.Z)
}

func (v *Vec3) DistanceToSquared(k *Vec3) float32 {
	return LengthSquared(v.X-k.X, v.Y-k.Y, v.Z-k.Z)
}

func (v *Vec3) Length() float32 {
	return Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v *Vec3) Normalize() {
	f := 1.0 / v.Length()
	v.X *= f
	v.Y *= f
	v.Z *= f
}

func (v1 *Vec3) EqEpsilon(v2 *Vec3) bool {
	return v1.DistanceTo(v2) < Epsilon
}

func Vec3SumOf(v1, v2 Vec3) Vec3 {
	return Vec3{v1.X + v2.X, v1.Y + v2.Y, v1.Z + v2.Z}
}
