# cat コマンド

## テストの実行

```sh
$ make test
```

## 実現した機能

- `mycat file1`:file1 の内容を表示する。
- `mycat < file1`:file1 の内容を表示する。
- `mycat file1 file2`:file1 および file2 の内容を連結し、端末上にその結果を表示する。
- `mycat --help`:コマンドの使用方法を表示する。
- `mycat --show-ends file1`:各行の末尾に$を追加した file1 の内容を表示する。
- `mycat -E file1`:各行の末尾に$を追加した file1 の内容を表示する。
