package main

import (
"crypto/sha512"
"fmt"
"io/ioutil"
"os"
"path/filepath"
)

var files = make(map[[sha512.Size]byte]string)
var scannedFiles = make(map[[sha512.Size]byte]string)
var hashMaporiginal = make(map[[sha512.Size]byte]string)

<<<<<<< HEAD

/**
func checkDuplicate(pathOriginal string, pathCopy string, info os.FileInfo, err error) error {
=======
func checkDuplicate(path string, info os.FileInfo, err error) error {
>>>>>>> master

	if err != nil {
		fmt.Println(err)
		return nil
	}

	if info.IsDir() { //skip folder
		return nil
	}

	data, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	hash := sha512.Sum512(data) //get the file sha512 hash
<<<<<<< HEAD
	files[hash] = pathOriginal  //store in map for comparison

	//read second and check path and

	dataCopy, err := ioutil.ReadFile(pathCopy)

	if v, ok := files[hash]; ok {
		fmt.Printf("%q is a duplicate of %q\n", pathOriginal, v)
		//os.Remove(path)
=======

	if v, ok := files[hash]; ok {
		fmt.Printf("%q is a duplicate of %q\n", path, v)
		os.Remove(path)
	} else {
		files[hash] = path //store in map for comparison
>>>>>>> master
	}

	return nil
}
**/

func scanFiles(path string, hashMap map[[sha512.Size]byte]string)  {

	var scan = func(path string, fileInfo os.FileInfo, ImpErr error) (err error) {
		if fileInfo.IsDir() { //skip folder
			return nil
		}

		data, err := ioutil.ReadFile(path)

		if err != nil {
			fmt.Println(err)
			return nil	
		}

		hash := sha512.Sum512(data) //get the file sha512 hash
		hashMap[hash] = path  


		return

	}

	err := filepath.Walk(path, scan)

	if err != nil {
		fmt.Println(err)
	}

}

func compareFiles(path string, hashMap map[[sha512.Size]byte]string)  { 

	var scan = func(path string, fileInfo os.FileInfo, ImpErr error) (err error) {
		if fileInfo.IsDir() { //skip folder
			return nil
		}

		data, err := ioutil.ReadFile(path)

		if err != nil {
			fmt.Println(err)
			return nil	
		}

		hash := sha512.Sum512(data) //get the file sha512 hash

		if v, ok := hashMap[hash]; ok {
			fmt.Printf("%q is a duplicate of %q\n", path, v)
		//os.Remove(path)
		} else {
			fmt.Printf("%q not a in copy_directory\n", path)
		}


		return
	}


	err := filepath.Walk(path, scan)

	if err != nil {
		fmt.Println(err)
	}

}



func main() {

	if len(os.Args) != 2 {
		fmt.Print("USAGE: %s <target_directory> \n", os.Args[0])
		os.Exit(0)
	}
	dir := os.Args[1]

	fmt.Print("Begin checking all files in original directory\n")

	pathOriginal := os.Args[1]
	pathCopy := os.Args[2]

	scanFiles(pathOriginal, scannedFiles)

	compareFiles(pathCopy, scannedFiles)

}
