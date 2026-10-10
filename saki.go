package main

import (
	"fmt"
	"math"
	"os"
)

const rho = 7800.0

const c11_ferrite = 230e9
const c12_ferrite = 135e9
const c44_ferrite = 115e9
const tau_c_ferrite = 150e6

const c11_austenite = 198e9
const c12_austenite = 125e9
const c44_austenite = 122e9
const tau_c_austenite = 100e6

// math Block 3x3 or 3x1 format
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

func vec_minus3x1(vec1, vec2 [3]float64) [3]float64 {
	return [3]float64{vec1[0] - vec2[0], vec1[1] - vec2[1], vec1[2] - vec2[2]}
}

func vec_scale3x1(vec [3]float64, scal float64) [3]float64 {
	return [3]float64{vec[0] * scal, vec[1] * scal, vec[2] * scal}
}
func outer_mult_vec3x3(vec1, vec2 [3]float64) [3][3]float64 {
	var result [3][3]float64
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			result[i][j] = vec1[i] * vec2[j]
		}
	}
	return result
}
func sum_of_all_elements_tens3x3(tensor [3][3]float64) float64 {
	sum := 0.0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			sum += tensor[i][j]
		}
	}
	return sum
}
func sign_of_num(x float64) float64 {
	if x >= 0.0 {
		return 1.0
	} else {
		return -1.0
	}
}
func matr_sum3x3(matr1, matr2 [3][3]float64) [3][3]float64 {
	var result [3][3]float64
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			result[i][j] = matr1[i][j] + matr2[i][j]
		}
	}
	return result
}

func matr_and_number_mul3x3(matr [3][3]float64, number float64) [3][3]float64 {
	return [3][3]float64{
		{matr[0][0] * number, matr[0][1] * number, matr[0][2] * number},
		{matr[1][0] * number, matr[1][1] * number, matr[1][2] * number},
		{matr[2][0] * number, matr[2][1] * number, matr[2][2] * number},
	}
}
func double_contract3x3(a, b [3][3]float64) float64 {
	sum := 0.0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			sum += a[i][j] * b[i][j]
		}
	}
	return sum
}
func to_radians(degrees float64) float64 {
	return degrees * (math.Pi / 180.0)
}

func transponir_tens3x3(tens [3][3]float64) [3][3]float64 {
	var res [3][3]float64
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			res[i][j] = tens[j][i]
		}
	}
	return res
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

// end of math block
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

func compute_plastic_velocity_gradient(stress_tens [3][3]float64, burgers_directions, normal_direction [12][3]float64, critical_stress, gamma_0, exponent float64) [3][3]float64 {
	var plastic_gradient [3][3]float64
	for system := 0; system < 12; system++ {
		slip_tensor := outer_mult_vec3x3(burgers_directions[system], normal_direction[system])
		resolved_stress := double_contract3x3(stress_tens, slip_tensor)
		if math.Abs(resolved_stress) > critical_stress {
			plastic_rate := (gamma_0 * math.Pow((math.Abs(resolved_stress)/critical_stress), exponent) * sign_of_num(resolved_stress))
			plastic_gradient = matr_sum3x3(plastic_gradient, matr_and_number_mul3x3(slip_tensor, plastic_rate))
		}
	}
	return plastic_gradient
}

// ----------------------Bond_class-------------------------------
type Bond struct {
	particle1        Particle
	particle2        Particle
	Bond_vec         [3]float64
	Length0          float64
	Weight           float64
	Bond_status      int
	Bond_type        int
	Critical_stretch float64
}

func new_bond(particle1, particle2 Particle) Bond {
	bondvec := [3]float64{(particle2.X_curr[0] - particle1.X_curr[0]), (particle2.X_curr[1] - particle1.X_curr[1]), (particle2.X_curr[2] - particle1.X_curr[2])}
	length := linalg_norm3x1(bondvec)
	weight := math.Exp(-(length * length))
	bondtype := 0
	critical_stretch := 0.0
	if particle1.Phase_id == particle2.Phase_id {
		if particle1.Phase_id == 0 {
			bondtype = 0
			critical_stretch = 0.04
		} else {
			bondtype = 1
			critical_stretch = 0.06
		}
	} else {
		bondtype = 2
		critical_stretch = 0.005
	}
	return Bond{
		particle1:        particle1,
		particle2:        particle2,
		Bond_vec:         bondvec,
		Length0:          length,
		Weight:           weight,
		Bond_status:      1,
		Bond_type:        bondtype,
		Critical_stretch: critical_stretch,
	}
}

//----------------------Bond_class-------------------------------

// ----------------------Particle_class-------------------------------
type Particle struct {
	Index  int64
	X_ref  [3]float64
	X_curr [3]float64

	Phase_id int

	U_curr [3]float64
	U_prev [3]float64

	Force [3]float64

	F   [3][3]float64
	F_p [3][3]float64

	K_tensor [3][3]float64
	K_inv    [3][3]float64

	Stress [3][3]float64

	Angle float64

	R [3][3]float64

	Burgers_vec     [12][3]float64
	Normal_vec      [12][3]float64
	Elastic_tensor  [3][3][3][3]float64
	Critical_stress float64
	NeighborsIdx    []int64
}

