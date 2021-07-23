// B"H
/*
/*
Package utils
*/
package utils

import (
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	external "github.com/dustin/go-humanize"
)

var quietDownload bool

// ---
// DownloadFile helpers from same source
// ---
// WriteCounter counts the number of bytes written to it. It implements to the io.Writer interface
// and we can pass this into io.TeeReader() which will report progress on each write cycle.
type WriteCounter struct {
	Total uint64
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.PrintProgress()
	return n, nil
}

func (wc WriteCounter) PrintProgress() {
	if !quietDownload {
		// Clear the line by using a character return to go back to the start and remove
		// the remaining characters by filling it with spaces
		fmt.Printf("\r%s", strings.Repeat(" ", 35))
		// Return again and print current status of download
		// We use the humanize package to print the bytes in a meaningful way (e.g. 10 MB)
		fmt.Printf("\rDownloading... %s complete", external.Bytes(wc.Total))
	}
}

// ---
// End of helpers
// ---

/*
DownloadFile
https://golangcode.com/download-a-file-with-progress/
This code is copy paste with quiet flag that I added from the above blog post and it working great.
*/
// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory. We pass an io.TeeReader
// into Copy() to report progress on the download.
func DownloadFile(filepath string, url string, quiet bool) error {

	// quietDownload := "dsf"
	if quiet {
		quietDownload = true
	}

	// TODO add flag
	var insecureSkipVerify bool = true
	// Skip ssl vrification.
	httpTransport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: insecureSkipVerify},
		Proxy:           http.ProxyFromEnvironment,
	}
	client := &http.Client{Transport: httpTransport}

	// Create the file, but give it a tmp file extension, this means we won't overwrite a
	// file until it's downloaded, but we'll remove the tmp extension once downloaded.
	out, err := os.Create(filepath + ".tmp")
	if err != nil {
		return err
	}
	// Get the data

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		out.Close()
		return err
	}

	defer resp.Body.Close()
	// Create our progress reporter and pass it to be used alongside our writer
	counter := &WriteCounter{}
	if _, err = io.Copy(out, io.TeeReader(resp.Body, counter)); err != nil {
		out.Close()
		return err
	}
	// The progress use the same line so print a new line once it's finished downloading
	fmt.Print("\n")
	// Close the file without defer so it can happen before Rename()
	out.Close()
	if err = os.Rename(filepath+".tmp", filepath); err != nil {
		return err
	}

	fileModeNumber := int(0755)
	fileMode := os.FileMode(fileModeNumber)
	if err = os.Chmod(filepath, fileMode); err != nil {
		return err
	}

	return nil
}

// GetFileContent return response body (without closing body no reason for it)
// Actually I see now the this is just get request function that needs a little touch up
func GetFileContent(url string) []byte {

	// TODO add flag
	var insecureSkipVerify bool = true
	// Skip ssl vrification.
	httpTransport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: insecureSkipVerify},
		Proxy:           http.ProxyFromEnvironment,
	}

	// Create http request
	client := &http.Client{Transport: httpTransport}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Do the http request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	// resp, err := http.Get(url)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return body
}
