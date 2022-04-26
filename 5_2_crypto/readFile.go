package main

import (
	// "flag"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// readFile()
	hashFile()

	// Compare by sha256
	log.Println("Compare")
	t1 := fileCompareByChunk("image/png", "image/png")

}

func readFile() {
	content, err := ioutil.ReadFile("hello.txt")
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("Hello W")
	fmt.Print(content)
}

func hashFile() {
	// argsWithProg := os.Args
	// argsWithoutProg := os.Args[1:]
	arg_1 := os.Args[1]
	arg_2 := os.Args[2]
	arg_3 := os.Args[3]
	// // fmt.Println("var1 = ", reflect.TypeOf(arg))
	// file, err := os.Open(arg)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()

	// // fmt.Println(argsWithProg)
	// // fmt.Println(argsWithoutProg)
	// // fmt.Println(arg)

	// hash := sha256.New()
	// if _, err := io.Copy(hash, file); err != nil {
	// 	log.Fatal(err)
	// }

	photo_1, err := ioutil.ReadFile(arg_1)

	f, err := os.Open("odd.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("%x", h.Sum(nil))
	fmt.Printf("%x", h.Sum(nil))
}

func compareBySha256Simple()  {
	content1, err1 := ioutil.ReadFile("odd.txt")
	if err1 != nil {
		log.Fatal(err1)
	}
	content2, err2 := ioutil.ReadFile("odd.txt")
	if err2 != nil {
		log.Fatal(err2)
	}
	content3, err3 := ioutil.ReadFile("odd.txt")
	if err3 != nil {
		log.Fatal(err3)
	}

	sum1 := sha256.Sum256(content1)
	sum2 := sha256.Sum256(content2)
	sum3 := sha256.Sum256(content2)

	equal1 := sum1 == sum2
	equal2 := sum1 == sum3
	equal3 := sum2 == sum2

	fmt.Println("image 1 et 2: ", equal1)
	fmt.Println("image 1 et 3: ", equal2)
	fmt.Println("image 2 et 3: ", equal3)
}

func fileCompareByChunk(file1, file2 string) bool {
	// _ : ignorer la var retourner à cet endroit

	/*
	os.Open // Ouvrir le fichier
	bufio.NewReader // J'ai mis un lecteur
	reeader.Read // Lire
	
	*/
	return false
}