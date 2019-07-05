package gomu

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/vansante/go-ffprobe"
)

var (
	fileInfo os.FileInfo
	err      error
)

type Ffprobe struct {
	Format struct {
		Filename       string `json:"filename"`
		NbStreams      int    `json:"nb_streams"`
		NbPrograms     int    `json:"nb_programs"`
		FormatName     string `json:"format_name"`
		FormatLongName string `json:"format_long_name"`
		StartTime      string `json:"start_time"`
		Duration       string `json:"duration"`
		Size           string `json:"size"`
		BitRate        string `json:"bit_rate"`
		ProbeScore     int    `json:"probe_score"`
		Tags           struct {
			MajorBrand       string    `json:"major_brand"`
			MinorVersion     string    `json:"minor_version"`
			CompatibleBrands string    `json:"compatible_brands"`
			CreationTime     time.Time `json:"creation_time"`
		} `json:"tags"`
	} `json:"format"`
}

type CsvLine struct {
	Column1 string
	Column2 string
	Column3 string
}

// RemoveDuplicates : Remove duplicates from a []string
func RemoveDuplicates(elements []string) []string {
	encountered := map[string]bool{}
	result := []string{}

	for v := range elements {
		if encountered[elements[v]] == true {
		} else {
			encountered[elements[v]] = true
			result = append(result, elements[v])
		}
	}
	return result
}

// CreatePNG :  .png file from source content
func CreatePNG(mp4 string, png string) {
	fmt.Println("-- CREATE PNG -- ", png)
	cmd := exec.Command("ffmpeg", "-y", "-ss", "0", "-t", "11", "-i", mp4, "-filter_complex", "[0:v] palettegen", png)
	cmd.Run()
}

// CreateStill : Create single .jpg file from a source
func CreateStill(mp4 string, still string) {
	fmt.Println("-- CREATE STILL -- ", still[:len(still)-4])
	CreateDirectories("jpg")
	newFilename := still[:len(still)-4] + "-still.jpeg"
	cmd5 := exec.Command("ffmpeg", "-i", mp4, "-f", "image2", newFilename)

	cmd5.Run()
}

// CreateStillBundle :  Generate a bundle of .jpg files based on frame extraction
func CreateStillBundle(mp4 string, still string) {
	fmt.Println("-- CREATE JPGS -- ", still[:len(still)-4])
	isolatedVideo := mp4[59 : len(mp4)-4]
	trimDir := strings.SplitAfter(still[:len(still)-4], "/edits")
	removeEditDir := trimDir[0][:len(trimDir[0])-5]
	newDir := removeEditDir + "jpg/" + isolatedVideo

	fmt.Println(newDir)

	CreateDirectories(newDir)

	check := newDir + "/" + isolatedVideo + "-frame-%04d.jpg"
	cmd5 := exec.Command("ffmpeg", "-i", mp4, check)

	cmd5.Run()
}

// MoveStill : Move .png to folder
func MoveFile(source string, destination string) {
	fmt.Println("-- MOVE STILL --")
	fmt.Println(destination)
	cmd4 := exec.Command("mv", source, destination)
	cmd4.Run()
}

// MovePNG : Move .png to folder
func MovePNG(source string, destination string) {
	fmt.Println("-- MOVE PNG --")
	fmt.Println(destination)
	cmd4 := exec.Command("mv", source, destination)
	cmd4.Run()
}

// BytesToString : just that
func BytesToString(data []byte) string {
	return string(data[:])
}

// CreateLowResGIF : Create .gif file from source content and .png palatte
func CreateLowResGIF(mp4 string, gif string) {
	fmt.Println("-- CREATE GIF -- ", gif)
	cmd := exec.Command("ffmpeg", "-ss", "0", "-t", "11", "-i", mp4, "-filter_complex", "[0:v] fps=24,scale=w=480:h=-1,split [a][b];[a] palettegen=stats_mode=single [p];[b][p] paletteuse=new=1", gif)
	cmd.Run()
}

// CreateGIF : Create .gif file from source content and .png palatte
func CreateGIF(mp4 string, gif string) {
	fmt.Println("-- CREATE GIF -- ", gif)
	cmd := exec.Command("ffmpeg", "-y", "-ss", "0", "-t", "11", "-i", mp4, "-filter_complex", "[0:v] fps=60,scale=w=480:h=-1,split [a][b];[a] palettegen=stats_mode=single [p];[b][p] paletteuse=new=1", gif)
	cmd.Run()
}

// MoveGIF : Move .gif to folder
func MoveGIF(source string, destination string) {
	fmt.Println("-- MOVE GIF --")
	cmd := exec.Command("mv", source, destination)
	cmd.Run()
}

// CreateDirectories : Create /gif directories
func CreateDirectories(dir string) {
	fmt.Println("-- DIRS CREATED --")
	cmd3 := exec.Command("mkdir", dir)
	cmd3.Run()
}

