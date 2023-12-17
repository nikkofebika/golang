package helper

var version = "1.0.0" // tidak bisa diakses dari luar package

var Version = "1.0.0"

// hanya bisa diakses dari package yang sama
func sayGoodBye(name string) string {
	return "Good bye " + name
}

func SayHello(name string) string {
	return "Hallo " + name
}
