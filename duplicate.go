package main

import (
	"crypto/sha512"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var files = make(map[[sha512.Size]byte]string)

func storeHash(path string, info os.FileInfo, err error) map[[64]byte]string {
	if err != nil {
		fmt.Println(err)

	}

	if info.IsDir() { //skip folder
		fmt.Println("is folder")
	}

	data, err := ioutil.ReadFile(path)

	hash := sha512.Sum512(data) //get the file sha512 hash
	files[hash] = path          //store in map for comparison

	return files
}

func checkDuplicate(pathOriginal string, pathCopy string, info os.FileInfo, err error) error {

	if err != nil {
		fmt.Println(err)
		return nil
	}

	if info.IsDir() { //skip folder
		return nil
	}

	//read all pathOriginal and store the hash

	data, err := ioutil.ReadFile(pathOriginal)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	hash := sha512.Sum512(data) //get the file sha512 hash
	files[hash] = pathOriginal  //store in map for comparison

	//read second and check path and

	dataCopy, err := ioutil.ReadFile(pathCopy)

	if v, ok := files[hash]; ok {
		fmt.Printf("%q is a duplicate of %q\n", pathOriginal, v)
		//os.Remove(path)
	}

	return nil
}

func main() {

	if len(os.Args) != 3 {
		fmt.Print("USAGE: %s <original_directory> <copy_directory\n", os.Args[0])
		os.Exit(0)
	}

	pathOriginal := os.Args[1]
	pathCopy := os.Args[2]

	hash := filepath.Walk(pathOriginal, checkDuplicate)
	err2 := filepath.Walk(pathCopy, checkDuplicate)

	if err1 != nil {
		fmt.Println(err1)
		os.Exit(1)
	}

	if err2 != nil {
		fmt.Println(err2)
		os.Exit(1)
	}
}
