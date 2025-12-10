package DB

import (
	"database/sql"
	"APPS_TEST/DATA"
	_ "github.com/lib/pq"
	"log"
	"fmt"
	"os" // Digunakan untuk os.Getenv
	"github.com/joho/godotenv" // Library untuk membaca file .env
)

var Conn *sql.DB

type Response struct{

	Status bool
	Msg string

}
func cetak( nilai interface{} ) {
	fmt.Println( nilai )
}




// Fungsi cetak dan tipe Response tetap sama

func Init() {

	cetak("DB INIT")

	// Memuat variabel dari file .env.
	// Jika gagal (misal file tidak ada), aplikasi akan melanjutkan dan
	// mencoba membaca variabel dari lingkungan OS.
	err := godotenv.Load()
	if err != nil {
		log.Println("INFO: File .env tidak ditemukan. Menggunakan variabel lingkungan OS.")
	}

	// Mengambil nilai koneksi dari Environment Variables (ENV)
	// Jika variabel tidak ada di .env maupun ENV OS, gunakan nilai default.
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" { dbHost = "localhost" }

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" { dbPort = "5432" }

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" { dbUser = "postgres" }

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" { dbPassword = "admin" } // Sesuaikan default Anda

	dbName := os.Getenv("DB_NAME")
	if dbName == "" { dbName = "postgres" }

	// Bentuk connection string menggunakan nilai dari variabel lingkungan
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName,
	)

	// Perubahan pada inisialisasi koneksi
	Conn, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Gagal connect:", err)
	}

	err = Conn.Ping()
	if err != nil {
		log.Fatal("DB tidak bisa ping:", err)
	}

	log.Println("Database Connected")
}

func QueryExec( query string ) bool {
	_, err := Conn.Exec(query)
	if err != nil {
		cetak(err)
		return false
	}else{
		return true 
	}
}

func QueryExecReturnID(query string) (bool, int) {

	query = query + "RETURNING id"
	var idBaru int
	err := Conn.QueryRow(query).Scan(&idBaru)
	if err != nil {
		cetak("Gagal eksekusi query")
		cetak(err)
		return false, 0
	}else{
		return true, idBaru
	}
}

func QueryDataMahasiswa(query string) []DATA.DataMahasiswa {
	obj_sql, err := Conn.Query(query)
	if err != nil {
		return []DATA.DataMahasiswa{} //menghentikan laju fungsi
	}
	defer obj_sql.Close()
	//Jika SQL berhasil di eksesuksi
	var result [] DATA.DataMahasiswa
	for obj_sql.Next() {
		var data_row DATA.DataMahasiswa
		err := obj_sql.Scan(&data_row.Id, &data_row.Nama, &data_row.TanggalLahir, &data_row.Gender, &data_row.Jurusan, &data_row.Hobi )
		if err != nil {
			// cetak("SCAN ROW TIDAK ADA")
		}else{
			// cetak("SCAN ROW ADA")
		}
		result = append(result, data_row)
	}

	cetak(len( result ))
	//Mengembalikan array kosong kalo gak ada data, dan mengembalikan array index multi dimensi isi DATA.DataMahasiswa
	return result
}
func QueryDataRowMahasiswa(query string) *DATA.DataMahasiswa {
	data := QueryDataMahasiswa(query)
	if len(data) > 0 {
		return &data[0] // pointer ke row pertama
	} else {
		return nil
	}

	//Akan mengembalikan nilai row bentuk DATA.DataMahasiswa jika ada. Tapi kalo gak ada, akan mengembalikan nil
}
func Tambah_mahasiswa( input DATA.InputMahasiswa ) Response {

	//==== Menambah data mahasiswa =====
	query := fmt.Sprintf(
		"INSERT INTO mahasiswa (nama, tanggal_lahir, gender) VALUES ('%s', '%s', %d)",
		input.Nama,
		input.TanggalLahir,
		input.Gender,
	)

	response := Response{}
	param_db := true


	result, newIdMahasiswa := QueryExecReturnID( query )
	if result == true{
		//==== Menambah data mahasiswa_jurusan =====
		query = fmt.Sprintf(
			"INSERT INTO mahasiswa_jurusan (id_mahasiswa, id_jurusan) VALUES ( %d, %d )",
			newIdMahasiswa, input.IDJurusan,
		)
		result = QueryExec( query )
		if result == false{
			param_db = false
		}

		//==== Menambah data mahasiswa_hobi =====
		query = fmt.Sprintf(
			"INSERT INTO mahasiswa_hobi (id_mahasiswa, id_hobi) VALUES ( %d, %d )",
			newIdMahasiswa, input.IDHobi,
		)
		result = QueryExec( query )
		if result == false{
			param_db = false
		}

	}else{
		param_db = false
	}

	if param_db == true {
		response.Status = param_db 
		response.Msg = "Berhasil menambahkan data"
	}else{
		sql := fmt.Sprintf(
			"DELETE mahasiswa WHERE id = %id",
			newIdMahasiswa,
		)
		result = QueryExec( sql )

		sql = fmt.Sprintf(
			"DELETE mahasiswa_jurusan WHERE id_mahasiswa = %id",
			newIdMahasiswa,
		)
		result = QueryExec( sql )

		sql = fmt.Sprintf(
			"DELETE mahasiswa_hobi WHERE id_mahasiswa = %id",
			newIdMahasiswa,
		)
		result = QueryExec( sql )

	}

	response.Status = param_db
	cetak( response )
	return response
}   

