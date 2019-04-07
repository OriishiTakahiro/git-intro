# Gitとは

Gitとは一言にバージョン管理ツールと呼ばれるもので，

- Git
- Mecurial
- SVN (Subversion)
- CVS (Concurrent Versions System)
- Bazzar
- Monotone

などの種類が在り，Gitはその中でおそらく一番メジャーであろうものになっています．

プログラムを開発を行うとき，ほとんどの場合において

- 長期間の開発
- 複数人での共有
- 度重なる変更や修正
- バグや機能修正に伴うロールバック

といった状態にあると思います．

バージョン管理システムは，ある時点をスナップショットのような形で残し管理するツールで，これらに伴う問題を解決する非常に良い手段となっています．

また殆どの開発現場においてバージョン管理ツールは導入されており，ITエンジニアにとっても必須のツールとなっています．

(またこれは個人的な意見ですが，開発現場にバージョン管理ツールが導入されていない時点で，その現場はブラックだと言ってもいいと思いますね...)

必須スキルと化したGitについて，是非肌感を掴んでいってください．

# Gitがないとき，あるとき

まずはGitがどんな役割を果たすものであるのか，開発の実例を通して見ていきましょう．

ここではプログラミング言語Goで簡単な計算ツールGalcを開発する過程を見ていきます．

初期のGalcは足し算しかなく，`go run galc.go add 1 2`とすると1と2の足し算が返ってきます．

```go
package main

import (
	"fmt"
	"os"
	sc "strconv"
)

func main() {
	lhs, _ := sc.Atoi(os.Args[2])
	rhs, _ := sc.Atoi(os.Args[3])
	switch os.Args[1] {
	case "add":
		fmt.Println(lhs + rhs)
	}
}
```

## Gitがないとき

### 増えるgalc\_x.x.x.zip

まず最初のgalcのコードを`galc_0.0.1.zip`という形で保存します．

次に引き算機能を入れてみましょう．

```go
package main

import (
	"fmt"
	"os"
	sc "strconv"
)

func main() {
	lhs, _ := sc.Atoi(os.Args[2])
	rhs, _ := sc.Atoi(os.Args[3])
	switch os.Args[1] {
	case "add":
		fmt.Println(lhs + rhs)
	case "sub":
		fmt.Println(lhs - rhs)
	}
}
```

これが`galc_0.0.2.zip`となるわけですね．

次に割り算を入れてみましょう．

```go
package main

import (
	"fmt"
	"os"
	sc "strconv"
)

func main() {
	lhs, _ := sc.Atoi(os.Args[2])
	rhs, _ := sc.Atoi(os.Args[3])
	switch os.Args[1] {
	case "add":
		fmt.Println(lhs + rhs)
	case "sub":
		fmt.Println(lhs - rhs)
	case "div":
		fmt.Println(lhs / rhs)
	}
}
```

`galc_0.0.3.zip`の完成です．

つぎに`os.Args[2]`と`os.Args[3]`が整数でないときにエラーが出て落ちるので，変換チェックを入れないとですね．

```go
package main

import (
	"fmt"
	"os"
	sc "strconv"
)

func main() {
	lhs, lerr := sc.Atoi(os.Args[2])
	rhs, rerr := sc.Atoi(os.Args[3])
	if lerr != nil || rerr != nil {
		fmt.Println("2, 3番目の引数の何れかが整数に変換できません")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "add":
		fmt.Println(lhs + rhs)
	case "sub":
		fmt.Println(lhs - rhs)
	case "div":
		fmt.Println(lhs / rhs)
	}
}
```

これで`galc_0.0.3.zip`ができました．

これが延々と続くわけなので，機能開発を行う度にファイルが増えます．

### その後の地獄

 - 0除算エラー対応
 - 今まで複数人でやってきたコードを混ぜる
 - 冪乗計算の追加
 - 掛け算の追加
 - リリース
 - 実数への対応

これで8個のzipファイルができました．

ここで「乗除の計算は要らないです」と言われたらどうしますか?

全てのzipファイルを開いて，差分を手直しして新しいzipファイルを作って...

そんなことしていたら，動かないものができますよね?

### つらみの総括

現場では複数人が同時並行で進めるわけです．こんなのやってたら死にます．

(独りでも死にます)

この方法では

- ファイル管理が面倒
- どのバージョンで何を追加，修正したのか不明
- いちいち展開しないと覗けない
- 一部の修正だけ無効にしたい，取り込みたいが難しい
- 複数人の開発者でのコードが同期しない
- それぞれのファイルで何を実装したか，修正したかを別途管理しないといけない

といった問題があります．

## Gitがあるとき

Gitがない場合で問題になったことは，Gitを__正しく使えば__ほとんど解決します．

Gitは簡単なツール操作で以下のことを行ってくれます．

- スナップショットの作成
- スナップショットのメタデータ作成と管理  
  (いつ？誰が？どんな変更を加えたか？)
- ログの閲覧
- 差分の閲覧
- 複数人のコード共有
- 複数人のコードを混ぜ合わせる
