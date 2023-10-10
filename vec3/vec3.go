package vec3

import "github.com/o5h/glm/math"

type Vec3 struct{ X, Y, Z float32 }

func New(x, y, z float32) Vec3 { return Vec3{x, y, z} }

func EqEpsilon(v, k Vec3) bool       { return Distance(v, k) < math.Epsilon }
func Distance(v, k Vec3) float32     { return math.LengthXYZ(v.X-k.X, v.Y-k.Y, v.Z-k.Z) }
func DotProduct(v1, v2 Vec3) float32 { return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z }
func CrossProduct(v, k Vec3) Vec3 {
	return Vec3{v.Y*k.Z - v.Z*k.Y, v.Z*k.X - v.X*k.Y, v.X*k.Y - v.Y*k.X}
}
