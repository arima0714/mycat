# cat コマンド

## 実現したい機能

* ファイルが存在しなければ標準入力から読み込む
* dockerコンテナでビルドできるようにする
    * マルチステージビルドを使う
* Makefile内で`docker build`, `docker run` をできるようにする

## 実現した機能

* `cat file1`：file1の内容を表示する
* `cat file1 file2`：file1およびfile2の内容を連結し、端末上にその結果を表示する