func QueryDataJurusan(query string) []DATA.DataJurusan {
	obj_sql, err := Conn.Query(query)
	if err != nil {
		return []DATA.DataJurusan{} // Menghentikan fungsi jika query gagal
	}
	defer obj_sql.Close()

	var result []DATA.DataJurusan
	for obj_sql.Next() {
		var data_row DATA.DataJurusan
		err := obj_sql.Scan(&data_row.Id, &data_row.Nama)
		if err != nil {
			// bisa log error di sini jika mau
			cetak(err)
			continue
		}
		result = append(result, data_row)
	}

	cetak(len(result))
	return result
}

func Tambah_jurusan( input DATA.InputJurusan ) Response {

	response := Response{}
	query := fmt.Sprintf(
		"INSERT INTO jurusan (nama) VALUES ( '%s' )",
		input.Nama,
	)
	result, id := QueryExecReturnID( query )
	if result == true {
		response.Status = true
		response.Msg = "Berhasil menambahkan jurusan"
	}else{
		response.Status = false
		response.Msg = "Gagal menambahkan jurusan"
	}

	cetak( response )
	cetak( id )
	return response 
}


func QueryDataHobi(query string) []DATA.DataHobi {
	obj_sql, err := Conn.Query(query)
	if err != nil {
		return []DATA.DataHobi{} // Menghentikan fungsi jika query gagal
	}
	defer obj_sql.Close()

	var result []DATA.DataHobi
	for obj_sql.Next() {
		var data_row DATA.DataHobi
		err := obj_sql.Scan(&data_row.Id, &data_row.Nama)
		if err != nil {
			// bisa log error di sini jika mau
			cetak(err)
			continue
		}
		result = append(result, data_row)
	}

	cetak(len(result))
	return result
}

func Tambah_hobi( input DATA.InputHobi ) Response {
	response := Response{}
	query := fmt.Sprintf(
		"INSERT INTO hobi (nama) VALUES ( '%s' )",
		input.Nama,
	)
	result, id := QueryExecReturnID( query )
	if result == true {
		response.Status = true
		response.Msg = "Berhasil menambahkan hobi"
	}else{
		response.Status = false
		response.Msg = "Gagal menambahkan hobi"
	}

	cetak( response )
	cetak( id )
	return response 
}






