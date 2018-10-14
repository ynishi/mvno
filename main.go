package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func createMoveToPath(prefix, oldpath, newdir string) string {
	j := 0
	candidateFilename := createName(prefix, j, oldpath)
	for !isAvailable(newdir, candidateFilename) {
		j = j + 1
		candidateFilename = createName(prefix, j, oldpath)
	}
	return filepath.Join(newdir, candidateFilename)
}

func createName(prefix string, number int, filename string) string {
	numberStr := strconv.Itoa(number)
	ext := filepath.Ext(filename)
	created := fmt.Sprintf("%s%s%s", prefix, numberStr, ext)
	return created
}

func isAvailable(dir string, filename string) bool {
	path := filepath.Join(dir, filename)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return true
	} else {
		return false
	}
}

func moveByCopy(oldPath, destPath string) error {
	in, err := os.Open(oldPath)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer func() {
		cerr := out.Close()
		if cerr != nil {
			if err != nil {
				err = fmt.Errorf("%v:%v", cerr, err)
			} else {
				err = cerr
			}
		}
		if err != nil {
			if _, err := os.Stat(destPath); os.IsExist(err) {
				rerr := os.Remove(destPath)
				if rerr != nil {
					if err != nil {
						err = fmt.Errorf("%v:%v", rerr, err)
					} else {
						err = rerr
					}
				}
			}
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return err
	}
	if err = out.Sync(); err != nil {
		return err
	}
	in.Close()
	if err = os.Remove(oldPath); err != nil {
		return err
	}
	return nil
}

func main() {
	if len(os.Args) < 4 {
		log.Fatal("required args,$1:prefix,$2..n-1:target files $n: dir to move in.")
	}
	newDir := os.Args[len(os.Args)-1]
	prefix := os.Args[1]
	for i := 2; i <= len(os.Args)-2; i++ {
		oldPath := os.Args[i]
		moveToPath := createMoveToPath(prefix, oldPath, newDir)
		if err := moveByCopy(oldPath, moveToPath); err != nil {
			log.Fatal(fmt.Sprintf("failed to mv,from:%v:to:%v", oldPath, newDir), err)
		}
	}
}
