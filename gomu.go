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
	"strconv"
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

type NlpTagCsv struct {
	Column1 string
	Column2 string
}

// ReadStrainNlpTagCsvSetup ... setup for the nlptagcsv
func ReadStrainNlpTagCsvSetup() [][]string {
	masterNLP := []string{"(", ")", ",", ":", ".", "''", "``", "#", "$", "CC", "CD", "DT", "EX", "FW", "IN", "JJ", "JJR", "JJS", "LS", "MD", "NN", "NNP", "NNPS", "NNS", "PDT", "POS", "PRP", "PRP$", "RB", "RBR", "RBS", "RP", "SYM", "TO", "UH", "VB", "VBD", "VBG", "VBN", "VBP", "VBZ", "WDT", "WP", "WP$", "WRB"}

	var getthemall [][]string
	getthemall = append(getthemall, []string{
		"Strain File",
		masterNLP[0],
		masterNLP[1],
		masterNLP[2],
		masterNLP[3],
		masterNLP[4],
		masterNLP[5],
		masterNLP[6],
		masterNLP[7],
		masterNLP[8],
		masterNLP[9],
		masterNLP[10],
		masterNLP[11],
		masterNLP[12],
		masterNLP[13],
		masterNLP[14],
		masterNLP[15],
		masterNLP[16],
		masterNLP[17],
		masterNLP[18],
		masterNLP[19],
		masterNLP[20],
		masterNLP[21],
		masterNLP[22],
		masterNLP[23],
		masterNLP[24],
		masterNLP[25],
		masterNLP[26],
		masterNLP[27],
		masterNLP[28],
		masterNLP[29],
		masterNLP[30],
		masterNLP[31],
		masterNLP[32],
		masterNLP[33],
		masterNLP[34],
		masterNLP[35],
		masterNLP[36],
		masterNLP[37],
		masterNLP[38],
		masterNLP[39],
		masterNLP[40],
		masterNLP[41],
		masterNLP[42],
		masterNLP[43],
		masterNLP[44]})
	return getthemall
}

