package main

import "math"

//import "fmt"
//math Block 3x3 or 3x1 format
func normalize3x1(v [3]float64) [3]float64 {
	length := math.Sqrt(v[0]*v[0] + v[1]*v[1] + v[2]*v[2])
	return [3]float64{v[0] / length, v[1] / length, v[2] / length}
}

func linalg_norm3x1(v [3]float64) float64 {
	a := math.Sqrt(v[0]*v[0] + v[1]*v[1] + v[2]*v[2])
	return a
}

func matr_mul3x3(a, b [3][3]float64) [3][3]float64 {
	var c [3][3]float64
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				c[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return c
}

func matr_vec_mul3x1(a [3][3]float64, vec [3]float64) [3]float64 {
	var result [3]float64
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			result[i] += a[i][j] * vec[j]
		}
	}
	return result
}

func vec_add3x1(vec1, vec2 [3]float64) [3]float64 {
	return [3]float64{vec1[0] + vec2[0], vec1[1] + vec2[1], vec1[2] + vec2[2]}
}

func vec_scale3x1(vec [3]float64, scal float64) [3]float64 {
	return [3]float64{vec[0] * scal, vec[1] * scal, vec[2] * scal}
}

//math Block generic format

func normalize(v []float64) []float64 {
	rows := len(v)
	sum := 0.0
	for i := 0; i < rows; i++ {
		sum += v[i] * v[i]
	}
	length := math.Sqrt(sum)
	if length == 0 {
		return v
	}
	for i := 0; i < rows; i++ {
		v[i] /= length
	}
	return v
}

func linalg_norm(v []float64) float64 {
	rows := len(v)
	sum := 0.0
	for i := 0; i < rows; i++ {
		sum += v[i] * v[i]
	}
	length := math.Sqrt(sum)
	return length
}

func matr_mul(a, b [][]float64) [][]float64 {
	rowsA := len(a)
	colsA := len(a[0])
	rowsB := len(b)
	colsB := len(b[0])

	if colsA != rowsB {
		panic("colsA != rowsB")
	}

	result := make([][]float64, rowsA)
	for i := range result {
		result[i] = make([]float64, colsB)
	}

	for i := 0; i < rowsA; i++ {
		for j := 0; j < colsB; j++ {
			for k := 0; k < colsA; k++ {
				result[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return result
}

func matr_vec_mul(a [][]float64, vec []float64) []float64 {
	rowsA := len(a)
	colsA := len(a[0])

	if colsA != len(vec) {
		panic("colsA != length of vec")
	}
	result := make([]float64, rowsA)
	for i := 0; i < rowsA; i++ {
		for j := 0; j < len(vec); j++ {
			result[i] += a[i][j] * vec[j]
		}
	}
	return result
}

func build_elastic_tensor(c11, c12, c44 float64) [3][3][3][3]float64 {
	var tensor [3][3][3][3]float64
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == j {
				tensor[i][i][i][i] = c11

				for k := 0; k < 3; k++ {
					if k != i {
						tensor[i][i][k][k] = c12
					}
				}
			} else {
				tensor[i][j][i][j] = c44
			}
		}
	}
	return tensor
}

func rotate_elas_tens(elas_tens [3][3][3][3]float64, rotation_tens [3][3]float64) [3][3][3][3]float64 {
	var rotated_elas_tens [3][3][3][3]float64
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {
					for p := 0; p < 3; p++ {
						for q := 0; q < 3; q++ {
							for r := 0; r < 3; r++ {
								for s := 0; s < 3; s++ {
									rotated_elas_tens[i][j][k][l] += (rotation_tens[p][i] * rotation_tens[q][j] * rotation_tens[r][k] * rotation_tens[s][l] * elas_tens[p][q][r][s])
								}
							}
						}
					}
				}
			}
		}
	}
	return rotated_elas_tens
}

