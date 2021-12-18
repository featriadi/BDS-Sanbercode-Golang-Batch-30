package models

type NilaiMahasiswa struct {
	Nama        string `json:"nama"`
	MataKuliah  string `json:"mata_kuliah"`
	IndeksNilai string `json:"index_nilai"`
	Nilai       uint   `json:"nilai"`
	ID          uint   `json:"id"`
}
