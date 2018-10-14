package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func createMoveToPath(prefix string, start int, format, oldpath, newdir string) (string, int) {
	j := start
	candidateFilename := createName(prefix, j, format, oldpath)
	for !isAvailable(newdir, candidateFilename) {
		j = j + 1
		candidateFilename = createName(prefix, j, format, oldpath)
	}
	return filepath.Join(newdir, candidateFilename), j
}

func createName(prefix string, number int, numberFormat, filename string) string {
	ext := filepath.Ext(filename)
	nameFormat := "%s" + numberFormat + "%s"
	created := fmt.Sprintf(nameFormat, prefix, number, ext)
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

var (
	numFormatOpt = flag.String("format", "%d", "format of number(supported printf int)")
	startNumOpt  = flag.Int("start", 0, "start number of filename, default:0")
	prefixOpt    = flag.String("prefix", "", "prefix of filename, default:\"\"")
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s (-h):\n", os.Args[0])
		fmt.Printf("  %s [OPTIONS] $1:..n-1(target files) $n(dir to move in)\n", os.Args[0])
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 2 {
		flag.Usage()
		log.Fatal("args required")
	}
	newDir := args[len(args)-1]
	prefix := *prefixOpt
	start := *startNumOpt
	format := *numFormatOpt
	for i := 0; i <= len(args)-2; i++ {
		oldPath := args[i]
		moveToPath, latest := createMoveToPath(prefix, start, format, oldPath, newDir)
		start = latest + 1
		if err := moveByCopy(oldPath, moveToPath); err != nil {
			log.Fatal(fmt.Sprintf("failed to mv,from:%v:to:%v:", oldPath, newDir), err)
		}
	}
}