// ReadStrainNlpTagCsv ... read individual csv file contain strain NLP tags
func ReadStrainNlpTagCsv(inputFile string) [][]string {

	// fmt.Println("gomu --- ", inputFile)
	incomingFile := inputFile[65:len(inputFile)]

	// Open CSV file
	f, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		panic(err)
	}

	var matrix [][]string
	// Loop through lines & turn into object
	for _, line := range lines {
		data := NlpTagCsv{
			Column1: line[0],
			Column2: line[1],
		}
		matrix = append(matrix, []string{data.Column1, data.Column2})
	}

	var nlpTag01, nlpTag02, nlpTag03, nlpTag04, nlpTag05, nlpTag06, nlpTag07, nlpTag08, nlpTag09, nlpTag10, nlpTag11, nlpTag12, nlpTag13, nlpTag14, nlpTag15, nlpTag16, nlpTag17, nlpTag18, nlpTag19, nlpTag20, nlpTag21, nlpTag22, nlpTag23, nlpTag24, nlpTag25, nlpTag26, nlpTag27, nlpTag28, nlpTag29, nlpTag30, nlpTag31, nlpTag32, nlpTag33, nlpTag34, nlpTag35, nlpTag36, nlpTag37, nlpTag38, nlpTag39, nlpTag40, nlpTag41, nlpTag42, nlpTag43, nlpTag44, nlpTag45 [][]string

	var getthemall [][]string
	// masterNLP := []string{"(", ")", ",", ":", ".", "''", "``", "#", "$", "CC", "CD", "DT", "EX", "FW", "IN", "JJ", "JJR", "JJS", "LS", "MD", "NN", "NNP", "NNPS", "NNS", "PDT", "POS", "PRP", "PRP$", "RB", "RBR", "RBS", "RP", "SYM", "TO", "UH", "VB", "VBD", "VBG", "VBN", "VBP", "VBZ", "WDT", "WP", "WP$", "WRB"}

	// getthemall = append(getthemall, []string{
	// 	"Strain File",
	// 	masterNLP[0],
	// 	masterNLP[1],
	// 	masterNLP[2],
	// 	masterNLP[3],
	// 	masterNLP[4],
	// 	masterNLP[5],
	// 	masterNLP[6],
	// 	masterNLP[7],
	// 	masterNLP[8],
	// 	masterNLP[9],
	// 	masterNLP[10],
	// 	masterNLP[11],
	// 	masterNLP[12],
	// 	masterNLP[13],
	// 	masterNLP[14],
	// 	masterNLP[15],
	// 	masterNLP[16],
	// 	masterNLP[17],
	// 	masterNLP[18],
	// 	masterNLP[19],
	// 	masterNLP[20],
	// 	masterNLP[21],
	// 	masterNLP[22],
	// 	masterNLP[23],
	// 	masterNLP[24],
	// 	masterNLP[25],
	// 	masterNLP[26],
	// 	masterNLP[27],
	// 	masterNLP[28],
	// 	masterNLP[29],
	// 	masterNLP[30],
	// 	masterNLP[31],
	// 	masterNLP[32],
	// 	masterNLP[33],
	// 	masterNLP[34],
	// 	masterNLP[35],
	// 	masterNLP[36],
	// 	masterNLP[37],
	// 	masterNLP[38],
	// 	masterNLP[39],
	// 	masterNLP[40],
	// 	masterNLP[41],
	// 	masterNLP[42],
	// 	masterNLP[43],
	// 	masterNLP[44]})

	i := 0
	for i < len(matrix)-1 {
		i++
		// collect = append(collect, []string{matrix[i][0], matrix[i][1]})
		switch matrix[i][1] {
		case "(":
			nlpTag01 = append(nlpTag01, []string{matrix[i][0], matrix[i][1]})
			// fmt.Println(nlpTag01)
		case ")":
			nlpTag02 = append(nlpTag02, []string{matrix[i][0], matrix[i][1]})
			// fmt.Println(nlpTag02)
		case ",":
			nlpTag03 = append(nlpTag03, []string{matrix[i][0], matrix[i][1]})
			// fmt.Println(nlpTag03)
		case ":":
			nlpTag04 = append(nlpTag04, []string{matrix[i][0], matrix[i][1]})
			// fmt.Println(nlpTag04)
		case ".":
			nlpTag05 = append(nlpTag05, []string{matrix[i][0], matrix[i][1]})
			// fmt.Println(nlpTag05)
		case "''":
			nlpTag06 = append(nlpTag06, []string{matrix[i][0], matrix[i][1]})
			// fmt.Println(nlpTag06)
		case "``":
			nlpTag07 = append(nlpTag07, []string{matrix[i][0], matrix[i][1]})
			// fmt.Println(nlpTag07)
		case "#":
			nlpTag08 = append(nlpTag08, []string{matrix[i][0], matrix[i][1]})
			// fmt.Println(nlpTag08)
		case "$":
			nlpTag09 = append(nlpTag09, []string{matrix[i][0], matrix[i][1]})
			// fmt.Println(nlpTag09)
		case "CC":
			nlpTag10 = append(nlpTag10, []string{matrix[i][0], matrix[i][1]})
			// fmt.Println(nlpTag10)
		case "CD":
			nlpTag11 = append(nlpTag11, []string{matrix[i][0], matrix[i][1]})
			// fmt.Println(nlpTag11)
		case "DT":
			nlpTag12 = append(nlpTag12, []string{matrix[i][0], matrix[i][1]})
			// fmt.Println(nlpTag12)
		case "EX":
			nlpTag13 = append(nlpTag13, []string{matrix[i][0], matrix[i][1]})
			// fmt.Println(nlpTag13)
		case "FW":
			nlpTag14 = append(nlpTag14, []string{matrix[i][0], matrix[i][1]})
			// fmt.Println(nlpTag14)
		case "IN":
			nlpTag15 = append(nlpTag15, []string{matrix[i][0], matrix[i][1]})
			// fmt.Println(nlpTag15)
		case "JJ":
			nlpTag16 = append(nlpTag16, []string{matrix[i][0], matrix[i][1]})
			// fmt.Println(nlpTag16)
		case "JJR":
			nlpTag17 = append(nlpTag17, []string{matrix[i][0], matrix[i][1]})
			// fmt.Println(nlpTag17)
		case "JJS":
			nlpTag18 = append(nlpTag18, []string{matrix[i][0], matrix[i][1]})
			// fmt.Println(nlpTag18)
		case "LS":
			nlpTag19 = append(nlpTag19, []string{matrix[i][0], matrix[i][1]})
			// fmt.Println(nlpTag19)
		case "MD":
			nlpTag20 = append(nlpTag20, []string{matrix[i][0], matrix[i][1]})
			// fmt.Println(nlpTag20)
		case "NN":
			nlpTag21 = append(nlpTag21, []string{matrix[i][0], matrix[i][1]})
			// fmt.Println(nlpTag21)
		case "NNP":
			nlpTag22 = append(nlpTag22, []string{matrix[i][0], matrix[i][1]})
			// fmt.Println(nlpTag22)
		case "NNPS":
			nlpTag23 = append(nlpTag23, []string{matrix[i][0], matrix[i][1]})
			// fmt.Println(nlpTag23)
		case "NNS":
			nlpTag24 = append(nlpTag24, []string{matrix[i][0], matrix[i][1]})
			// fmt.Println(nlpTag24)
		case "PDT":
			nlpTag25 = append(nlpTag25, []string{matrix[i][0], matrix[i][1]})
			// fmt.Println(nlpTag25)
		case "POS":
			nlpTag26 = append(nlpTag26, []string{matrix[i][0], matrix[i][1]})
			// fmt.Println(nlpTag26)
		case "PRP":
			nlpTag27 = append(nlpTag27, []string{matrix[i][0], matrix[i][1]})
			// fmt.Println(nlpTag27)
		case "PRP$":
			nlpTag28 = append(nlpTag28, []string{matrix[i][0], matrix[i][1]})
			// fmt.Println(nlpTag28)
		case "RB":
			nlpTag29 = append(nlpTag29, []string{matrix[i][0], matrix[i][1]})
			// fmt.Println(nlpTag29)
		case "RBR":
			nlpTag30 = append(nlpTag30, []string{matrix[i][0], matrix[i][1]})
			// fmt.Println(nlpTag30)
		case "RBS":
			nlpTag31 = append(nlpTag31, []string{matrix[i][0], matrix[i][1]})
			// fmt.Println(nlpTag31)
		case "RP":
			nlpTag32 = append(nlpTag32, []string{matrix[i][0], matrix[i][1]})
			// fmt.Println(nlpTag32)
		case "SYM":
			nlpTag33 = append(nlpTag33, []string{matrix[i][0], matrix[i][1]})
			// fmt.Println(nlpTag33)
		case "TO":
			nlpTag34 = append(nlpTag34, []string{matrix[i][0], matrix[i][1]})
			// fmt.Println(nlpTag34)
		case "UH":
			nlpTag35 = append(nlpTag35, []string{matrix[i][0], matrix[i][1]})
			// fmt.Println(nlpTag35)
		case "VB":
			nlpTag36 = append(nlpTag36, []string{matrix[i][0], matrix[i][1]})
			// fmt.Println(nlpTag36)
		case "VBD":
			nlpTag37 = append(nlpTag37, []string{matrix[i][0], matrix[i][1]})
			// fmt.Println(nlpTag37)
		case "VBG":
			nlpTag38 = append(nlpTag38, []string{matrix[i][0], matrix[i][1]})
			// fmt.Println(nlpTag38)
		case "VBN":
			nlpTag39 = append(nlpTag39, []string{matrix[i][0], matrix[i][1]})
			// fmt.Println(nlpTag39)
		case "VBP":
			nlpTag40 = append(nlpTag40, []string{matrix[i][0], matrix[i][1]})
			// fmt.Println(nlpTag40)
		case "VBZ":
			nlpTag41 = append(nlpTag41, []string{matrix[i][0], matrix[i][1]})
			// fmt.Println(nlpTag41)
		case "WDT":
			nlpTag42 = append(nlpTag42, []string{matrix[i][0], matrix[i][1]})
			// fmt.Println(nlpTag42)
		case "WP":
			nlpTag43 = append(nlpTag43, []string{matrix[i][0], matrix[i][1]})
			// fmt.Println(nlpTag43)
		case "WP$":
			nlpTag44 = append(nlpTag44, []string{matrix[i][0], matrix[i][1]})
			// fmt.Println(nlpTag44)
		case "WRB":
			nlpTag45 = append(nlpTag45, []string{matrix[i][0], matrix[i][1]})
			// fmt.Println(nlpTag45)
		}
	}

	getthemall = append(getthemall, []string{
		incomingFile,
		strconv.Itoa(len(nlpTag01)),
		strconv.Itoa(len(nlpTag02)),
		strconv.Itoa(len(nlpTag03)),
		strconv.Itoa(len(nlpTag04)),
		strconv.Itoa(len(nlpTag05)),
		strconv.Itoa(len(nlpTag06)),
		strconv.Itoa(len(nlpTag07)),
		strconv.Itoa(len(nlpTag08)),
		strconv.Itoa(len(nlpTag09)),
		strconv.Itoa(len(nlpTag10)),
		strconv.Itoa(len(nlpTag11)),
		strconv.Itoa(len(nlpTag12)),
		strconv.Itoa(len(nlpTag13)),
		strconv.Itoa(len(nlpTag14)),
		strconv.Itoa(len(nlpTag15)),
		strconv.Itoa(len(nlpTag16)),
		strconv.Itoa(len(nlpTag17)),
		strconv.Itoa(len(nlpTag18)),
		strconv.Itoa(len(nlpTag19)),
		strconv.Itoa(len(nlpTag20)),
		strconv.Itoa(len(nlpTag21)),
		strconv.Itoa(len(nlpTag22)),
		strconv.Itoa(len(nlpTag23)),
		strconv.Itoa(len(nlpTag24)),
		strconv.Itoa(len(nlpTag25)),
		strconv.Itoa(len(nlpTag26)),
		strconv.Itoa(len(nlpTag27)),
		strconv.Itoa(len(nlpTag28)),
		strconv.Itoa(len(nlpTag29)),
		strconv.Itoa(len(nlpTag30)),
		strconv.Itoa(len(nlpTag31)),
		strconv.Itoa(len(nlpTag32)),
		strconv.Itoa(len(nlpTag33)),
		strconv.Itoa(len(nlpTag34)),
		strconv.Itoa(len(nlpTag35)),
		strconv.Itoa(len(nlpTag36)),
		strconv.Itoa(len(nlpTag37)),
		strconv.Itoa(len(nlpTag38)),
		strconv.Itoa(len(nlpTag39)),
		strconv.Itoa(len(nlpTag40)),
		strconv.Itoa(len(nlpTag41)),
		strconv.Itoa(len(nlpTag42)),
		strconv.Itoa(len(nlpTag43)),
		strconv.Itoa(len(nlpTag44)),
		strconv.Itoa(len(nlpTag45))})
	// fmt.Println(getthemall)
	return getthemall
}

