package download

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"regexp"
	"strings"
	"sync"
)

func Worker(work chan string, wg *sync.WaitGroup) {
	for workUrl := range work {
		DownloadUrl(workUrl)
	}
	wg.Done()
}

func DownloadUrl(workUrl string) {
	u, err := url.Parse(workUrl)
	if err != nil {
		fmt.Printf("Invalid URL %s, skipping.\n", workUrl)
		return
	}
	//parentDir := downloadDir
	ext := path.Ext(u.Path)
	// fmt.Printf("ext: %s\n", ext)
	if IsSanitizeUrl {
		sanitizedName := SanitizeUrl(u.Scheme + u.Host + TrimSuffix(u.Path, ext) + "?" + u.RawQuery)
		fmt.Printf("sanitizedName: %s\n", sanitizedName)
	}
	// filePath := u.Path + "/"

	filePath, err := GetPrefixPath(u.Path)
	if err != nil {
		fmt.Println("Error while getting the destination file prefix path, skipping:", err)
		return
	}

	fileName := u.Path // filePath + ext
	fullPath := fmt.Sprintf("%v%s", DownloadDir, filePath)
	fullName := fmt.Sprintf("%v%s", DownloadDir, fileName) // parentDir + fileName

	err = os.MkdirAll(fullPath, 0700)
	if err != nil {
		fmt.Println("Error creating directory, skipping:", err)
		return
	}

	_, err = os.Stat(fullName)
	if err == nil {
		fmt.Printf("File %s already exists, skipping\n", fileName)
		return
	}

	out, err := os.Create(fullName)
	if err != nil {
		fmt.Println("Error creating file, skipping:", err)
		return
	}
	defer out.Close()

	resp, err := http.Get(workUrl)
	if err != nil {
		fmt.Println("Error getting URL, skipping:", err)
		return
	}
	defer resp.Body.Close()

	bytes, err := io.Copy(out, resp.Body)
	if err != nil {
		fmt.Println("Error writing file, skipping:", err)
		return
	}

	fmt.Printf("%d bytes: %s\n", bytes, fileName)
}

func SanitizeUrl(url string) string {
	return regexp.MustCompile("\\s+").ReplaceAllLiteralString(StripChars(url), "-")
}

func StripChars(str string) string {
	toStrip := [...]string{
		"~", "`", "!", "@", "#", "$", "%", "^", "&", "*",
		"(", ")", "_", "=", "+", "[", "{", "]", "}", "\\",
		"|", ";", ":", "\"", "'", "&#8216;", "&#8217;",
		"&#8220;", "&#8221;", "&#8211;", "&#8212;", "—",
		"–", ",", "<", ".", ">", "/", "?",
	}

	for _, char := range toStrip {
		str = strings.Replace(str, char, "", -1)
	}

	return str
}

func TrimSuffix(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		return s[:len(s)-len(suffix)]
	}

	return s
}

func AddWorkFromCsv(work chan string, filePath string) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	csvReader := csv.NewReader(file)

	for {
		line, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		work <- line[0]
	}
}
