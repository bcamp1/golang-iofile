package iofile

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// Ask - Prints out prompt, then returns user input
func Ask(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')

	if strings.HasSuffix(text, "\n") {
		text = text[:len(text)-len("\n")]
	}
	return text
}

// Get - Returns user input
func Get() string {
	return Ask("> ")
}

// GetBlank is like get but no arrow
func GetBlank() string {
	return Ask("")
}

// Append - Append to a file's text
func Append(fname string, message string) {
	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile(fname, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte(message)); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

// WriteAndReplace writes to a givn file
func WriteAndReplace(fname string, message string) {
	RemoveFile(fname)
	Append(fname, message)
}

// ReadLine reads a given line of file
func ReadLine(fname string, lineNum int) string {
	r, _ := os.Open(fname)
	var line string
	var lastLine int
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		lastLine++
		if lastLine == lineNum {
			// you can return sc.Bytes() if you need output in []bytes
			return sc.Text()
		}
	}
	return line
}

// PrintPage - Prints html code from given url
func PrintPage(url string) {
	fmt.Printf("HTML code of %s ...\n", url)
	resp, err := http.Get(url)
	// handle the error if there is one
	if err != nil {
		panic(err)
	}
	// do this now so it won't be forgotten
	defer resp.Body.Close()
	// reads html as a slice of bytes
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	// show the HTML code as a string %s
	fmt.Printf("%s\n", html)
}

// RemoveFile removes a given file "fname"
func RemoveFile(fname string) {
	err := os.Remove(fname)

	if err != nil {
		fmt.Println(err)
		return
	}
}

// FileExists checks if a given file exists
func FileExists(fname string) bool {
	if _, err := os.Stat("/path/to/whatever"); os.IsNotExist(err) {
		return false
	}
	return true
}

// PrintFile ...
func PrintFile(fname string) {

}
