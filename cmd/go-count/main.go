package main

var (
	portNumber string = "3000"
)

type application struct{}

func main() {

	app := &application{}

	app.setup()

	app.start()

}