// WalkDirectory: walk directory to find a set of files within said directory
func WalkDirectory(dir string) []string {
	var collector []string
	fileList := []string{}
	filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return nil
	})

	for _, file := range fileList {
		collector = append(collector, file)
	}
	// fmt.Println(collector)
	return collector
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

func ReadDirRmDups(dir string) []string {
	var collectFiles []string

	// for looper := 0; looper < len(dir); looper++ {
	fileList := []string{}
	filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return nil
	})

	for _, file := range fileList {
		collectFiles = append(collectFiles, file)
	}

	rmDups := RemoveDuplicates(collectFiles)

	return rmDups
}

func SearchForThem(typer string, rmDups []string) []string {
	var savePng []string
	switch typer {
	case "gifs":
		fmt.Println("find the gifs")
		i := 0
		for i < len(rmDups)-1 {
			i++
			h := strings.Contains(rmDups[i], typer+"/")
			switch h {
			case true:
				if !strings.Contains(rmDups[i], ".DS_Store") {
					savePng = append(savePng, rmDups[i])
				}
			}
		}
	case "jpg":
		fmt.Println("find the jpg")
		i := 0
		for i < len(rmDups)-1 {
			i++
			h := strings.Contains(rmDups[i], typer+"/")
			switch h {
			case true:
				if !strings.Contains(rmDups[i], ".DS_Store") {
					savePng = append(savePng, rmDups[i])
				}
			}
		}
	case "edits":
		fmt.Println("find the mp4s")
		i := 0
		for i < len(rmDups)-1 {
			i++
			h := strings.Contains(rmDups[i], typer+"/")
			switch h {
			case true:
				if !strings.Contains(rmDups[i], ".DS_Store") {
					savePng = append(savePng, rmDups[i])
				}
			}
		}
	case "png":
		fmt.Println("find pngs")
		i := 0
		for i < len(rmDups)-1 {
			i++
			h := strings.Contains(rmDups[i], typer+"/")
			switch h {
			case true:
				if !strings.Contains(rmDups[i], ".DS_Store") {
					savePng = append(savePng, rmDups[i])
				}
			}
		}
	case "raw":
		fmt.Println("find raw")
		i := 0
		for i < len(rmDups)-1 {
			i++
			h := strings.Contains(rmDups[i], typer+"/")
			switch h {
			case true:
				if !strings.Contains(rmDups[i], ".DS_Store") {
					savePng = append(savePng, rmDups[i])
				}
			}
		}
	}
	return savePng
}

