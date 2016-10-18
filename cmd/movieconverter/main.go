package main

import (
	"flag"
	"log"
	"os"

	mc "github.com/sky0621/work-movieconverter"
)

var version = "0.1.0"

func main() {
	log.Println("[START]movieconverter")
	arg, err := parseFlag()
	if err != nil {
		os.Exit(1) // おそらくバリデーションに失敗
	}
	if arg == nil {
		os.Exit(0)
	}
	mc.Run(arg)
	log.Println("[END]movieconverter")
}

func parseFlag() (*mc.Arg, error) {
	var showVersion bool

	var targetDir string
	var sleepSecond string
	var movieSuffix string
	var outputDir string
	var logDir string

	flag.BoolVar(&showVersion, "v", false, "バージョン")
	flag.StringVar(&targetDir, "t", ".", "監視対象ディレクトリ")
	flag.StringVar(&sleepSecond, "s", "60", "監視間隔（秒）")
	flag.StringVar(&movieSuffix, "x", "mp4", "監視対象動画ファイル拡張子群")
	flag.StringVar(&outputDir, "o", "out", "変換結果出力先ディレクトリ")
	flag.StringVar(&logDir, "l", ".", "ログ出力先ディレクトリ")
	flag.Parse()
	if showVersion {
		log.Println("version:", version)
		return nil, nil
	}

	return mc.NewArg(targetDir, sleepSecond, movieSuffix, outputDir, logDir)
}