// func new_particle(index, phase_id int, r, f, f_p [3][3]float64, ferrite_elastic_tensor, austenite_elastic_tensor [3][3][3][3]float64, critical_stress, angle, tau_c_ferrite, tau_c_austenite float64, bBCC, nBCC, bFCC, nFCC [12][3]float64, x_ref, x_curr [3]float64) Particle {
// 	angle_rad := to_radians(angle)

// 	cos_A := math.Cos(angle_rad)
// 	sin_A := math.Sin(angle_rad)

// 	r = [3][3]float64{
// 		{cos_A, -sin_A, 0.0},
// 		{sin_A, cos_A, 0.0},
// 		{0.0, 0.0, 1.0},
// 	}

// 	var I [3][3]float64
// 	for i := 0; i < 3; i++ {
// 		I[i][i] = 1.0
// 	}
// 	f = I
// 	f_p = I
// 	const (
// 		FERRITE   = 0
// 		AUSTENITE = 1
// 	)
// 	var (
// 		b_rot       [12][3]float64
// 		n_rot       [12][3]float64
// 		elastic_rot [3][3][3][3]float64
// 		tau_c       float64
// 	)
// 	if phase_id == FERRITE {
// 		for i := 0; i < 12; i++ {
// 			b_rot[i] = matr_vec_mul3x1(r, bBCC[i])
// 			n_rot[i] = matr_vec_mul3x1(r, nBCC[i])
// 		}
// 		elastic_rot = rotate_elas_tens(ferrite_elastic_tensor, r)

// 		tau_c = tau_c_ferrite

// 	} else {
// 		for i := 0; i < 12; i++ {
// 			b_rot[i] = matr_vec_mul3x1(r, bFCC[i])
// 			n_rot[i] = matr_vec_mul3x1(r, nFCC[i])
// 		}
// 		elastic_rot = rotate_elas_tens(austenite_elastic_tensor, r)

// 		tau_c = tau_c_austenite
// 	}
// 	return Particle{
// 		Index:           index,
// 		X_ref:           x_ref,
// 		X_curr:          x_ref,
// 		Phase_id:        phase_id,
// 		Angle:           angle_rad,
// 		R:               r,
// 		F:               f,
// 		F_p:             f_p,
// 		Burgers_vec:     b_rot,
// 		Normal_vec:      n_rot,
// 		Elastic_tensor:  elastic_rot,
// 		Critical_stress: tau_c,
// 		// Force, U_curr, U_prev, K_tensor, K_inv, Stress зануляются автоматически
// 	}
// }

func new_particle(index int64, x_ref [3]float64, phase_id int) Particle {
	var I [3][3]float64
	for i := 0; i < 3; i++ {
		I[i][i] = 1.0
	}
	return Particle{
		Index:    index,
		X_ref:    x_ref,
		X_curr:   x_ref,
		Phase_id: phase_id,
		F:        I,
		F_p:      I,
	}
}
func (p *Particle) InitPhysics(angle float64, ferrite_tens, austenite_tens [3][3][3][3]float64, bBCC, nBCC, bFCC, nFCC [12][3]float64) {
	p.Angle = to_radians(angle)
	cos_A := math.Cos(p.Angle)
	sin_A := math.Sin(p.Angle)

	p.R = [3][3]float64{
		{cos_A, -sin_A, 0.0},
		{sin_A, cos_A, 0.0},
		{0.0, 0.0, 1.0},
	}

	if p.Phase_id == 0 { // Феррит
		for i := 0; i < 12; i++ {
			p.Burgers_vec[i] = matr_vec_mul3x1(p.R, bBCC[i])
			p.Normal_vec[i] = matr_vec_mul3x1(p.R, nBCC[i])
		}
		p.Elastic_tensor = rotate_elas_tens(ferrite_tens, p.R)
		p.Critical_stress = tau_c_ferrite
	} else { // Аустенит
		for i := 0; i < 12; i++ {
			p.Burgers_vec[i] = matr_vec_mul3x1(p.R, bFCC[i])
			p.Normal_vec[i] = matr_vec_mul3x1(p.R, nFCC[i])
		}
		p.Elastic_tensor = rotate_elas_tens(austenite_tens, p.R)
		p.Critical_stress = tau_c_austenite
	}
}

func (p *Particle) FindNeighbors(particles []Particle, horizon float64) {
	for _, otherparticle := range particles {
		if p.Index == otherparticle.Index {
			continue
		}
		xi := vec_minus3x1(otherparticle.X_ref, p.X_ref)
		dist := linalg_norm3x1(xi)
		if dist <= horizon {
			p.NeighborsIdx = append(p.NeighborsIdx, otherparticle.Index)
		}
	}
}

//----------------------Particle_class-------------------------------

