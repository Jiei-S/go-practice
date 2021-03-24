# プログラムの構成と実行
## 実行  
~~~  
$ go run main.go
~~~  
## ビルド  
### オプション指定なし  
カレントディレクトリ内全てのファイルをビルド対象とする  
実行ファイル名: カレントディレクトリの名前  
~~~  
$ go build
~~~  
### 関数mainを含むファイルの指定  
実行ファイル名: main    
~~~  
$ go build main.go
~~~  
### -oオプション  
~~~  
$ go build -o <実行ファイル名> <コンパイル対象ファイル>
~~~  
## テスト  
~~~  
$ go test <テスト対象パッケージ>
~~~  
### -vオプション  
個々のテストの詳細表示  
~~~  
$ go test -v <テスト対象パッケージ>
=== RUN   TestSoccer
--- PASS: TestSoccer (0.00s)
=== RUN   TestBaseball
--- PASS: TestBaseball (0.00s)
=== RUN   TestBasketball
--- PASS: TestBasketball (0.00s)
PASS
ok      <パッケージのパス>   0.017s
~~~  
