package movieconverter

import (
	"reflect"
	"testing"
)

func TestReadPrevious_NoFile(t *testing.T) {
	t.Log("[仕様]prev.listが存在しない場合 -> nil を返す。")
	prev := PreviousMovieList{TargetDir: "testdata/movielist/previous_nofile"}
	movieList, err := prev.ReadMovieList()
	if err != nil {
		t.Fatalf("エラーが発生しました。 %s\n", err)
	}
	if movieList != nil {
		t.Fatal("nil を返しませんでした。")
	}
}

func TestReadPrevious_NoDir(t *testing.T) {
	t.Log("[仕様]存在しないパスを指定した場合 -> nil を返す。")
	prev := PreviousMovieList{TargetDir: "testdata/movielist/previous_nodir"}
	movieList, err := prev.ReadMovieList()
	if err != nil {
		t.Fatalf("エラーが発生しました。 %s\n", err)
	}
	if movieList != nil {
		t.Fatal("nil を返しませんでした。")
	}
}

func TestReadPrevious_Exist(t *testing.T) {
	t.Log("[仕様]prev.listが存在する場合 -> 読み込み、動画ファイル情報リストの構造体を返す。")
	prev := PreviousMovieList{TargetDir: "testdata/movielist/previous_exist"}
	movieList, err := prev.ReadMovieList()
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
	prev := PreviousMovieList{TargetDir: "testdata/movielist/previous_abnormal"}
	movieList, err := prev.ReadMovieList()
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

func TestReadCurrent_Exist(t *testing.T) {
	t.Log("[仕様]指定のディレクトリ配下に動画ファイルが存在する場合 -> 読み込み、動画ファイル情報リストの構造体を返す。")
	curr := CurrentMovieList{TargetDir: "./testdata/movielist/current_exist/", MovieSuffix: []string{".mp4"}}
	movieList, err := curr.ReadMovieList()
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
			MovieFile{MovieFileName: "sample01.mp4", MovieUpdDatetime: "2016-10-19 11:48:04.2648115 +0900 JST"},
			MovieFile{MovieFileName: "sample02.mp4", MovieUpdDatetime: "2016-10-19 11:48:19.0122648 +0900 JST"},
		},
	}
	return &movieList
}
