package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var id = 50001

func main() {
	err := filepath.Walk("./temp", walkFunc)
	if err != nil {
		fmt.Println(err)
	}
}

func walkFunc(path string, info os.FileInfo, err error) error {
	extension := getExtension(path)
	if extension != "jpg" {
		return nil
	}
	//fmt.Println(path)
	//fmt.Println("temp/" + strconv.Itoa(id) + ".jpg")

	err = os.Rename(path, "temp/"+strconv.Itoa(id)+".jpg")
	id++

	return nil
}

func getExtension(path string) string {
	pos := strings.LastIndex(path, ".") + 1
	return path[pos:]
}
