package main

import (
	"fmt"
	"github.com/go-turk/konya/database"
	"github.com/go-turk/konya/urun"
)

func main() {
	db, err := database.Connection()
	if err != nil {
		fmt.Println("Database Connection Hatası:", err)
		return
	}
	urunRepo, err := urun.NewUrunRepository(db)
	if err != nil {
		fmt.Println("Urun Repo Hatası:", err)
		return
	}
	backendLv2Rozeti := urun.NewUrun("Backend Lv2 Rozeti", "Siyah", 25.15)
	if err := urunRepo.Create(backendLv2Rozeti); err != nil {
		fmt.Println("Backend Lv2 Rozeti Eklenemedi", err)
		return
	}
	fmt.Println(backendLv2Rozeti)
}
