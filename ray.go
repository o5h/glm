package glm

type Ray struct {
	Origin    Vec3
	Direction Vec3
}

func RayCast(touch Vec2, origin Vec3, proj, view Mat4x4) *Ray {

	rayClip := Vec4{X: touch.X, Y: touch.Y, Z: -1, W: 1}

	var inverseProj Mat4x4
	inverseProj.CopyInverseFrom(proj)
	rayEye := Vec4{
		X: inverseProj[0]*rayClip.X + inverseProj[4]*rayClip.Y + inverseProj[8]*rayClip.Z + inverseProj[12]*rayClip.W,
		Y: inverseProj[1]*rayClip.X + inverseProj[5]*rayClip.Y + inverseProj[9]*rayClip.Z + inverseProj[13]*rayClip.W,
		Z: -1,
		W: 0}

	var inverseView Mat4x4
	inverseView.CopyInverseFrom(view)
	rayWor := inverseView.MulVec4(rayEye)
	r := Ray{}
	r.Origin = origin
	r.Direction.SetXYZ(rayWor.X, rayWor.Y, rayWor.Z)
	r.Direction.Normalize()
	return &r
}
