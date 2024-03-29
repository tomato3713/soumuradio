# soumuradio
soumuradio is a Go client library for accessing the radio station etc.
information search API of Ministry of Internal Affairs and Communications in
Japan

API Document Link: [Web-API機能（無線局等情報検索）について](https://www.tele.soumu.go.jp/j/musen/webapi/)

[![GoDev][godev-image]][godev-url]

## Install

```sh
go get github.com/tomato3713/soumuradio
```

## Usage

### Import

```sample01.go
import github.com/tomato3713/soumuradio
```

## Example
```sh
# List API Example
go run example/get-list/main.go
# Number API Example
go run example/get-number/main.go
```

### Number acquisition API

```go
cli, err := soumuradio.NewClient(http.DefaultClient)
if err != nil {
    os.Exit(1)
}

opts := soumuradio.NewNumOptions(soumuradio.License, soumuradio.Amateur)
result, err := cli.GetTotalCount(nil, opts)
if err != nil {
    os.Exit(1)
}

fmt.Printf("Amateur Radio TotalCount is %d\n", result)
// example output -> Amateur Radio TotalCount is 393183
```

### Logger

You can set debug option if you print out library logs for debug.

```go
soumuradio.EnableDebug()
// Print out debug logs like under line
// [go-soumuradio]2019/11/05 15:32:52 request URL: https://www.tele.soumuradio.go.jp/musen/num?MC=1&OF=2&OW=AT&ST=1
```

If you use your original logger, you can set it.

```go
soumuradio.SetLogger(log.New(os.Stdout, "[myapp]", log.LstdFlags)
```

## Note

このサービスは、総務省 電波利用ホームページのWeb-API 機能を利用して取得した情報
をもとに作成しているが、サービスの内容は総務省によって保証されたものではない

[godev-image]: https://pkg.go.dev/badge/github.com/tomato3713/soumuradio
[godev-url]: https://pkg.go.dev/github.com/tomato3713/soumuradio

