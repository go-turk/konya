package matematikLib

import "testing"

func TestMatematik_Topla(t *testing.T) {
	type input struct {
		a int
		b int
	}
	type excepted struct {
		sonuc int
	}
	testcases := []struct {
		name     string
		input    input
		excepted excepted
	}{
		{
			name: "-2-2=-4",
			input: input{
				a: -2,
				b: -2,
			},
			excepted: excepted{
				sonuc: -4,
			},
		},
		{
			name: "0+0=0",
			input: input{
				a: 0,
				b: 0,
			},
			excepted: excepted{
				sonuc: 0,
			},
		},
		{
			name: "0+(-1)=-1",
			input: input{
				a: 0,
				b: -1,
			},
			excepted: excepted{
				sonuc: -1,
			},
		},
		{
			name: "1+1=0",
			input: input{
				a: 1,
				b: 1,
			},
			excepted: excepted{
				sonuc: 0,
			},
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			mat := NewMatematik()
			result := mat.Topla(tt.input.a, tt.input.b)
			if result != tt.excepted.sonuc {
				t.Errorf("%s için beklenen sonuç %d ancak %d geldi", tt.name, tt.excepted.sonuc, result)
			}
		})
	}

}

func TestMatematik_Cikar(t *testing.T) {

}

func TestMatematik_Carp(t *testing.T) {

}

func TestMatematik_Bol(t *testing.T) {

}
