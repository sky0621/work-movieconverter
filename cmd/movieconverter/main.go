package main

import "log"

// 指定のディレクトリ（※指定なしはカレント）を監視（指定秒間隔（※指定なしは10分）でポーリング）し、
// 配下の動画ファイル（★拡張子は要検討）のCREATE/UPDATE/DELETEがあれば、
// 外部コマンド「ffmpeg」（★指定可能パラメータは要検討）を実行して、指定の出力先（※指定なしはカレント直下の out ディレクトリ）に変換結果を出力する。
func main() {
	log.Println("[START]movieconverter")

	log.Println("[END]movieconverter")
}
