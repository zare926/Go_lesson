package main

import (
	"fmt"
	"log"
	"os"
	// "fmt"
	// "fmt"
	// "os"
)

func main() {

	// Exitは処理を終了させる
	// os.Exit(1)
	// fmt.Println("start")

	// log.Fatalもプログラムを終了させる
	// _, err := os.Open("A.txt")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// Args
	// fmt.Println(os.Args[0])
	// fmt.Println(os.Args[1])
	// fmt.Println(os.Args[2])
	// fmt.Println(os.Args[3])

	// fmt.Printf("length=%d\n", len(os.Args))

	// for i, v := range os.Args {
	// 	fmt.Println(i, v)
	// }

	// ファイル削除
	// Open

	// f, err := os.Open("test.txt")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// defer f.Close()

	// ファイル操作
	// Create

	// f, _ := os.Create("foo.txt")
	// f.Write([]byte("Hello\n"))
	// f.WriteAt([]byte("Golang"), 6)
	// f.Seek(0, os.SEEK_END)
	// f.WriteString("Yaah")

	// ファイル操作
	// Read
	// f, err := os.Open("foo.txt")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// defer f.Close()

	// bs := make([]byte, 128)
	// n, err := f.Read(bs)
	// fmt.Println(n)
	// fmt.Println(string(bs))

	// bs2 := make([]byte, 128)

	// nn, err := f.ReadAt(bs2, 11)
	// fmt.Println(nn)
	// fmt.Println(string(bs2))

	// OpenFile

	f, err := os.OpenFile("foo.txt", os.O_RDONLY, 06666)

	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	bs := make([]byte, 128)
	n, err := f.Read(bs)
	fmt.Println(n)
	fmt.Println(string(bs))
}
