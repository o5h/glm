package glm

type AABB struct {
	Min, Max Vec3
}

func (aabb *AABB) ContainsVec3(k Vec3) bool {
	return aabb.Min.X <= k.X &&
		aabb.Min.Y <= k.Y &&
		aabb.Min.Z <= k.Z &&
		aabb.Max.X >= k.X &&
		aabb.Max.Y >= k.Y &&
		aabb.Max.Z >= k.Z
}
