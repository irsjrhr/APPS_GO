package main

import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
    "APPS_TEST/DB" 
    "APPS_TEST/DATA"
)

func cetak( nilai interface{} ) {
    fmt.Println( nilai )
}

func main() {

    DB.Init()


    // Inisialisasi router Gin
    router := gin.Default()

    //=============================== ROUTE MAHASISWA ============================
    router.GET("/mahasiswa", func(c *gin.Context) {
        data_mahasiswa := DB.QueryDataMahasiswa( "SELECT  m.id AS mahasiswa_id, m.nama AS mahasiswa_nama, m.tanggal_lahir, m.gender, j.nama AS jurusan_nama, h.nama AS hobi_nama FROM mahasiswa m LEFT JOIN mahasiswa_jurusan mj ON m.id = mj.id_mahasiswa LEFT JOIN jurusan j ON mj.id_jurusan = j.id LEFT JOIN mahasiswa_hobi mh ON m.id = mh.id_mahasiswa LEFT JOIN hobi h ON mh.id_hobi = h.id ORDER BY m.id")
        var msg string 
        if len( data_mahasiswa ) > 0 {
            msg = "Berhasil mengambil data"
        }else {
            msg = "Data tidak ditemukan!"
        }
        c.JSON(http.StatusOK, gin.H{
            "status": true,
            "message": msg,
            "data" : data_mahasiswa,
        })
    })
    router.GET("/mahasiswa/:id", func(c *gin.Context) {
        id := c.Param("id")
        row_mahasiswa := DB.QueryDataRowMahasiswa( "SELECT  m.id AS mahasiswa_id, m.nama AS mahasiswa_nama, m.tanggal_lahir, m.gender, j.nama AS jurusan_nama, h.nama AS hobi_nama FROM mahasiswa m LEFT JOIN mahasiswa_jurusan mj ON m.id = mj.id_mahasiswa LEFT JOIN jurusan j ON mj.id_jurusan = j.id LEFT JOIN mahasiswa_hobi mh ON m.id = mh.id_mahasiswa LEFT JOIN hobi h ON mh.id_hobi = h.id WHERE m.id="+id+";")   
        var msg string 
        if row_mahasiswa != nil {
            msg = "Berhasil mengambil data"
        }else {
            msg = "Data tidak ditemukan!"
        }
        c.JSON(http.StatusOK, gin.H{
            "status": true,
            "message": msg,
            "data" : row_mahasiswa,
            "id" : id,
        })
    })
    router.POST("/mahasiswa", func(c *gin.Context) {
        var input DATA.InputMahasiswa

        // Bind JSON input ke struct
        if err := c.ShouldBindJSON(&input); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "status": false,
                "message": "JSON tidak valid",
                "error":   err.Error(),
            })
            return
        }   

        //Tambah Data 
        tambah_msh := DB.Tambah_mahasiswa( input )
        c.JSON(http.StatusOK, gin.H{
            "status" : tambah_msh.Status,
            "message": "Data mahasiswa berhasil disimpan",
            "data":    input,
        })
    })




    //=============================== ROUTE JURUSAN ============================
    router.GET("/jurusan", func(c *gin.Context) {
        data_jurusan := DB.QueryDataJurusan( "SELECT * FROM jurusan")
        var msg string 
        if len( data_jurusan ) > 0 {
            msg = "Berhasil mengambil data"
        }else {
            msg = "Data tidak ditemukan!"
        }
        c.JSON(http.StatusOK, gin.H{
            "status": true,
            "message": msg,
            "data" : data_jurusan,
        })
    })
    router.POST("/jurusan", func(c *gin.Context) {
        var input DATA.InputJurusan

        // Bind JSON input ke struct
        if err := c.ShouldBindJSON(&input); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "status": false,
                "message": "JSON tidak valid",
                "error":   err.Error(),
            })
            return
        }   

        //Tambah Data 
        tambah_jurusan := DB.Tambah_jurusan( input )
        c.JSON(http.StatusOK, gin.H{
            "status" : tambah_jurusan.Status,
            "message": "Data jurusan berhasil disimpan",
            "data":    input,
        })
    })



    //=============================== ROUTE HOBI ============================
    router.GET("/hobi", func(c *gin.Context) {
        data_hobi := DB.QueryDataHobi( "SELECT * FROM hobi")
        var msg string 
        if len( data_hobi ) > 0 {
            msg = "Berhasil mengambil data"
        }else {
            msg = "Data tidak ditemukan!"
        }
        c.JSON(http.StatusOK, gin.H{
            "status": true,
            "message": msg,
            "data" : data_hobi,
        })
    })
    router.POST("/hobi", func(c *gin.Context) {
        var input DATA.InputHobi

        // Bind JSON input ke struct
        if err := c.ShouldBindJSON(&input); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "status": false,
                "message": "JSON tidak valid",
                "error":   err.Error(),
            })
            return
        }   

        //Tambah Data 
        tambah_hobi := DB.Tambah_hobi( input )
        c.JSON(http.StatusOK, gin.H{
            "status" : tambah_hobi.Status,
            "message": "Data hobi berhasil disimpan",
            "data":    input,
        })
    })

    // Jalankan server di port 8080
    router.Run(":8080")

}



