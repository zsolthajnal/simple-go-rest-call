package main

func main() {
	targetIP, err := readParams()
	if err != nil {
		panic(err)
	}

	if err := restGetCall(targetIP); err != nil {
		panic(err)
	}
}
