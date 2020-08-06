package main

import (
	"bytes"
	"fmt"
	"github.com/otiai10/gosseract/v2"
	"math/rand"
	"path/filepath"

	"os"
	"os/exec"
	"strconv"
	"time"
)

func main() {
	st := makeTimestamp()
	tmpDir := os.TempDir()
	tmpDir = filepath.Join(tmpDir, RandStringRunes(5))
	err := os.MkdirAll(tmpDir, 0775)
	if err != nil {
		panic(err)
	}
	fmt.Println(tmpDir)
	defer os.RemoveAll(tmpDir)

	fileTemplate := filepath.Join(tmpDir, "%d.png")

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("mutool", "draw", "-F", "png", "-o", fileTemplate, os.Args[1], "1-10")
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err = cmd.Run()
	fmt.Println(err)
	fmt.Println(stdout.String(), stderr.String())

	for i := 1; i <= 10; i++ {
		name := strconv.Itoa(i)
		name += ".png"
		fileName := filepath.Join(tmpDir, name)
		fmt.Println("=================== file name: ", fileName)
		ocrImg(fileName)
	}

	et := makeTimestamp()
	fmt.Printf("time: %dsm\n", et-st)
}

func ocrImg(file string) {

	client := gosseract.NewClient()
	defer client.Close()
	//client.SetLanguage("chi_sim+eng")
	client.SetImage(file)
	fmt.Println(client.Text())

}

func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
