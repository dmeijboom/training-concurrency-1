package main

func main() {
	url, close := startServer()
	defer close()

	for i := 0; i < 100; i++ {
		err := doRequest(url)
		if err != nil {
			panic(err)
		}
	}
}
