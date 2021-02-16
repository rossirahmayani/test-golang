package database

var connection string

func init()  {
	connection = "MySQL" // utk inisialisasi variabel
}

func GetDatabase() string {
	return connection
}
