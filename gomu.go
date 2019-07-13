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
