package main

import (
	"errors"
	sc "strconv"
)

func validate(args []string) error {
	// 長さが4未満 (command, lhs, rhsが全て与えられていない) ならエラー
	if len(args) < 4 {
		return errors.New("引数の長さが足りません (4つ以上)")
	}
	// args[2], args[3]がInt値に変換できるか？
	if _, err := sc.Atoi(args[2]); err != nil {
		return errors.New("2番目の引数がIntに変換できませんでした\n" + err.Error())
	}
	if _, err := sc.Atoi(args[3]); err != nil {
		return errors.New("3番目の引数がIntに変換できませんでした\n" + err.Error())
	}
	return nil
}
