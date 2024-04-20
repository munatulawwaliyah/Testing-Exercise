package lingkaran

import "math"

func HitungLuas(jariJari float64) float64 {
    return math.Pi * jariJari * jariJari
}

func HitungKeliling(jariJari float64) float64 {
    return 2 * math.Pi * jariJari
}
