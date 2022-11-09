package urun_test

import (
	"errors"
	"github.com/go-turk/konya/urun"
	"gorm.io/gorm"
	"testing"
)

func TestNewUrunRepository(t *testing.T) {
	type input struct {
		db *gorm.DB
	}
	type expected struct {
		urunRepository urun.UrunRepository
		err            error
	}
	testcases := []struct {
		name     string
		input    input
		expected expected
	}{
		{
			name: "db is nil",
			input: input{
				db: nil,
			},
			expected: expected{
				urunRepository: nil,
				err:            errors.New("db is nil"),
			},
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			result, err := urun.NewUrunRepository(tt.input.db)
			if err != nil {
				if err.Error() != tt.expected.err.Error() {
					t.Errorf("expected: %v, got: %v", tt.expected.err, err)
				}
			}
			if result != tt.expected.urunRepository {
				t.Errorf("expected: %v, got: %v", tt.expected.urunRepository, result)
			}
		})
	}
}

func TestUrunRepository_Create(t *testing.T) {

}

func TestUrunRepository_Update(t *testing.T) {

}

func TestUrunRepository_Delete(t *testing.T) {

}

func TestUrunRepository_FindAll(t *testing.T) {

}

func TestUrunRepository_FindByID(t *testing.T) {

}

func TestUrunRepository_DeleteAll(t *testing.T) {

}
