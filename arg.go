package movieconverter

import (
	"fmt"
	"strconv"
	"strings"

	"gopkg.in/go-playground/validator.v9"
)

// Arg プログラム引数
type Arg struct {
	// TargetDir 監視対象ディレクトリ
	TargetDir string `validate:"required,max=256"` // [MEMO]（他のパス系もだけど）厳密にやるなら、ファイルパスとして正しいかのチェックも必要
	// SleepSecond 監視間隔（秒）
	SleepSecond int `validate:"required"` // [MEMO]最大値のチェックは入れた方がいいかな
	// MovieSuffix 監視対象動画ファイル拡張子群
	MovieSuffix []string `validate:"required"` // [MEMO]厳密にやるなら mp4 など取りうる値を定義してチェック
	// OutputDir 変換結果出力先ディレクトリ
	OutputDir string `validate:"required,max=256"`
	// LogDir ログ出力先ディレクトリ
	LogDir string `validate:"required,max=256"`
}

// NewArg ...
func NewArg(targetDir string, sleepSecond string, movieSuffix string, outputDir string, logDir string) (*Arg, error) {

	iSleepSecond, err := strconv.Atoi(sleepSecond)
	if err != nil {
		return nil, fmt.Errorf("プログラム引数(sleepSecond)のバリデーションエラー [ERROR] %s\n", err)
	}

	aryMovieSuffix := strings.Split(movieSuffix, ",")
	// [MEMO]空文字をスプリットした結果のスライス要素数は０ではないみたい・・・。
	for _, msfx := range aryMovieSuffix {
		if msfx == "" {
			return nil, fmt.Errorf("プログラム引数(movieSuffix)のバリデーションエラー [ERROR] Key: 'Arg.MovieSuffix' Error:Field validation for 'MovieSuffix' failed on the 'required' tag\n")
		}
	}

	arg := Arg{
		TargetDir:   targetDir,
		SleepSecond: iSleepSecond,
		MovieSuffix: aryMovieSuffix,
		OutputDir:   outputDir,
		LogDir:      logDir,
	}

	validate := validator.New()
	err = validate.Struct(arg)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return nil, fmt.Errorf("プログラム引数のバリデーションエラー [ERROR] %s\n", err.(*validator.InvalidValidationError))
		}
		return nil, fmt.Errorf("プログラム引数のバリデーションエラー [ERROR] %s\n", err.(validator.ValidationErrors))
	}

	return &arg, nil
}
