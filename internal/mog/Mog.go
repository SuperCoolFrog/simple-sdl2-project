package mog

import (
	"io"
	"os"
	"time"
)

/*
fmt.Fprintf(mog.MW, "\n %s something %d, once = %t", me.Id, 2, true)

fmt.Fprintln(mog.MW, "\n hello")
*/
var MW io.Writer = nil

func Init() error {
	dir := "mogs/"
	filePath := dir + time.Now().Format("2006.01.02.15.04.05") + ".mog.txt"
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	MW = io.MultiWriter(os.Stdout, file)
	return nil
}
