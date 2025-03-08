package main

import "fmt"

func luasLingkaran1(r float64, luasLingkaran *float64, kelilingLingkaran *float64) {
	*luasLingkaran = 3.14 * r * r
	*kelilingLingkaran = 2 * 3.14 * r
}

func luasPersegi1(sisi float64, luasPersegi *float64, kelilingPersegi *float64) {
	*luasPersegi = sisi * sisi
	*kelilingPersegi = 4 * sisi
}

func hitungTotal(kelilingLingkaran, kelilingPersegi, luasLingkaran, luasPersegi float64, totalKeliling *float64, totalLuas *float64) {
	*totalKeliling = kelilingLingkaran + kelilingPersegi
	*totalLuas = luasLingkaran + luasPersegi
}

func main() {
	var jari, s float64
	var luasLing, kelilingLing, luasPers, kelilingPers, totalKel, totalLuas float64
	fmt.Scan(&jari, &s)

	for jari != 0 && s != 0 {
		luasLingkaran1(jari, &luasLing, &kelilingLing)
		luasPersegi1(s, &luasPers, &kelilingPers)
		hitungTotal(kelilingLing, kelilingPers, luasLing, luasPers, &totalKel, &totalLuas)

		fmt.Printf("%.2f %.2f %.2f %.2f %.2f %.2f %.2f %.2f\n", jari, s, luasLing, kelilingLing, luasPers, kelilingPers, totalKel, totalLuas)
		fmt.Scan(&jari, &s)
	}

}
