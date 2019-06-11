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
	// cmd2 := exec.Command("ffmpeg", "-y", "-ss", "0", "-t", "13", "-i", mp4, "-filter_complex", "[0:v] palettegen", png)

	cmd2 := exec.Command("ffmpeg", "-y", "-ss", "0", "-t", "11", "-i", mp4, "-filter_complex", "[0:v] palettegen", png)

	cmd2.Run()
}

// CreateGIF : Create .gif file from source content and .png palatte
func CreateGIF(mp4 string, gif string) {
	fmt.Println("-- CREATE GIF -- ", gif)
	// cmd3 := exec.Command("ffmpeg", "-y", "-ss", "0", "-t", "13", "-i", mp4, "-filter_complex", "[0:v] fps=15,scale=w=480:h=-1,split [a][b];[a] palettegen=stats_mode=single [p];[b][p] paletteuse=new=1", gif)
	// cmd3 := exec.Command("ffmpeg", "-y", "-ss", "0", "-t", "11", "-i", mp4, "-filter_complex", "[0:v] fps=15,scale=w=1280:h=-1,split [a][b];[a] palettegen=stats_mode=single [p];[b][p] paletteuse=new=1", gif)
	cmd3 := exec.Command("ffmpeg", "-y", "-ss", "0", "-t", "11", "-i", mp4, "-filter_complex", "[0:v] fps=60,scale=w=480:h=-1,split [a][b];[a] palettegen=stats_mode=single [p];[b][p] paletteuse=new=1", gif)

	cmd3.Run()
}

// MoveGIF : Move .gif to folder
func MoveGIF(source string, destination string) {
	fmt.Println("-- MOVE GIF --")
	cmd4 := exec.Command("mv", source, destination)
	cmd4.Run()
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

// ProbeFiles : ffprobe input files
func ProbeFiles(dir string, files []string, folderdates []string) [][]string {
	var matrix [][]string
	matrix = append(matrix, []string{
		"Filename",
		"FolderDate",
		"Edit Date",
		"Edit Day",
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

		data, err := ffprobe.GetProbeData(new, 5000*time.Millisecond)
		if err != nil {
			log.Panicf("Error getting data: %v", err)
		}

		buf, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			log.Panicf("Error unmarshalling: %v", err)
		}
		log.Print(string(buf))

		var probed Ffprobe
		if err := json.Unmarshal(buf, &probed); err != nil {
			panic(err)
		}

		ffprobeFilename := probed.Format.Filename
		cleanFilename := filepath.Base(ffprobeFilename)

		fmt.Println(cleanFilename)
		fmt.Println(probed.Format.Duration)
		fmt.Println(probed.Format.Tags.CreationTime)
		fmt.Println(probed.Format.FormatLongName)
		fmt.Println(probed.Format.Size)

		unixdate := string(probed.Format.Tags.CreationTime.Format(time.RFC850))

		s := strings.Split(unixdate, ",")
		day := s[0]
		date := s[1][1:11]
		time := s[1][11:19]
		loc := s[1][20:23]

		fmt.Println("______________________________________")
		matrix = append(matrix, []string{
			cleanFilename,
			folderdates[h],
			date,
			day,
			time,
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

// CrawlAndCollect : asdf
func CrawlAndCollect(searchdirectory string, searchType string) ([]string, []string) {
	var collector []string
	var folderCollect []string
	var files3 []string
	var testfiles []string

	searchDir := searchdirectory

	fileList := []string{}
	filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return nil
	})

	for _, file := range fileList {
		collector = append(collector, file)
	}

	g := RemoveDuplicates(collector)

	for i := 0; i < len(g); i++ {
		fileInfo, err = os.Stat(g[i])
		if err != nil {
			log.Fatal(err)
		}
		if fileInfo.IsDir() == true {
			if (fileInfo.Name() != "edits") && (fileInfo.Name() != "raw") && (fileInfo.Name() != searchdirectory[22:len(searchdirectory)]) && (fileInfo.Name() != "gifs") && (fileInfo.Name() != "cuts") && (fileInfo.Name() != "stills") {
				folderCollect = append(folderCollect, fileInfo.Name())

			}
		}
	}

	for k := 1; k < len(folderCollect); k++ {
		newDir := searchDir + "/" + folderCollect[k] + searchType
		err := filepath.Walk(newDir, func(path string, f os.FileInfo, err error) error {

			if !f.IsDir() {
				if (f.Name()[len(f.Name())-4:]) == ".mp4" {
					fmt.Println("MP4")
					r, err2 := regexp.MatchString(".mp4", f.Name())
					if err2 == nil && r {
						testfiles = append(testfiles, newDir)
						files3 = append(files3, newDir+"/"+f.Name())
						fmt.Println(f.Name())
					} else {

					}
				} else if f.Name()[len(f.Name())-4:] == ".mov" {
					fmt.Println("MOV")
					r, err2 := regexp.MatchString(".mov", f.Name())
					if err2 == nil && r {
						testfiles = append(testfiles, newDir)
						files3 = append(files3, newDir+"/"+f.Name())
						fmt.Println(f.Name())
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
	return testfiles, files3
}
