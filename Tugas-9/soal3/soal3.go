package soal3

import "strconv"

func LuasPersegi(sisi int, kondisi bool) interface{} {
	var secret interface{}

	if sisi != 0 && kondisi {
		luas := sisi * sisi
		secret = "luas persegi dengan sisi " + strconv.Itoa(sisi) + " cm adalah " + strconv.Itoa(luas) + " cm"
	} else if sisi != 0 && !kondisi {
		secret = sisi
	} else if sisi == 0 && kondisi {
		secret = "Maaf anda belum menginput sisi dari persegi"
	} else {
		secret = nil
	}

	return secret
}
