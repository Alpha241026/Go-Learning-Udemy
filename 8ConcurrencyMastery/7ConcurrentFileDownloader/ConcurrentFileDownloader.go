package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func DownloadFile(url, destDir string) error {
	//extract the filename from the URL
	filename := filepath.Base(url)
	//build the full destination path
	filePath := filepath.Join(destDir, filename)

	//create the destination file on disk
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	fmt.Println("Downloading", url)
	start := time.Now()

	//download the file from the remote server
	resp, err := http.Get(url)
	if err != nil {
		_ = os.Remove(filePath)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		_ = os.Remove(filePath)
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	//copy the downloaded bytes into the local file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("Download %s took %s\n", filename, time.Since(start))
	return nil
}

func SequentialDownloader(urls []string, destDir string) error {
	//ensure the destination directory exists
	if err := os.MkdirAll(destDir, 0755); err != nil {
		return err
	}

	start := time.Now()
	//download files one after another
	for _, url := range urls {
		if err := DownloadFile(url, destDir); err != nil {
			fmt.Println("Error downloading", url, err)
			continue
		}
	}
	fmt.Printf("Download %s took %s\n", urls, time.Since(start))
	return nil
}

// report returned by each download worker
type Result struct {
	URL      string
	Filename string
	Size     int64
	Duration time.Duration
	Error    error
}

func ConcurrentDownloader(urls []string, destDir string, maxConcurrent int) error {
	if err := os.MkdirAll(destDir, 0755); err != nil {
		return err
	}

	//channel used by workers to send download results
	results := make(chan Result)

	//wait for all download workers to finish
	var wg sync.WaitGroup

	// Limit the number of simultaneous downloads
	limiter := make(chan struct{}, maxConcurrent)
	for _, url := range urls {
		wg.Add(1)
		//launch worker per url; pass the loop variable as a parameter to avoid closure issues
		go func(url string) {
			defer wg.Done()

			//acquire a concurrency slot
			limiter <- struct{}{}

			//release the slot when finished
			defer func() { <-limiter }()

			start := time.Now()
			filename := filepath.Base(url)
			filePath := filepath.Join(destDir, filename)

			out, err := os.Create(filePath)
			if err != nil {
				//every worker reports its outcome through the shared results channel
				results <- Result{URL: url, Error: err}
				return
			}
			defer out.Close()

			resp, err := http.Get(url)
			if err != nil {
				results <- Result{URL: url, Error: err}
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				results <- Result{URL: url, Error: fmt.Errorf("bad status: %s", resp.Status)}
				return
			}

			size, err := io.Copy(out, resp.Body)
			if err != nil {
				results <- Result{URL: url, Error: err}
				return
			}
			//measure how long this individual download takes
			timeSince := time.Since(start)
			results <- Result{URL: url, Filename: filename, Size: size, Duration: timeSince, Error: nil}
		}(url)
	}

	//close the results channel after all workers complete
	go func() {
		wg.Wait()
		close(results)
	}()

	var totalSize int64
	var errors []error
	start := time.Now()

	//process results until the channel is closed
	for result := range results {
		//aggregate successful downloads and collect any errors.
		if result.Error != nil {
			fmt.Printf("Error downloading %s: %s\n", result.URL, result.Error.Error())
			errors = append(errors, result.Error)
		} else {
			totalSize += result.Size
			fmt.Printf("Downloaded %s (%d bytes) in %s\n", result.Filename, result.Size, result.Duration)
		}
	}

	startedSince := time.Since(start)
	fmt.Printf("All downloads completed in %s, Total: %d bytes\n", startedSince, totalSize)
	if len(errors) > 0 {
		return fmt.Errorf("errors downloading: %+v", errors)
	}
	return nil
}

func main() {
	//example URLs to download
	urls := []string{"https://file-examples.com/storage/fe062d525c682cad199a167/2017/10/file_example_JPG_1MB.jpg", "https://go.dev/images/go-logo-white.svg"}

	err := ConcurrentDownloader(urls, "./", 3)
	if err != nil {
		fmt.Println(err)
		return
	}

	log.Println("Done")
}
