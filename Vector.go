package main

import "math"

// type Vector2D [2]float64

//Vector3D
type Vector3D [3]float64

// type Vector4D [4]float64

//Components
func (v Vector3D) Components() (x, y, z float64) {
	return v[0], v[1], v[2]
}

//Cross
func (v Vector3D) Cross(i Vector3D) Vector3D {
	return Vector3D{v[1]*i[2] - v[2]*i[1], v[2]*i[0] - v[0]*i[2], v[0]*i[1] - v[1]*i[0]}
}

//Length
func (v Vector3D) Length() float64 {
	return math.Sqrt(v[0]*v[0] + v[1]*v[1] + v[2]*v[2])
}

//Dot
func (v Vector3D) Dot(i Vector3D) float64 {
	return v[0]*i[0] + v[1]*i[1] + v[2]*i[2]
}

//Normalize
func (v Vector3D) Normalize() Vector3D {
	length := v.Length()
	return Vector3D{v[0] / length, v[1] / length, v[2] / length}
}

//Add
func (v Vector3D) Add(i Vector3D) Vector3D {
	return Vector3D{v[0] + i[0], v[1] + i[1], v[2] + i[2]}
}

//Subtract
func (v Vector3D) Subtract(i Vector3D) Vector3D {
	return Vector3D{v[0] - i[0], v[1] - i[1], v[2] - i[2]}
}

//AddScalar
func (v Vector3D) AddScalar(s float64) Vector3D {
	return Vector3D{v[0] + s, v[1] + s, v[2] + s}
}

//SubtractScalar
func (v Vector3D) SubtractScalar(s float64) Vector3D {
	return Vector3D{v[0] - s, v[1] - s, v[2] - s}
}

//MultiplyScalar
func (v Vector3D) MultiplyScalar(s float64) Vector3D {
	return Vector3D{v[0] * s, v[1] * s, v[2] * s}
}

//DivideScalar
func (v Vector3D) DivideScalar(s float64) Vector3D {
	return Vector3D{v[0] / s, v[1] / s, v[2] / s}
}

//X
func (v Vector3D) X() float64 {
	return v[0]
}

//Y
func (v Vector3D) Y() float64 {
	return v[1]
}

//Z
func (v Vector3D) Z() float64 {
	return v[2]
}
