package movieconverter

import (
	"reflect"
	"testing"
)

func TestReadPrevious_NoFile(t *testing.T) {
	t.Log("[仕様]prev.listが存在しない場合 -> nil を返す。")
	movieList, err := ReadPrevious("testdata/movielist/previous_nofile")
	if err != nil {
		t.Fatalf("エラーが発生しました。 %s\n", err)
	}
	if movieList != nil {
		t.Fatal("nil を返しませんでした。")
	}
}

func TestReadPrevious_Exist(t *testing.T) {
	t.Log("[仕様]prev.listが存在する場合 -> 読み込み、動画ファイル情報リストの構造体を返す。")
	movieList, err := ReadPrevious("testdata/movielist/previous_exist")
	if err != nil {
		t.Fatalf("エラーが発生しました。 %s\n", err)
	}
	if movieList == nil {
		t.Fatal("nil を返しました。")
	}
	expected := buildValidExpected()
	if !reflect.DeepEqual(movieList, expected) {
		t.Fatalf("結果と期待値が要素レベルで異なっています。\n[実績値]：%#v\n[期待値]：%#v", movieList, expected)
	}
}

func TestReadPrevious_AbNormal(t *testing.T) {
	t.Log("[仕様]prev.listが存在するが不正な行が含まれる場合 -> 不正な行以外を読み込み、動画ファイル情報リストの構造体を返す。")
	movieList, err := ReadPrevious("testdata/movielist/previous_abnormal")
	if err != nil {
		t.Fatalf("エラーが発生しました。 %s\n", err)
	}
	if movieList == nil {
		t.Fatal("nil を返しました。")
	}
	expected := buildValidExpected()
	if !reflect.DeepEqual(movieList, expected) {
		t.Fatalf("結果と期待値が要素レベルで異なっています。\n[実績値]：%#v\n[期待値]：%#v", movieList, expected)
	}
}

func buildValidExpected() *MovieList {
	movieList := MovieList{
		[]MovieFile{
			MovieFile{MovieFileName: "Sample01.mp4", MovieUpdDatetime: "2016-10-18 17:03:49"},
			MovieFile{MovieFileName: "Sample02.mp4", MovieUpdDatetime: "2016-09-05 05:52:00"},
		},
	}
	return &movieList
}
