package lingkaran_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"Lingkaran/Lingkaran"
)
var _ = Describe("Lingkaran", func() {
	Describe("HitungLuas", func() {
		It("Menghitung luas lingkaran dengan jari-jari 4", func() {
			luas := lingkaran.HitungLuas(4)
			Expect(luas).To(Equal(50.26548245743669))
		})

		It("Menghitung luas lingkaran dengan jari-jari 10", func() {
			luas := lingkaran.HitungLuas(10)
			Expect(luas).To(Equal(314.1592653589793))
		})
	})

	Describe("HitungKeliling", func() {
		It("Menghitung keliling lingkaran dengan jari-jari 5", func() {
			keliling := lingkaran.HitungKeliling(5)
			Expect(keliling).To(Equal(31.41592653589793))
		})

		It("Menghitung keliling lingkaran dengan jari-jari 8", func() {
			keliling := lingkaran.HitungKeliling(8)
			Expect(keliling).To(Equal(50.26548245743669))
		})
	})

})