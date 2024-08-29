package fscompare

import (
	"bufio"
	"fmt"
	"os"
)

func CompareFS(filePathOld, filePathNew string) {
	processAdded(filePathOld, filePathNew)
	processRemode(filePathOld, filePathNew)
}

func processRemode(filePathOld, filePathNew string) {

	fileOld, err := os.Open(filePathOld)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer fileOld.Close()

	fileNew, err := os.Open(filePathNew)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer fileNew.Close()

	scannerOld := bufio.NewScanner(fileOld)
	scannerNew := bufio.NewScanner(fileNew)

	for scannerOld.Scan() {
		found := false
		path1 := scannerOld.Text()

		_, err := fileNew.Seek(0, 0)
		if err != nil {
			println("Error seeking file:", err)
			return
		}

		for scannerNew.Scan() {
			path2 := scannerNew.Text()
			if path1 == path2 {
				found = true
				break
			}
		}

		if !found {
			fmt.Println("REMOVED:", path1)
		}

	}
}

func processAdded(filePathOld, filePathNew string) {

	fileOld, err := os.Open(filePathOld)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer fileOld.Close()

	fileNew, err := os.Open(filePathNew)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer fileNew.Close()

	scannerOld := bufio.NewScanner(fileOld)
	scannerNew := bufio.NewScanner(fileNew)

	for scannerNew.Scan() {
		found := false
		path1 := scannerNew.Text()

		_, err := fileOld.Seek(0, os.SEEK_SET)
		if err != nil {
			println("Error seeking file:", err)
			return
		}

		for scannerOld.Scan() {
			path2 := scannerOld.Text()
			if path1 == path2 {
				found = true
				break
			}
		}
		if !found {
			fmt.Println("ADDED:", path1)
		}
	}
}
