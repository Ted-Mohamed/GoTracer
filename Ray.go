package main

//Ray
type Ray struct {
	Origin, Direction Vector3D
}

//Point
func (r Ray) Point(t float64) Vector3D {
	return r.Origin.Add(r.Direction.MultiplyScalar(t))
}
