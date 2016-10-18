package movieconverter

// Arg プログラム引数
type Arg struct {
	// TargetDir 監視対象ディレクトリ
	TargetDir string
	// SleepSecond 監視間隔（秒）
	SleepSecond int
	// MovieSuffix 監視対象動画ファイル拡張子群
	MovieSuffix []string
	// OutputDir 変換結果出力先ディレクトリ
	OutputDir string
}
