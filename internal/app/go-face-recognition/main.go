package main

import(
	"fmt"
	"github.com/Kagami/go-face"
)

const dataDir = "testdata"

func main() {
	fmt.Println("Facial Recognition System v0.01")

	rec, err := face.NewRecognizer(dataDir)
	if err != nil {
		fmt.Println("Cannot Initialize recognizer")
	}
	defer rec.Close()

	fmt.Println("Recognizer Initialized")
}