// Acquire files
func CreateStillBundle2(files string) {
	striped := files[:len(files)-4]
	newMp4 := striped + ".mp4"
	newJpg := striped + ".jpg"
	jpgDir := striped[:50] + "/jpg"

	editName := strings.SplitAfter(striped, "/edits")

	fmt.Println(newMp4, newJpg, jpgDir, editName[1][1:])

	// jpgDir2 := striped[:50] + "/jpg/" + editName[1][1:]
	// fmt.Println(jpgDir2)
	// cmd3 := exec.Command("mkdir", jpgDir2)
	// cmd3.Run()

	check := jpgDir + "/" + editName[1][1:] + "/" + editName[1][1:] + "-frame-%04d.jpg"
	fmt.Println(check)
	cmd5 := exec.Command("ffmpeg", "-i", newMp4, check)

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
	fmt.Println("source, destination : ", source, destination)
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
	// matrix = append(matrix, []string{
	// 	"Filename",
	// 	"File Type",
	// 	"Folder Date",
	// 	"Folder Day Number",
	// 	"Edit Date",
	// 	"Edit Day Number",
	// 	"Edit Day",
	// 	"Month",
	// 	"Year",
	// 	"Time",
	// 	"Timezone",
	// 	"Duration",
	// 	"Size",
	// 	"Bitrate",
	// 	"Format",
	// 	"Formant Long"})

	matrix = append(matrix, []string{
		"Filename",
		"FolderDate",
		"Folder Month",
		"Folder Day",
		"Folder Year",
		"Edit Date",
		"Edit Month",
		"Edit Day",
		"Edit Year",
		"Edit Day Number",
		"Timestamp",
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
			log.Panicf("Error getting data:  %v", err)
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

		// ffprobeFilename := probed.Format.Filename
		// cleanName := filepath.Base(ffprobeFilename)

		// unixdate := string(probed.Format.Tags.CreationTime.Format(time.RFC850))

		// s := strings.Split(unixdate, ",")
		// // date := s[1][1:11]
		// day := s[0]
		// dayNum := s[1][1:3]
		// month := s[1][4:7]
		// year := "20" + s[1][8:11]
		// edittime := s[1][11:19]
		// loc := s[1][20:23]
		// folderDay := folderdates[h][3:5]

		// // fmt.Println(date, day, dayNum, month, year, time, loc)

		// fmt.Println(cleanName[:len(cleanName)-4], cleanName[len(cleanName)-4:], folderdates[h], month, folderDay, year, folderDay, day, dayNum, month, year, edittime, loc, probed.Format.Duration, probed.Format.Tags.CreationTime, probed.Format.FormatLongName, probed.Format.Size)

		ffprobeFilename := probed.Format.Filename
		cleanFilename := filepath.Base(ffprobeFilename)

		fmt.Println(cleanFilename)
		fmt.Println(probed.Format.Duration)
		fmt.Println(probed.Format.Tags.CreationTime)
		fmt.Println(probed.Format.FormatLongName)
		fmt.Println(probed.Format.Size)

		unixdate := string(probed.Format.Tags.CreationTime.Format(time.RFC850))

		s := strings.Split(unixdate, ",")
		folderMonth := folderdates[h][:2]
		folderDay := folderdates[h][3:5]
		folderYear := folderdates[h][len(folderdates[h])-4:]
		editDate := s[1][1:11]
		editMonth := s[1][4:7]
		editDay := s[0]
		editYear := "20" + s[1][8:11]
		editDayNumber := s[1][1:3]
		timestamp := s[1][11:19]
		loc := s[1][20:23]

		fmt.Println("______________________________________")
		matrix = append(matrix, []string{
			// cleanFilename[:len(cleanFilename)-4],
			cleanFilename,
			folderdates[h],
			folderMonth,
			folderDay,
			folderYear,
			editMonth,
			editDate,
			editDay,
			editYear,
			editDayNumber,
			timestamp,
			loc,
			probed.Format.Duration,
			probed.Format.Size,
			probed.Format.BitRate,
			probed.Format.FormatName,
			probed.Format.FormatLongName})
		// matrix = append(matrix, []string{
		// 	cleanName[:len(cleanName)-4],
		// 	// cleanName[len(cleanName)-4:],
		// 	folderdates[h],
		// 	folderDay,
		// 	// date,
		// 	day,
		// 	dayNum,
		// 	month,
		// 	year,
		// 	edittime,
		// 	loc,
		// 	probed.Format.Duration,
		// 	probed.Format.Size,
		// 	probed.Format.BitRate,
		// 	probed.Format.FormatName,
		// 	probed.Format.FormatLongName})
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

func PrettyPrinter(strains ...[]string) string {
	// var output []string
	b, err := json.MarshalIndent(strains, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	return (string(b))
}

func ExportImageMagickHeaders() [][]string {
	var matrix [][]string
	matrix = append(matrix, []string{
		"Image",
		"Format",
		"Mime type",
		"Class",
		"Geometry",
		"Resolution",
		"Print size",
		"Units",
		"Colorspace",
		"Type",
		"Base Type",
		"Endianess",
		"Depth",
		"Red",
		"Green",
		"Blue",
		"Alpha",
		"Pixels",
		"red min",
		"red max",
		"red mean",
		"red standard deviation",
		"red kurtosis",
		"red skewness",
		"red entropy",
		"green min",
		"green max",
		"green mean",
		"green standard deviation",
		"green kurtosis",
		"green skewness",
		"green entropy",
		"blue min",
		"blue max",
		"blue mean",
		"blue standard deviation",
		"blue kurtosis",
		"blue skewness",
		"blue entropy",
		"alpha min",
		"alpha max",
		"alpha mean",
		"alpha standard deviation",
		"alpha kurtosis",
		"alpha skewness",
		"alpha entropy",
		"imagstats min",
		"imagstats max",
		"imagstats mean",
		"imagstats standard deviation",
		"imagstats kurtosis",
		"imagstats skewness",
		"imagstats entropy",
		"Colors",
		"Rendering intent",
		"Gamma",
		"chromaticity red primary",
		"chromaticity green primary",
		"chromaticity blue primary",
		"chromaticity white point",
		"Matte color",
		"Background color",
		"Border color",
		"Transparent color",
		"Interlace",
		"Intensity",
		"Compose",
		"Page geometry",
		"Dispose",
		"Iterations",
		"Compression",
		"Orientation",
		"Prop date create",
		"Prop date modify",
		"png:IHDR.bit-depth-orig",
		"png:IHDR.bit_depth",
		"png:IHDR.color-type-orig",
		"png:IHDR.color_type",
		"png:IHDR.interlace_method",
		"png:IHDR.width,height",
		"png:pHYs",
		"png:sRGB",
		"Prop signature",
		"Artifacts verbose",
		"Tainted",
		"Filesize",
		"Number pixels",
		"Pixels per second",
		"User time",
		"Elapsed time",
		"Version",
	})

	return matrix
}

// EditCsv : struct for edit analysis csv
type EditCsv struct {
	Column1  string
	Column2  string
	Column3  string
	Column4  string
	Column5  string
	Column6  string
	Column7  string
	Column8  string
	Column9  string
	Column10 string
	Column11 string
	Column12 string
	Column13 string
	Column14 string
	Column15 string
	Column16 string
	Column17 string
}

func ReadEditCsv(input string) [][]string {
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
	// fmt.Println(len(lines))
	var fromEditAnalysis [][]string

	// Loop through lines & turn into object
	for _, line := range lines {

		data := EditCsv{
			Column1:  line[0],
			Column2:  line[1],
			Column3:  line[2],
			Column4:  line[3],
			Column5:  line[4],
			Column6:  line[5],
			Column7:  line[6],
			Column8:  line[7],
			Column9:  line[8],
			Column10: line[9],
			Column11: line[10],
			Column12: line[11],
			Column13: line[12],
			Column14: line[13],
			Column15: line[14],
			Column16: line[15],
			Column17: line[16],
		}
		fromEditAnalysis = append(fromEditAnalysis, []string{
			data.Column1,
			data.Column2,
			data.Column3,
			data.Column4,
			data.Column5,
			data.Column6,
			data.Column7,
			data.Column8,
			data.Column9,
			data.Column10,
			data.Column11,
			data.Column12,
			data.Column13,
			data.Column14,
			data.Column15,
			data.Column16,
			data.Column17,
		})
	}

	return fromEditAnalysis
}

type MagickCsv struct {
	Column1  string
	Column2  string
	Column3  string
	Column4  string
	Column5  string
	Column6  string
	Column7  string
	Column8  string
	Column9  string
	Column10 string
	Column11 string
	Column12 string
	Column13 string
	Column14 string
	Column15 string
	Column16 string
	Column17 string
	Column18 string
	Column19 string
	Column20 string
	Column21 string
	Column22 string
	Column23 string
	Column24 string
	Column25 string
	Column26 string
	Column27 string
	Column28 string
	Column29 string
	Column30 string
	Column31 string
	Column32 string
	Column33 string
	Column34 string
	Column35 string
	Column36 string
	Column37 string
	Column38 string
	Column39 string
	Column40 string
	Column41 string
	Column42 string
	Column43 string
	Column44 string
	Column45 string
	Column46 string
	Column47 string
	Column48 string
	Column49 string
	Column50 string
	Column51 string
	Column52 string
	Column53 string
	Column54 string
	Column55 string
	Column56 string
	Column57 string
	Column58 string
	Column59 string
	Column60 string
	Column61 string
	Column62 string
	Column63 string
	Column64 string
	Column65 string
	Column66 string
	Column67 string
	Column68 string
	Column69 string
	Column70 string
	Column71 string
	Column72 string
	Column73 string
	Column74 string
	Column75 string
	Column76 string
	Column77 string
	Column78 string
	Column79 string
	Column80 string
	Column81 string
	Column82 string
	Column83 string
	Column84 string
	Column85 string
	Column86 string
	Column87 string
	Column88 string
	Column89 string
	Column90 string
	Column91 string
}

// ReadCSV : asdf
func ReadMagickCSV(input string) [][]string {
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

	var fromMagickAnalysis [][]string

	// Loop through lines & turn into object
	for _, line := range lines {
		data := MagickCsv{
			Column1:  line[0],
			Column2:  line[1],
			Column3:  line[2],
			Column4:  line[3],
			Column5:  line[4],
			Column6:  line[5],
			Column7:  line[6],
			Column8:  line[7],
			Column9:  line[8],
			Column10: line[9],
			Column11: line[10],
			Column12: line[11],
			Column13: line[12],
			Column14: line[13],
			Column15: line[14],
			Column16: line[15],
			Column17: line[16],
			Column18: line[17],
			Column19: line[18],
			Column20: line[19],
			Column21: line[20],
			Column22: line[21],
			Column23: line[22],
			Column24: line[23],
			Column25: line[24],
			Column26: line[25],
			Column27: line[26],
			Column28: line[27],
			Column29: line[28],
			Column30: line[29],
			Column31: line[30],
			Column32: line[31],
			Column33: line[32],
			Column34: line[33],
			Column35: line[34],
			Column36: line[35],
			Column37: line[36],
			Column38: line[37],
			Column39: line[38],
			Column40: line[39],
			Column41: line[40],
			Column42: line[41],
			Column43: line[42],
			Column44: line[43],
			Column45: line[44],
			Column46: line[45],
			Column47: line[46],
			Column48: line[47],
			Column49: line[48],
			Column50: line[49],
			Column51: line[50],
			Column52: line[51],
			Column53: line[52],
			Column54: line[53],
			Column55: line[54],
			Column56: line[55],
			Column57: line[56],
			Column58: line[57],
			Column59: line[58],
			Column60: line[59],
			Column61: line[60],
			Column62: line[61],
			Column63: line[62],
			Column64: line[63],
			Column65: line[64],
			Column66: line[65],
			Column67: line[66],
			Column68: line[67],
			Column69: line[68],
			Column70: line[69],
			Column71: line[70],
			Column72: line[71],
			Column73: line[72],
			Column74: line[73],
			Column75: line[74],
			Column76: line[75],
			Column77: line[76],
			Column78: line[77],
			Column79: line[78],
			Column80: line[79],
			Column81: line[80],
			Column82: line[81],
			Column83: line[82],
			Column84: line[83],
			Column85: line[84],
			Column86: line[85],
			Column87: line[86],
			Column88: line[87],
			Column89: line[88],
			Column90: line[89],
			Column91: line[90],
		}
		fromMagickAnalysis = append(fromMagickAnalysis, []string{data.Column1,
			data.Column2,
			data.Column3,
			data.Column4,
			data.Column5,
			data.Column6,
			data.Column7,
			data.Column8,
			data.Column9,
			data.Column10,
			data.Column11,
			data.Column12,
			data.Column13,
			data.Column14,
			data.Column15,
			data.Column16,
			data.Column17,
			data.Column18,
			data.Column19,
			data.Column20,
			data.Column21,
			data.Column22,
			data.Column23,
			data.Column24,
			data.Column25,
			data.Column26,
			data.Column27,
			data.Column28,
			data.Column29,
			data.Column30,
			data.Column31,
			data.Column32,
			data.Column33,
			data.Column34,
			data.Column35,
			data.Column36,
			data.Column37,
			data.Column38,
			data.Column39,
			data.Column40,
			data.Column41,
			data.Column42,
			data.Column43,
			data.Column44,
			data.Column45,
			data.Column46,
			data.Column47,
			data.Column48,
			data.Column49,
			data.Column50,
			data.Column51,
			data.Column52,
			data.Column53,
			data.Column54,
			data.Column55,
			data.Column56,
			data.Column57,
			data.Column58,
			data.Column59,
			data.Column60,
			data.Column61,
			data.Column62,
			data.Column63,
			data.Column64,
			data.Column65,
			data.Column66,
			data.Column67,
			data.Column68,
			data.Column69,
			data.Column70,
			data.Column71,
			data.Column72,
			data.Column73,
			data.Column74,
			data.Column75,
			data.Column76,
			data.Column77,
			data.Column78,
			data.Column79,
			data.Column80,
			data.Column81,
			data.Column82,
			data.Column83,
			data.Column84,
			data.Column85,
			data.Column86,
			data.Column87,
			data.Column88,
			data.Column89,
			data.Column90,
			data.Column91})
	}
	return fromMagickAnalysis
}

func CreateImageMagickAnalysis(inputFile string) [][]string {
	out, err := exec.Command("magick", "identify", "-verbose", inputFile).Output()

	if err != nil {
		fmt.Printf("%s", err)
	}

	// fmt.Println("Command Successfully Executed")

	output := string(out[:])
	new := strings.SplitAfter(output, "\n")

	var input []string
	for h := 0; h < len(new); h++ {
		input = append(input, new[h])
	}

	var matrix, what [][]string

	for x := 0; x < len(input)-1; x++ {
		if x <= 1 {
			splitter := strings.SplitAfter(input[x][:len(input[x])-1], ": ")
			matrix = append(matrix, []string{splitter[1]})
			// fmt.Println(x, splitter[0], splitter[1])
		} else if (x > 1) && (x < 13) {
			splitter := strings.SplitAfter(input[x][2:len(input[x])-1], ": ")
			matrix = append(matrix, []string{splitter[1]})
			// fmt.Println(x, splitter[0], splitter[1])
		} else if (x > 13) && (x < 18) {
			splitter := strings.SplitAfter(input[x][2:len(input[x])-1], ": ")
			matrix = append(matrix, []string{splitter[1]})
			// fmt.Println(x, splitter[0], splitter[1])
		} else if (x > 18) && (x < 20) {
			// fmt.Println("__", input[x][4:len(input[x])-1])
			splitter := strings.SplitAfter(input[x][4:len(input[x])-1], ": ")
			matrix = append(matrix, []string{splitter[1]})
		} else if (x > 20) && (x < 28) {
			// fmt.Println("__", input[x][6:len(input[x])-1])
			splitter := strings.SplitAfter(input[x][4:len(input[x])-1], ": ")
			matrix = append(matrix, []string{splitter[1]})
		} else if (x > 28) && (x < 36) { // Red:
			// fmt.Println("__red ", input[x][6:len(input[x])-1])
			splitter := strings.SplitAfter(input[x][4:len(input[x])-1], ": ")
			matrix = append(matrix, []string{splitter[1]})
		} else if (x > 36) && (x < 44) { // Green:
			// fmt.Println("__green ", input[x][6:len(input[x])-1])
			splitter := strings.SplitAfter(input[x][6:len(input[x])-1], ": ")
			matrix = append(matrix, []string{splitter[1]})
		} else if (x > 44) && (x < 52) { // Blue:
			// fmt.Println("__blue ", input[x][6:len(input[x])-1])
			splitter := strings.SplitAfter(input[x][6:len(input[x])-1], ": ")
			matrix = append(matrix, []string{splitter[1]})
		} else if (x > 53) && (x < 61) { // Alpha:
			// fmt.Println("__alpha ", input[x][6:len(input[x])-1])
			splitter := strings.SplitAfter(input[x][6:len(input[x])-1], ": ")
			matrix = append(matrix, []string{splitter[1]})
		} else if (x > 61) && (x < 63) { // colors
			splitter := strings.SplitAfter(input[x][2:len(input[x])-1], ": ")
			// fmt.Println("__color", splitter[1])
			matrix = append(matrix, []string{splitter[1]})
		} else if (x > len(input)-42) && (x < len(input)-39) {
			splitter := strings.SplitAfter(input[x][2:len(input[x])-1], ": ")
			// fmt.Println("__rendering intent", splitter[1])
			matrix = append(matrix, []string{splitter[1]})
		} else if (x > len(input)-39) && (x < len(input)-33) {
			splitter := strings.SplitAfter(input[x][4:len(input[x])-1], ": ")
			// fmt.Println(splitter[1])
			matrix = append(matrix, []string{splitter[1]})
		} else if (x > len(input)-34) && (x < len(input)-22) {
			splitter := strings.SplitAfter(input[x][:len(input[x])-1], ": ")
			// fmt.Println(splitter[1])
			matrix = append(matrix, []string{splitter[1]})
		} else if (x > len(input)-22) && (x < len(input)-10) {
			splitter := strings.SplitAfter(input[x][:len(input[x])-1], ": ")
			// fmt.Println(splitter[1])
			matrix = append(matrix, []string{splitter[1]})
		} else if (x > len(input)-10) && (x < len(input)) {
			splitter := strings.SplitAfter(input[x][:len(input[x])-1], ": ")
			// fmt.Println(splitter[1])
			matrix = append(matrix, []string{splitter[1]})
		}
	}

	what = append(what, []string{
		matrix[0][0],
		matrix[1][0],
		matrix[2][0],
		matrix[3][0],
		matrix[4][0],
		matrix[5][0],
		matrix[6][0],
		matrix[7][0],
		matrix[8][0],
		matrix[9][0],
		matrix[10][0],
		matrix[11][0],
		matrix[12][0],
		matrix[13][0],
		matrix[14][0],
		matrix[15][0],
		matrix[16][0],
		matrix[17][0],
		matrix[18][0],
		matrix[19][0],
		matrix[20][0],
		matrix[21][0],
		matrix[22][0],
		matrix[23][0],
		matrix[24][0],
		matrix[25][0],
		matrix[26][0],
		matrix[27][0],
		matrix[28][0],
		matrix[29][0],
		matrix[30][0],
		matrix[31][0],
		matrix[32][0],
		matrix[33][0],
		matrix[34][0],
		matrix[35][0],
		matrix[36][0],
		matrix[37][0],
		matrix[38][0],
		matrix[39][0],
		matrix[40][0],
		matrix[41][0],
		matrix[42][0],
		matrix[43][0],
		matrix[44][0],
		matrix[45][0],
		matrix[46][0],
		matrix[47][0],
		matrix[48][0],
		matrix[49][0],
		matrix[50][0],
		matrix[51][0],
		matrix[52][0],
		matrix[53][0],
		matrix[54][0],
		matrix[55][0],
		matrix[56][0],
		matrix[57][0],
		matrix[58][0],
		matrix[59][0],
		matrix[60][0],
		matrix[61][0],
		matrix[62][0],
		matrix[63][0],
		matrix[64][0],
		matrix[65][0],
		matrix[66][0],
		matrix[67][0],
		matrix[68][0],
		matrix[69][0],
		matrix[70][0],
		matrix[71][0],
		matrix[72][0],
		matrix[73][0],
		matrix[74][0],
		matrix[75][0],
		matrix[76][0],
		matrix[77][0],
		matrix[78][0],
		matrix[79][0],
		matrix[80][0],
		matrix[81][0],
		matrix[82][0],
		matrix[83][0],
		matrix[84][0],
		matrix[85][0],
		matrix[86][0],
		matrix[87][0],
		matrix[88][0],
		matrix[89][0],
		matrix[90][0],
	})

	return what
}
