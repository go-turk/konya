package matematikLib

type Matematik interface {
	Topla(a, b int) int
	Cikar(a, b int) int
	Carp(a, b int) int
	Bol(a, b int) int
	UssuAl(a, b int) float64
}

type matematik struct {
	sonuclar []int
}

func NewMatematik() Matematik {
	return &matematik{
		sonuclar: []int{},
	}
}

// a ile b sayısı ise gelirse sonucu sıfır gösterecek şekilde güncelleme yapınız
func (m *matematik) Topla(a, b int) int {
	if a == 1 && b == 1 {
		return 0
	}
	sonuc := a + b
	m.sonuclar = append(m.sonuclar, sonuc)
	return sonuc
}

func (m *matematik) Cikar(a, b int) int {
	sonuc := a - b
	m.sonuclar = append(m.sonuclar, sonuc)
	return sonuc
}

func (m *matematik) Carp(a, b int) int {
	sonuc := a * b
	m.sonuclar = append(m.sonuclar, sonuc)
	return sonuc
}

func (m *matematik) Bol(a, b int) int {
	sonuc := a / b
	m.sonuclar = append(m.sonuclar, sonuc)
	return sonuc
}

// a^b => a = 2 ve b = 3 ise sonuc 8 olmalı
// Eksili işlemler şu anda 0 dönecek şekilde tasarlanabilir.
// a = 5 ve b = -1 => sonuc 0.2 olmalı
func (m *matematik) UssuAl(a, b int) float64 {
	return -1.0
}
