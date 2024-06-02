# Introduction
Go学習用環境

# 環境構築
Macを想定

```shell
$ make install
```

# コマンド実行
## main.goを実行
```shell
$ make main

# or...

$ docker compose exec go bash -c 'go run /go/src/app/main.go'
```

## コンテナにアタッチ
```shell
$ make go

# or...

$ docker compose exec go bash
```

# 標準入力取得
```go
scanner := bufio.NewScanner(os.Stdin)
var a int
if scanner.Scan() {
	a, _ = strconv.Atoi(scanner.Text())
}

var b []int
if scanner.Scan() {
	line := scanner.Text()
	// 空白(スペースやタブ)を区切り文字として文字列を分割
	fields := strings.Fields(line)
	for _, field := range fields {
		num, _ := strconv.Atoi(field)
		b = append(b, num)
	}
}

fmt.Printf("a: %d\n", a)
fmt.Printf("b: %d\n", b)
```