// CreateGifExtension : create .gif directory extension
func CreateGifExtension(folderCollect []string, files3 []string, testfiles []string) []string {
	fmt.Println("starting")
	var gifhold []string
	for k3 := 0; k3 < len(folderCollect); k3++ {

		for lo := 0; lo < len(files3); lo++ {

			noExtension := testfiles[lo][:len(testfiles[lo])-6]
			newDir := noExtension + "/gifs"
			fmt.Println(noExtension)
			gifhold = append(gifhold, newDir)
		}
	}
	fmt.Println("file moved to /gif folder")
	return gifhold
}

// ExportCSV : export CSV
func ExportCSV(inputdata [][]string, csvFile string) {
	filename := csvFile
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("FAILED TO CREATE CSV", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range inputdata {
		err := writer.Write(value)
		if err != nil {
			log.Fatal("FAILED TO WRITE CSV", err)
		}
		// fmt.Println(value)
	}
}

const (
	layoutISO = "2006-01-02"
	layoutUS  = "January 2, 2006"
)

// ProbeFiles : ffprobe input files
func ProbeFiles(dir string, files []string, folderdates []string) [][]string {
	var matrix [][]string
	matrix = append(matrix, []string{
		"Filename",
		"File Type",
		"Folder Date",
		"Folder Day Number",
		"Edit Date",
		"Edit Day",
		"Edit Day Number",
		"Month",
		"Year",
		"Time",
		"Timezone",
		"Duration",
		"Size",
		"Bitrate",
		"Format",
		"Formant Long"})

	// Probe Video Files
	for h := 0; h < len(files); h++ {
		new := dir + "/" + folderdates[h] + "/edits/" + files[h]
		// fmt.Println(new)
		data, err := ffprobe.GetProbeData(new, 5000*time.Millisecond)
		if err != nil {
			log.Panicf("Error getting data: %v", err)
		}

		buf, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			log.Panicf("Error unmarshalling: %v", err)
		}
		// log.Print(string(buf))

		var probed Ffprobe
		if err := json.Unmarshal(buf, &probed); err != nil {
			panic(err)
		}

		ffprobeFilename := probed.Format.Filename
		cleanName := filepath.Base(ffprobeFilename)

		unixdate := string(probed.Format.Tags.CreationTime.Format(time.RFC850))

		s := strings.Split(unixdate, ",")
		// date := s[1][1:11]
		day := s[0]
		dayNum := s[1][1:3]
		month := s[1][4:7]
		year := "20" + s[1][8:11]
		edittime := s[1][11:19]
		loc := s[1][20:23]
		folderDay := folderdates[h][3:5]

		// fmt.Println(date, day, dayNum, month, year, time, loc)

		fmt.Println(cleanName[:len(cleanName)-4], cleanName[len(cleanName)-4:], folderdates[h], month, folderDay, year, folderDay, day, dayNum, month, year, edittime, loc, probed.Format.Duration, probed.Format.Tags.CreationTime, probed.Format.FormatLongName, probed.Format.Size)

		fmt.Println("______________________________________")
		matrix = append(matrix, []string{
			cleanName[:len(cleanName)-4],
			cleanName[len(cleanName)-4:],
			folderdates[h],
			folderDay,
			// date,
			day,
			dayNum,
			month,
			year,
			edittime,
			loc,
			probed.Format.Duration,
			probed.Format.Size,
			probed.Format.BitRate,
			probed.Format.FormatName,
			probed.Format.FormatLongName})
	}
	// fmt.Println(matrix)
	return matrix
}

// ReadCSV : asdf
func ReadCSV(input string) []string {
	filename := input

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		panic(err)
	}

	var fromEditAnalysis []string
	// Loop through lines & turn into object
	for _, line := range lines {
		data := CsvLine{
			Column1: line[0],
		}
		fromEditAnalysis = append(fromEditAnalysis, data.Column1)
	}
	return fromEditAnalysis
}

func ReadVimeoCSV(input string) ([]string, []string, []string) {
	// filename := "Leafly-Strains - Sheet1.csv"
	filename := input

	// Open CSV file
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		panic(err)
	}

	var fromEditAnalysis []string
	var id []string
	var dLink []string
	// Loop through lines & turn into object
	for _, line := range lines {
		data := CsvLine{
			Column2: line[1],
			Column1: line[2],
			Column3: line[5],
		}
		fromEditAnalysis = append(fromEditAnalysis, data.Column1)
		id = append(id, data.Column2)
		dLink = append(dLink, data.Column3)
	}
	return id, fromEditAnalysis, dLink
}

