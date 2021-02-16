package helper

func SayHello(name string) string { // public, bisa diakses luar
	return "Selamat pagi " + name
}

func sayGoodbye(name string) string {  // private, ga bisa diakses luar
	return "Samapi jumpa " + name
}