func compute_interparticle_force(stress_main_particle_tens, stress_neighbor_particle_tens, main_shape_tens, neighbor_shape_tens [3][3]float64, xi [3]float64, xi_norm float64, is_broken bool) [3]float64 {
	if is_broken == true {
		return [3]float64{0.0, 0.0, 0.0}
	}
	weight := math.Exp(-(xi_norm * xi_norm))
	force_main := vec_scale3x1(matr_vec_mul3x1(matr_mul3x3(stress_main_particle_tens, main_shape_tens), xi), weight)
	force_neighbor := vec_scale3x1(matr_vec_mul3x1(matr_mul3x3(stress_neighbor_particle_tens, neighbor_shape_tens), xi), weight)
	return vec_scale3x1(vec_add3x1(force_main, force_neighbor), 0.5)
}

func compute_plastic_velocity_gradient()
func main() {
	const rho = 7800.0

	const c11_ferrite = 230e9
	const c12_ferrite = 135e9
	const c44_ferrite = 115e9
	const tau_c_ferrite = 150e6
	//Направления скольжения феррита
	bBCC := [12][3]float64{
		{1.0, 1.0, 1.0}, {1.0, 1.0, 1.0},
		{1.0, 1.0, -1.0}, {1.0, 1.0, -1.0},
		{1.0, -1.0, 1.0}, {1.0, -1.0, 1.0},
		{-1.0, 1.0, 1.0}, {-1.0, 1.0, 1.0},
		{1.0, -1.0, -1.0}, {1.0, -1.0, -1.0},
		{-1.0, 1.0, -1.0}, {-1.0, 1.0, -1.0},
	}

	nBCC := [12][3]float64{
		{1.0, -1.0, 0.0}, {1.0, 0.0, -1.0},
		{1.0, -1.0, 0.0}, {0.0, 1.0, 1.0},
		{1.0, 1.0, 0.0}, {0.0, 1.0, 1.0},
		{1.0, 1.0, 0.0}, {1.0, 0.0, 1.0},
		{1.0, 1.0, 0.0}, {1.0, 0.0, -1.0},
		{1.0, 1.0, 0.0}, {0.0, 1.0, 1.0},
	}

	const c11_austenite = 198e9
	const c12_austenite = 125e9
	const c44_austenite = 122e9
	const tau_c_austenite = 100e6
	//Направления скольжения аустенита
	bFCC := [12][3]float64{
		{-1.0, 1.0, 0.0}, {1.0, 0.0, -1.0},
		{0.0, -1.0, 1.0}, {1.0, 0.0, 1.0},
		{-1.0, -1.0, 0.0}, {0.0, 1.0, -1.0},
		{-1.0, 0.0, 1.0}, {0.0, -1.0, -1.0},
		{1.0, 1.0, 0.0}, {1.0, -1.0, 0.0},
		{-1.0, 0.0, -1.0}, {0.0, 1.0, 1.0},
	}
	nFCC := [12][3]float64{
		{1.0, 1.0, 1.0}, {1.0, 1.0, 1.0},
		{1.0, 1.0, 1.0}, {-1.0, 1.0, 1.0},
		{-1.0, 1.0, 1.0}, {-1.0, 1.0, 1.0},
		{1.0, -1.0, 1.0}, {1.0, -1.0, 1.0},
		{1.0, -1.0, 1.0}, {-1.0, -1.0, 1.0},
		{-1.0, -1.0, 1.0}, {-1.0, -1.0, 1.0},
	}

	const gamma_0 = 2e-5
	const n_exp = 20.0
	for i := 0; i < 12; i++ {
		bBCC[i] = normalize3x1(bBCC[i])
		nBCC[i] = normalize3x1(nBCC[i])
		bFCC[i] = normalize3x1(bFCC[i])
		nFCC[i] = normalize3x1(nFCC[i])
	}
	//var ferrite_elas_tens [3][3][3][3]float64

	ferrite_elas_tens := build_elastic_tensor(c11_ferrite, c12_ferrite, c44_ferrite)
	austenite_elas_tens := build_elastic_tensor(c11_austenite, c12_austenite, c44_austenite)
	_ = austenite_elas_tens
	_ = ferrite_elas_tens
}
