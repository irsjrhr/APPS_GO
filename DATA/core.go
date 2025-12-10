package DATA


import (
    "time"
)

type InputMahasiswa struct{

    Nama string
    TanggalLahir string
    Gender int
    IDHobi int
    IDJurusan int

}
type InputJurusan struct{ Nama string }
type InputHobi struct{ Nama string }

type DataMahasiswa struct {
    Id  int
    Nama string
    TanggalLahir time.Time
    Gender int
    Jurusan string
    Hobi string
}

type DataJurusan struct {
    Id  int
    Nama string
}
type DataHobi struct {
    Id  int
    Nama string
}