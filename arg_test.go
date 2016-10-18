package movieconverter

import (
	"reflect"
	"testing"
)

func TestNewArg_Normal(t *testing.T) {
	t.Log("[仕様]正しいプログラム引数が与えられた場合 -> 期待された構造体Arg を返す。")
	arg, err := NewArg(
		".",
		"60",
		"mp4",
		"out",
		".")

	if err != nil {
		t.Fatalf("エラーが発生しました。 %s\n", err)
	}

	expected := &Arg{
		TargetDir:   ".",
		SleepSecond: 60,
		MovieSuffix: []string{"mp4"},
		OutputDir:   "out",
		LogDir:      "."}
	if !reflect.DeepEqual(arg, expected) {
		t.Fatalf("結果と期待値が要素レベルで異なっています。\n[実績値]：%#v\n[期待値]：%#v", arg, expected)
	}
}

func TestNewArg_InvalidTargetDir(t *testing.T) {
	t.Log("[仕様]プログラム引数(TargetDir)が不正な場合 -> エラーを返す。")
	_, err := NewArg(
		"",
		"60",
		"mp4",
		"out",
		".")

	if err == nil {
		t.Fatal("空文字を与えたが、エラーが発生しませんでした。 \n")
	}

	_, err = NewArg(
		"12345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789011234567890123456789012345678901234567890123456789012345678902345678901234567890123456789012345678901234567890123456789012345678901234567",
		"60",
		"mp4",
		"out",
		".")

	if err == nil {
		t.Fatal("256字を超えたが、エラーが発生しませんでした。 \n")
	}
}

func TestNewArg_InvalidSleepSecond(t *testing.T) {
	t.Log("[仕様]プログラム引数(SleepSecond)が不正な場合 -> エラーを返す。")
	_, err := NewArg(
		".",
		"",
		"mp4",
		"out",
		".")

	if err == nil {
		t.Fatal("空文字を与えたが、エラーが発生しませんでした。 \n")
	}

	_, err = NewArg(
		".",
		"1A0",
		"mp4",
		"out",
		".")

	if err == nil {
		t.Fatal("数値に変換できない文字を与えたが、エラーが発生しませんでした。 \n")
	}
}

func TestNewArg_InvalidMovieSuffix(t *testing.T) {
	t.Log("[仕様]プログラム引数(MovieSuffix)が不正な場合 -> エラーを返す。")
	_, err := NewArg(
		".",
		"60",
		"",
		".",
		".")

	if err == nil {
		t.Fatal("空文字を与えたが、エラーが発生しませんでした。 \n")
	}

	_, err = NewArg(
		".",
		"60",
		",",
		".",
		".")

	if err == nil {
		t.Fatal("空文字を与えたが、エラーが発生しませんでした。 \n")
	}
}

func TestNewArg_InvalidOutputDir(t *testing.T) {
	t.Log("[仕様]プログラム引数(OutputDir)が不正な場合 -> エラーを返す。")
	_, err := NewArg(
		".",
		"60",
		"mp4",
		"",
		".")

	if err == nil {
		t.Fatal("空文字を与えたが、エラーが発生しませんでした。 \n")
	}

	_, err = NewArg(
		".",
		"60",
		"mp4",
		"12345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789011234567890123456789012345678901234567890123456789012345678902345678901234567890123456789012345678901234567890123456789012345678901234567",
		".")

	if err == nil {
		t.Fatal("256字を超えたが、エラーが発生しませんでした。 \n")
	}
}

func TestNewArg_InvalidLogDir(t *testing.T) {
	t.Log("[仕様]プログラム引数(LogDir)が不正な場合 -> エラーを返す。")
	_, err := NewArg(
		".",
		"60",
		"mp4",
		".",
		"")

	if err == nil {
		t.Fatal("空文字を与えたが、エラーが発生しませんでした。 \n")
	}

	_, err = NewArg(
		".",
		"60",
		"mp4",
		".",
		"12345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789011234567890123456789012345678901234567890123456789012345678902345678901234567890123456789012345678901234567890123456789012345678901234567")

	if err == nil {
		t.Fatal("256字を超えたが、エラーが発生しませんでした。 \n")
	}
}
