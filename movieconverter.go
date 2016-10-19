package movieconverter

import (
	"log"
	"time"
)

// 指定のディレクトリ（※指定なしはカレント）を監視（指定秒間隔（※指定なしは60秒）でポーリング）し、
// 配下の動画ファイル（★拡張子は要検討）のCREATE/UPDATE/DELETEがあれば、
// 外部コマンド「ffmpeg」（★指定可能パラメータは要検討）を実行して、指定の出力先（※指定なしはカレント直下の out ディレクトリ）に変換結果を出力する。

// Run ...
func Run(arg *Arg) {
	for {
		// 指定ディレクトリ配下を最後に監視した際の動画ファイルリストを取得
		prev := PreviousMovieList{TargetDir: arg.TargetDir}
		previousList, perr := prev.ReadMovieList()
		if perr != nil {
			return
		}
		log.Println(previousList)

		// 指定ディレクトリ配下の動画ファイルの一覧から動画ファイルリストを取得
		curr := CurrentMovieList{TargetDir: arg.TargetDir, MovieSuffix: arg.MovieSuffix}
		currentList, cerr := curr.ReadMovieList()
		if cerr != nil {
			return
		}
		log.Println(currentList)

		time.Sleep(time.Duration(arg.SleepSecond) * time.Second)
	}
}

func runConvertMovies() {

}