// CrawlAndCollect : asdf
func CrawlAndCollect(searchdirectory string, searchType string) ([]string, []string) {
	var collector []string
	var folderCollect []string
	var files3 []string
	var testfiles []string

	searchDir := searchdirectory
	// fmt.Println(searchDir)

	fileList := []string{}
	filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return nil
	})

	for _, file := range fileList {
		collector = append(collector, file)
	}

	g := RemoveDuplicates(collector)

	// fmt.Println(searchdirectory[27:len(searchdirectory)])

	for i := 0; i < len(g); i++ {
		fileInfo, err = os.Stat(g[i])
		if err != nil {
			log.Fatal(err)
		}
		if fileInfo.IsDir() == true {
			if (fileInfo.Name() != "edits") && (fileInfo.Name() != "raw") && (fileInfo.Name() != "png") && (fileInfo.Name() != searchdirectory[27:len(searchdirectory)]) && (fileInfo.Name() != "gifs") && (fileInfo.Name() != "jpg") && (fileInfo.Name() != "cuts") && (fileInfo.Name() != "stills") {
				folderCollect = append(folderCollect, fileInfo.Name())
			}
		}
	}

	for k := 1; k < len(folderCollect); k++ {
		newDir := searchDir + "/" + folderCollect[k] + searchType
		fmt.Println(newDir)
		err := filepath.Walk(newDir, func(path string, f os.FileInfo, err error) error {

			if !f.IsDir() {
				if (f.Name()[len(f.Name())-4:]) == ".mp4" {
					// fmt.Println("MP4")
					r, err2 := regexp.MatchString(".mp4", f.Name())
					if err2 == nil && r {
						testfiles = append(testfiles, newDir)
						files3 = append(files3, newDir+"/"+f.Name())
						// fmt.Println(testfiles)
						// fmt.Println(f.Name())
					} else {

					}
				} else if f.Name()[len(f.Name())-4:] == ".mov" {
					// fmt.Println("MOV")
					r, err2 := regexp.MatchString(".mov", f.Name())
					if err2 == nil && r {
						testfiles = append(testfiles, newDir)
						files3 = append(files3, newDir+"/"+f.Name())
						// fmt.Println(f.Name())
					} else {

					}
				}
			}
			return nil
		})
		if err != nil {
			panic(err)
		}
	}
	// gifhold := CreateGifExtension(folderCollect, files3, testfiles)
	// fmt.Println(folderCollect, testfiles, files3)
	return testfiles, files3
}

// CrawlAndCollect : asdf
func CrawlAndCollectGIF(searchdirectory string, searchType string) ([]string, []string, []string) {
	var collector []string
	var folderCollect []string
	var files3 []string
	var testfiles []string

	searchDir := searchdirectory
	// fmt.Println(searchDir)

	fileList := []string{}
	filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return nil
	})

	for _, file := range fileList {
		collector = append(collector, file)
	}

	g := RemoveDuplicates(collector)

	// fmt.Println(searchdirectory[27:len(searchdirectory)])

	for i := 0; i < len(g); i++ {
		fileInfo, err = os.Stat(g[i])
		if err != nil {
			log.Fatal(err)
		}
		if fileInfo.IsDir() == true {
			if (fileInfo.Name() != "edits") && (fileInfo.Name() != "raw") && (fileInfo.Name() != searchdirectory[27:len(searchdirectory)]) && (fileInfo.Name() != "gifs") && (fileInfo.Name() != "png") && (fileInfo.Name() != "cuts") && (fileInfo.Name() != "stills") && (fileInfo.Name() != "jpg") {
				folderCollect = append(folderCollect, fileInfo.Name())
			}
		}
	}

	// create folders
	for k2 := 0; k2 < len(folderCollect); k2++ {
		newDir := searchDir + "/" + folderCollect[k2] + "/gifs"
		CreateDirectories(newDir)
	}

	for k3 := 0; k3 < len(folderCollect); k3++ {
		newDir := searchDir + "/" + folderCollect[k3] + "/png"
		CreateDirectories(newDir)
	}

	for k3 := 0; k3 < len(folderCollect); k3++ {
		newDir := searchDir + "/" + folderCollect[k3] + "/jpg"
		CreateDirectories(newDir)
	}

	// fmt.Println(folderCollect)

	for k := 1; k < len(folderCollect); k++ {
		newDir := searchDir + "/" + folderCollect[k] + searchType
		err := filepath.Walk(newDir, func(path string, f os.FileInfo, err error) error {
			fmt.Println(newDir)
			if !f.IsDir() {
				if (f.Name()[len(f.Name())-4:]) == ".mp4" {
					// fmt.Println("MP4")
					r, err2 := regexp.MatchString(".mp4", f.Name())
					if err2 == nil && r {
						testfiles = append(testfiles, newDir)
						files3 = append(files3, newDir+"/"+f.Name())
						// fmt.Println(testfiles)
						// fmt.Println(f.Name())
					} else {

					}
				} else if f.Name()[len(f.Name())-4:] == ".mov" {
					// fmt.Println("MOV")
					r, err2 := regexp.MatchString(".mov", f.Name())
					if err2 == nil && r {
						testfiles = append(testfiles, newDir)
						files3 = append(files3, newDir+"/"+f.Name())
						// fmt.Println(f.Name())
					} else {

					}
				}
			}
			return nil
		})
		if err != nil {
			panic(err)
		}
	}
	gifhold := CreateGifExtension(folderCollect, files3, testfiles)
	// fmt.Println(folderCollect, testfiles, files3)
	return testfiles, files3, gifhold
}