// ----------------------Bicrystall_class-------------------------------
type Bicrystal struct {
	N_ferrite      int64
	N_austenite    int64
	N_y            int64
	N_z            int64
	Spacing        float64
	Horizon_factor float64
	Volume         float64
	Horizon        float64
	Z_max          float64
	Particles      []Particle
}

func new_Bicrystal(n_ferrite, n_austenite, n_y, n_z int64, spacing, horizon, horizon_factor float64) Bicrystal {
	volume := spacing * spacing * spacing

	horizon = horizon_factor*spacing + 0.01
	z_max := (float64(n_z) - 1.0) * spacing

	return Bicrystal{
		N_ferrite:      n_ferrite,
		N_austenite:    n_austenite,
		N_y:            n_y,
		N_z:            n_z,
		Spacing:        spacing,
		Horizon_factor: horizon_factor,
		Volume:         volume,
		Horizon:        horizon,
		Z_max:          z_max,
		Particles:      []Particle{},
	}
}

func (bc *Bicrystal) GenerateParticles() {
	total_x := bc.N_austenite + bc.N_ferrite
	var particleIndex int64 = 0
	for i := int64(0); i < total_x; i++ {
		for j := int64(0); j < bc.N_y; j++ {
			for k := int64(0); k < bc.N_z; k++ {
				x := float64(i) * bc.Spacing
				y := float64(j) * bc.Spacing
				z := float64(k) * bc.Spacing
				x_ref := [3]float64{x, y, z}
				phase_id := 0
				if i >= bc.N_ferrite {
					phase_id = 1
				}

				bc.Particles = append(bc.Particles, new_particle(particleIndex, x_ref, phase_id))
				particleIndex++
			}
		}
	}
}
func (bc *Bicrystal) update_position(dt, damping_coeff, loading_vel float64) {
	for i := range bc.Particles {
		particle := &bc.Particles[i]
		if particle.X_ref[2] <= 0.0 {
			particle.U_curr = [3]float64{0.0, 0.0, 0.0}
		} else if particle.X_ref[2] >= bc.Z_max {
			particle.U_curr[2] += loading_vel * dt
		} else {
			next_displacement := vec_scale3x1(vec_add3x1(vec_add3x1(vec_minus3x1(vec_scale3x1(particle.U_curr, 2.0), particle.U_prev), vec_scale3x1(particle.Force, (dt*dt)/rho)), vec_scale3x1(particle.U_prev, 0.5*dt*damping_coeff)), 1/(1.0+0.5*damping_coeff*dt))
			particle.U_prev = particle.U_curr
			particle.U_curr = next_displacement
		}
		particle.X_curr = vec_add3x1(particle.X_ref, particle.U_curr)
	}
}
func (bc *Bicrystal) update_deformation_gradient() {
	for i := range bc.Particles {
		particle := &bc.Particles[i]
		var F [3][3]float64

		for _, neighborIdx := range particle.NeighborsIdx {
			neighbor := bc.Particles[neighborIdx]

			// xi = x_ref_neighbor - x_ref_particle
			xi := vec_minus3x1(neighbor.X_ref, particle.X_ref)

			// dy = x_curr_neighbor - x_curr_particle
			dy := vec_minus3x1(neighbor.X_curr, particle.X_curr)

			// Весовая функция weight = exp(-(||xi||^2))
			dist := linalg_norm3x1(xi)
			weight := math.Exp(-(dist * dist))

			// F += weight * (dy ⊗ xi) * Volume
			outer := outer_mult_vec3x3(dy, xi)
			scaledOuter := matr_and_number_mul3x3(outer, weight*bc.Volume)
			F = matr_sum3x3(F, scaledOuter)
		}

		// Итоговый градиент деформации: F = F * K_inv
		particle.F = matr_mul3x3(F, particle.K_inv)
	}
}

func (bc *Bicrystal) ExportToTXT(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Записываем заголовок (необязательно, но удобно для чтения)
	fmt.Fprintln(file, "Index, X, Y, Z, Phase_id")

	// Бежим по всем частицам и пишем их координаты и фазу
	for _, p := range bc.Particles {
		fmt.Fprintf(file, "%d, %.4f, %.4f, %.4f, %d\n",
			p.Index,
			p.X_ref[0], p.X_ref[1], p.X_ref[2],
			p.Phase_id,
		)
	}

	return nil
}

// ----------------------Bicrystall_class-------------------------------
func main() {

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

	// Создаем бикристалл:
	//n_ferrite = 10, n_austenite = 10, n_y = 5, n_z = 5, spacing = 1.0, horizon = ..., horizon_factor = 1.5
	bc := new_Bicrystal(10, 10, 5, 5, 1.0, 0.0, 1.5)

	// Генерируем сетку частиц
	bc.GenerateParticles()

	// Экспортируем в txt-файл
	err := bc.ExportToTXT("particles_grid.txt")
	if err != nil {
		fmt.Println("Ошибка при записи файла:", err)
		return
	}

	fmt.Printf("Успешно записано частиц: %d в файл particles_grid.txt\n", len(bc.Particles))
}
