package ray

import (
	"github.com/o5h/glm/f32/mat4x4"
	"github.com/o5h/glm/f32/vec2"
	"github.com/o5h/glm/f32/vec3"
	"github.com/o5h/glm/f32/vec4"
)

type Ray struct {
	Origin    vec3.Vec3
	Direction vec3.Vec3
}

func Cast(touch vec2.Vec2, origin vec3.Vec3, proj, view mat4x4.Mat4x4) *Ray {
	rayClip := vec4.Vec4{X: touch.X, Y: touch.Y, Z: -1, W: 1}
	var inverseProj mat4x4.Mat4x4
	inverseProj.CopyInverseFrom(proj)
	rayEye := vec4.Vec4{
		X: inverseProj[0]*rayClip.X + inverseProj[4]*rayClip.Y + inverseProj[8]*rayClip.Z + inverseProj[12]*rayClip.W,
		Y: inverseProj[1]*rayClip.X + inverseProj[5]*rayClip.Y + inverseProj[9]*rayClip.Z + inverseProj[13]*rayClip.W,
		Z: -1,
		W: 0}

	var inverseView mat4x4.Mat4x4
	inverseView.CopyInverseFrom(view)
	rayWor := inverseView.MulVec4(rayEye)

	direction := vec3.Vec3{X: rayWor.X, Y: rayWor.Y, Z: rayWor.Z}
	direction.Normalize()
	return &Ray{
		Origin:    origin,
		Direction: direction}
}
