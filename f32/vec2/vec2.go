package vec2

import "github.com/o5h/glm/f32/math"

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
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vec2) DistanceTo(k *Vec2) float32 {
	return math.LengthXY(v.X-k.X, v.Y-k.Y)
}

func (v *Vec2) Normalize() {
	f := 1.0 / v.Length()
	v.X *= f
	v.Y *= f
}

func DotProduct(v1, v2 Vec2) float32 {
	return v1.X*v2.X + v1.Y*v2.Y
}

func Vec2Equals(v1, v2 *Vec2) bool {
	if v1 == nil && v2 == nil {
		return true
	}
	if v1 == nil || v2 == nil {
		return false
	}
	if v1.DistanceTo(v2) < math.Epsilon {
		return true
	}
	return false
}
