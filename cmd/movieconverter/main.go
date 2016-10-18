package main

import (
	"flag"
	"log"
	"os"
	"strings"

	mc "github.com/sky0621/work-movieconverter"
)

var version = "0.1.0"

func main() {
	log.Println("[START]movieconverter")
	arg := parseFlag()
	if arg == nil {
		os.Exit(0)
	}
	mc.Run(arg)
	log.Println("[END]movieconverter")
}

func parseFlag() *mc.Arg {
	var showVersion bool

	var targetDir string
	var sleepSecond int
	var movieSuffix string
	var outputDir string

	flag.BoolVar(&showVersion, "v", false, "バージョン")
	flag.StringVar(&targetDir, "t", ".", "監視対象ディレクトリ")
	flag.IntVar(&sleepSecond, "s", 60, "監視間隔（秒）")
	flag.StringVar(&movieSuffix, "x", "mp4", "監視対象動画ファイル拡張子群")
	flag.StringVar(&outputDir, "o", "out", "変換結果出力先ディレクトリ")
	flag.Parse()
	if showVersion {
		log.Println("version:", version)
		return nil
	}

	return &mc.Arg{
		TargetDir:   targetDir,
		SleepSecond: sleepSecond,
		MovieSuffix: strings.Split(movieSuffix, ","),
		OutputDir:   outputDir,
	}
}
