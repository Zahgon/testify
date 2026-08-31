//go:build go1.21

package main

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"regexp"
	"slices"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("readme-gofmt: ")

	buf, err := os.ReadFile("README.md")
	if err != nil {
		log.Fatal(err)
	}

	gofmt, err := exec.LookPath("gofmt")
	if err != nil {
		log.Fatal(err)
	}

	buf = bytes.ReplaceAll(buf, []byte("\r"), nil)

	reBlock := regexp.MustCompile("(?s)\n```go\n(.*?\n)```")

	matches := reBlock.FindAllSubmatchIndex(buf, -1)

	changes := 0

	for i := len(matches) - 1; i >= 0; i-- {
		match := matches[i]
		block := buf[match[2]:match[3]]

		cmd := exec.Command(gofmt, "-s")
		cmd.Stdin = bytes.NewReader(block)
		var output bytes.Buffer
		cmd.Stdout = &output
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}

		outputBytes := bytes.ReplaceAll(output.Bytes(), []byte("\r"), []byte(""))
		if !bytes.Equal(outputBytes, block) {
			changes++
			buf = slices.Replace(buf, match[2], match[3], outputBytes...)
		}
	}

	if changes > 0 {
		cmd := exec.Command("diff", "-a", "-u", "README.md", "-")
		cmd.Stdin = bytes.NewReader(buf)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()

		os.Exit(2)
	}
}
