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
