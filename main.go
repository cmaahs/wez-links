package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/muesli/termenv"
	"github.com/sirupsen/logrus"
)

func main() {
	byteData := []byte{}
	fi, _ := os.Stdin.Stat()
	if (fi.Mode() & os.ModeNamedPipe) != 0 {
		// inData := ""
		// nBytes, nChunks := int64(0), int64(0)
		r := bufio.NewReader(os.Stdin)
		buf := make([]byte, 0, 4*1024)

		// TODO: process input LINE by LINE
		for {

			n, err := r.Read(buf[:cap(buf)])
			buf = buf[:n]

			if n == 0 {

				if err == nil {
					continue
				}

				if err == io.EOF {
					break
				}

				logrus.WithError(err).Fatal("broken pipe")
			}

			// nChunks++
			// nBytes += int64(len(buf))

			// fmt.Println(string(buf))
			// inData = fmt.Sprintf("%s%s", inData, string(buf))
			byteData = append(byteData, buf...)

			if err != nil && err != io.EOF {
				logrus.WithError(err).Fatal("error reading pipe")
			}
		}

		// fmt.Println("Bytes:", nBytes, "Chunks:", nChunks)
		// fmt.Printf("%s\n", byteData)
	} else {
		logrus.Fatal("no piped input")
	}

	termFormatter := termenv.NewOutput(os.Stdout)

	// NAME                                                   READY   STATUS    RESTARTS   AGE
	var lineData bytes.Buffer

	lineData.Write(byteData)
	scanner := bufio.NewScanner(&lineData)
	firstline := true
	processor := ""
	for scanner.Scan() {
		line := scanner.Text()
		if firstline {
			processor = inputType(line)
			firstline = false
		} else {
			switch processor {
			case "pod_list":
				line = decoratePodLine(line, termFormatter)
			case "secret_list":
				line = decorateSecretLine(line, termFormatter)
			case "namespace_list":
				line = decorateNamespaceLine(line, termFormatter)
			}
		}
		fmt.Printf("%s\n", line)
	}
}
