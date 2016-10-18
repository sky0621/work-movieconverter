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
	log.Println(arg)
	for {
		// 指定ディレクトリ配下を最後に監視した際の動画ファイルリストを取得
		prev := PreviousMovieList{TargetPath: arg.TargetDir}
		previousList, err := prev.ReadMovieList()
		if err != nil {
			return
		}
		log.Println(previousList)

		time.Sleep(time.Duration(arg.SleepSecond) * time.Second)
	}
}

func runConvertMovies() {

}
