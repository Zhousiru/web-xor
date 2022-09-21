package main

import (
	"log"
	"net/http"
	"os"
)

const key = "testkey"

func xor(b []byte) []byte {
	bKey := []byte(key)
	keyLen := len(bKey)

	var r []byte
	for index, v := range b {
		r = append(r, v^bKey[index%keyLen])
	}

	return r
}

func handleText(w http.ResponseWriter, r *http.Request) {
	text, err := os.ReadFile("./assets/text.txt")
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(xor(text))
}

func handleImage(w http.ResponseWriter, r *http.Request) {
	image, err := os.ReadFile("./assets/image.png")
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(xor(image))
}

func main() {
	http.HandleFunc("/text", handleText)
	http.HandleFunc("/image", handleImage)

	log.Println("listening on :8080")
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
