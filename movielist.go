package movieconverter

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// MovieList ... 動画ファイルリスト
type MovieList struct {
	MovieFiles []MovieFile
}

// MovieFile ... 動画ファイル
type MovieFile struct {
	MovieFileName    string
	MovieUpdDatetime string
}

// ReadPrevious ... 前回監視時の動画ファイル情報リストを取得
func ReadPrevious(dirPath string) (*MovieList, error) {
	file, err := os.OpenFile(filepath.Join(dirPath, "prev.list"), os.O_RDONLY, 0)
	if err != nil {
		return nil, nil
	}
	defer file.Close()

	var movieFiles []MovieFile
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if isInvalidLine(line) {
			log.Printf("[WARN]前回監視時の動画ファイル情報リストに不正な行が含まれています。この行はスキップします。 [%s]\n", line)
			continue
		}
		movieFiles = append(
			movieFiles,
			createMovieFile(line))
	}

	return &MovieList{MovieFiles: movieFiles}, nil
}

func isInvalidLine(line string) bool {
	lineSeps := strings.Split(line, ",")
	if len(lineSeps) != 2 {
		return true
	}
	if lineSeps[0] == "" || lineSeps[1] == "" {
		return true
	}
	return false
}

func createMovieFile(line string) MovieFile {
	lineSeps := strings.Split(line, ",")
	return MovieFile{
		MovieFileName:    lineSeps[0],
		MovieUpdDatetime: lineSeps[1]}
}

// ReadCurrent ... 最新の動画ファイル情報リストを取得
func ReadCurrent(dirPath string) (*MovieList, error) {

	return nil, nil
}
