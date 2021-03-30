# 言語の基本
## 変数定義  
### 明示的  
- 単体  
~~~  
var x int
~~~  
- 複数(同じ型)  
~~~  
var x, y, z int
~~~  
- 複数(異なる型)  
~~~  
var (
	x, y int
	z string
)
~~~  
### 暗黙的  
- 単体  
「ローカル変数」として定義できる  
~~~  
x := 1
~~~  
- 複数  
~~~  
var (
	x = 1
	y = 2
	z = "z"
)
~~~  
## 関数定義  
### 基本  
~~~  
func doSomething(x, y int) int {
	// 処理
}
~~~  
### 複数の戻り値  
~~~  
func doSomething(x, y, z int) (int, int) {
	v := x + z
	w := y + z
	return v, w
}
~~~  
### 無名関数  
~~~  
f := func(x, y int) int { return x + y }
~~~  
### 関数を返す関数  
~~~  
func doSomething() func() {
	return func() {
		// 処理
	}
}
doSomething()()
~~~  
### 関数を引数にとる関数  
~~~  
func doSomething(f func()) {
	f()
}
doSomething(func() {
	// 処理
})
~~~  
### クロージャ  
~~~  
func doSomething() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}
f := doSomething()
f("A") // => ""
f("B") // => "A"
~~~  
### クロージャによるジェネレータ  
~~~  
func doSomething() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
f := doSomething()
f() // => 1
f() // => 2
~~~  
## 定数定義  
~~~  
const (
	X = 1
	Y = 2
	Z // 2
)
~~~  
## スコープ  
先頭1文字を大文字にした場合、他パッケージから参照できる  
~~~  
// foo/app.go
func DoSomething() string {
	return "A"
}

func doSomething() string {
	return "A"
}

// main.go
import (
	"fmt"
	"./foo"
)
fmt.Println(foo.DoSomething()) // => A
fmt.Println(foo.doSomething()) // エラー
~~~  
## for  
~~~  
for i := 0; i < 10; i++ {
	// 処理
}
~~~  
### 範囲節  
~~~
for i, v := range array {
	// 処理
}
~~~  
### ラベル文  
~~~
LOOP:
for i := 1; i < 10; i++ {
	for j := 1; j < 10; j++ {
		if i%j == 0 {
			break LOOP
		}
	}
}
~~~  
## if  
### 基本型  
~~~  
if i % 15 == 0 {
	fmt.Println("Fizz Buzz")
} else if i % 3 == 0 {
	fmt.Println("Fizz")
} else if i % 5 == 0 {
	fmt.Println("Buzz")
} else {
	fmt.Println(i)
}
~~~  
## 省略型  
~~~  
if x, err := doSomething(); err != nil {
	// エラーケースの処理
}
~~~  
## switch  
### 基本型  
~~~  
switch n {
case 1:
	// 処理
case 2, 3: // 複数の値OK
	// 処理
default:
	// 処理
}
~~~  
### フォールスルー  
~~~
s := "あ"
switch s {
case "あ":
	s += "い"
	fallthrough
case "い":
	s += "う"
	fallthrough
case "う":
	s += "え"
	fallthrough
default:
	s += "お"
}
fmt.Println(s) // => あいうえお
~~~  
### case式  
~~~
switch {
case i%15 == 0:
	fmt.Println("Fizz Buzz")
case i%3 == 0:
	fmt.Println("Fizz")
case i%5 == 0:
	fmt.Println("Buzz")
default:
	fmt.Println(i)
}
~~~  
## defer  
~~~
func doSomething() {
	defer func() {
		fmt.Println("Defer1")
		fmt.Println("Defer2")
		fmt.Println("Defer3")
	}()
	fmt.Println("Done")
}
doSomething()

/* 出力 */
Done
Defer1
Defer2
Defer3
~~~  
## init  
パッケージの初期化が目的  
関数mainに先立って実行される  
~~~
func init() {
	fmt.Println("Init")
}

func main() {
	fmt.Println("Main")
}

/* 出力 */
Init
Main
~~~  
