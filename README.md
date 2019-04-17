# Gitハンズオン入門資料

[公開サイト](https://oriishitakahiro.github.io/git-intro/)

[![CircleCI](https://circleci.com/gh/OriishiTakahiro/git-intro/tree/master.svg?style=svg)](https://circleci.com/gh/OriishiTakahiro/git-intro/tree/master)

## 対象者

- Gitを使ったことがなく，Gitを使うとどんな世界が見えるのか知りたい人
- 研究やチーム開発で制作物をきちんと管理したい人
- エンジニアを目指す者として最低限のGitの知識が欲しい人

## 本記事で取り扱う範囲

本記事では初めてGitを使う人が，

- Gitの重要性
- Gitの勘所

を掴むことを目的としています．

そのため，実際にGitを使う流れに沿った内容となっており，また周辺の取り巻く環境も含めて俯瞰できるような内容にしたつもりです．

一方で初歩的な内容が多く，Gitの仕組みそのものから深く理解したい方にとっては非常に物足りない内容になっていると思いますので，ご容赦願います．

## 準備するもの

本記事の内容を実施するにあたり，以下のものが必要となります．

- GitHubアカウント
- Webブラウザが動作するPC
- インターネット利用環境
- プログラミング言語への慣れ  
  (Go言語への理解があるとなおよい)

## 全体の流れ

全体の流れについて簡単に説明します．

[Gitについて](docs/chap1.md)では，Gitがどういうものであるか，Gitが何故必要とされるのかを簡単に説明し，学習意欲を上げてもらいます．

[Git基礎](docs/chap2.md)では，Gitの構成やローカルのみで行う基本操作を解説し，体験してもらいます．
バージョン管理を手元で行いたいだけであればこの章までで十分です．

[GitHubと合わせて](docs/chap3.md)では，ホスティングサービスであるGitHubを利用し，手元の成果物をリモートサーバ上で管理する際の流れを体験してもらいます．
GitHubを使ってみたいけど複数人開発はしない，GitHubで成果物を公開したいだけといった方はこの章までで十分です．

[チームで使うGit](docs/chap4.md)では，複数人で共同開発する際の流れを体験してもらいます．
チーム開発を前提としているため，複数人での参加が前提となります．
またGo言語を用いた開発を演習があるため，それまでの章と比較し，ハードルは大きく上がります．

[Tips](docs/chap5.md)では，Gitを使う上で知っておいてほしい豆知識やワードについて解説します．

## コマンドの読み方

本記事の中では次のようにコードを記述している箇所があります．

先頭が次の文字で始まるものにはそれぞれ意味が異なります．

- '#'の行はコメントなので，読んでください
- '$'の行は実行するコマンドなので，ターミナルに入力してください
- '>'の行は実行結果なので，自分のターミナルで表示された画面と見くらべてください

```sh
# helloと表示する
$ echo hello
> hello
```

それとは別にファイルに記述してほしいプログラムや設定のコードなどは特に先頭文字を配置していません．

## その他

著者:  [@OriishiTakahiro](https://twitter.com/OriishiTakahiro)

本記事の使用，引用については特に制限をしません．

また本記事における間違いや語弊を招く表現などを発見した場合にはプルリクエストを頂けると非常に嬉しいです．
