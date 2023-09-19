package glm

var (
	ONE = Vec2{X: 1, Y: 1}
)

type Vec2 struct {
	X float32
	Y float32
}

func (v *Vec2) SetXY(x, y float32) {
	v.X, v.Y = x, y
}

func (v *Vec2) AddXY(x, y float32) {
	v.X += x
	v.Y += y
}

func (v *Vec2) Length() float32 {
	return Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vec2) DistanceTo(k *Vec2) float32 {
	return LengthXY(v.X-k.X, v.Y-k.Y)
}

func (v *Vec2) Normalize() {
	f := 1.0 / v.Length()
	v.X *= f
	v.Y *= f
}

func Vec2DotProduct(v1, v2 Vec2) float32 {
	return v1.X*v2.X + v1.Y*v2.Y
}

func (v1 *Vec2) EqEpsilon(v2 *Vec2) bool {
	return v1.DistanceTo(v2) < Epsilon
}
