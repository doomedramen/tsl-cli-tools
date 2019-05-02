package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) == 3 {

		fromPath := os.Args[1]
		toPath := os.Args[2]

		if _, err := os.Stat(fromPath); os.IsNotExist(err) {
			// fromPath does not exist, BAD!
			fmt.Println(fromPath + " does not exist")
			return
		}

		if _, err := os.Stat(toPath); err == nil {
			// toPath exists, BAD!
			fmt.Println(toPath + " already exists")
			return
		}

		initSum, err := checksum(fromPath)
		check(err)

		copyError := copyFile(fromPath, toPath)
		check(copyError)

		postSum, err := checksum(toPath)
		check(err)

		if sliceEq(initSum, postSum) {
			//log.Println("copied ok :D")
			fmt.Println("copied ok (" + hex.EncodeToString(initSum) + " === " + hex.EncodeToString(postSum) + ")")
			return
		} else {
			fmt.Println("did not copy")

			//delete bad copy
			if _, err := os.Stat(toPath); err == nil {
				//exists, delete it
				err := os.Remove(toPath)
				check(err)
				return
			}
			return
		}
	} else {
		fmt.Println("Usage: sc fileA fileB")
	}
}

func checksum(filePath string) ([]byte, error) {
	var result []byte
	file, err := os.Open(filePath)
	if err != nil {
		return result, err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return result, err
	}

	return hash.Sum(result), nil
}

func copyFile(src string, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, in)
	cerr := out.Close()
	if err != nil {
		return err
	}
	return cerr
}

func sliceEq(a, b []byte) bool {

	if a == nil && b == nil {
		return true;
	}

	if a == nil || b == nil {
		return false;
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
		os.Exit(1)
	}
}
