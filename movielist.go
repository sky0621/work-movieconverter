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

// ReadMovieList ... 動画ファイル情報リストを読み込むインタフェース
type ReadMovieList interface {
	ReadMovieList() (*MovieList, error)
}

// PreviousMovieList ... 前回監視時の動画ファイル情報リスト取得用の構造体
type PreviousMovieList struct {
	TargetDir string
}

// ReadMovieList ... 前回監視時に作成した動画ファイル情報リストを読み込み、構造体として返却
func (p *PreviousMovieList) ReadMovieList() (*MovieList, error) {
	file, err := os.OpenFile(filepath.Join(p.TargetDir, "prev.list"), os.O_RDONLY, 0)
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

// CurrentMovieList ... 最新の動画ファイル情報リスト取得用の構造体
type CurrentMovieList struct {
	TargetDir   string
	MovieSuffix []string
}

// ReadMovieList ... 指定ディレクトリ配下の動画ファイルの一覧から、動画ファイル情報のリストを構造体として返却
func (c *CurrentMovieList) ReadMovieList() (*MovieList, error) {
	log.Println(c.TargetDir)
	var movieFiles []MovieFile
	err := filepath.Walk(c.TargetDir, walkFunc(&movieFiles, c.MovieSuffix))
	if err != nil {
		log.Printf("指定ディレクトリ(%s)配下の動画ファイル一覧読み込み時にエラーが発生しました。 [ERROR]%s\n", c.TargetDir, err)
		return nil, err
	}
	return &MovieList{MovieFiles: movieFiles}, nil
}

// [MEMO]ファイルウォークの場合は、任意のデータを詰めたりするのにクロージャ使う必要あり
func walkFunc(movieFiles *[]MovieFile, movieSuffix []string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		ext := filepath.Ext(info.Name())
		// [MEMO]こんなユーティリティすら無い様子・・・。
		for _, suffix := range movieSuffix {
			if ext == suffix {
				*movieFiles = append(
					*movieFiles,
					MovieFile{
						MovieFileName:    info.Name(),
						MovieUpdDatetime: info.ModTime().String()})
			}
		}
		return nil
	}
}
