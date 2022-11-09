package urun

type Urun struct {
	ID    uint    `json:"id" gorm:"primaryKey"`
	Adi   string  `json:"adi"`
	Fiyat float64 `json:"fiyat"`
	Renk  string  `json:"renk"`
}

func NewUrun(ad, renk string, fiyat float64) *Urun {
	return &Urun{
		Adi:   ad,
		Fiyat: fiyat,
		Renk:  renk,
	}
}
