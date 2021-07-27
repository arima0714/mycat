# cat コマンド

## 実現したい機能

* ファイルが存在しなければ標準入力から読み込む

## 実現した機能

* `cat file1`：file1の内容を表示する
* `cat file1 file2`：file1およびfile2の内容を連結し、端末上にその結果を表示する
* dockerコンテナでビルドできるようにする
    * マルチステージビルドを使う
* Makefile内で`docker build`, `docker run` をできるようにする

# Dockerのマルチステージ・ビルド

イメージ構築において最も困難な目標の一つとして、イメージ容量を限りなく小さくするというものがある。それを実現するための手段としてshell tricksを駆使したり、レイヤを維持するロジックを用いたり、Dockerfileを複数用意するなどの必要があった。マルチステージビルドはこういった複雑な手法を使わず、シンプルにイメージ容量を小さくすることができる。[^dockerMultiStage]


`make build` , `make buildWithNotOptimization` で最適化済みコンテナ(`gopractice`)および非最適化コンテナ(`gopractice_not_optimization`)をビルドすることができる。
イメージのサイズは下記の通り。

```
docker images | grep gopractice
gopractice_not_optimization   latest    03f89e8722d8   4 minutes ago    867MB
gopractice                    latest    1083436f0346   7 minutes ago    10.1MB
```

$$
10.1 / 867 = 0.011...
$$

マルチステージ・ビルドを用いることで、イメージサイズを1％にまで削減することができた。

[^dockerMultiStage]:マルチステージ・ビルドを使う（https://docs.docker.jp/develop/develop-images/multistage-build.html）
