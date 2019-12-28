# go-ums 编译与运行 v0.1.0

## 1. 测试
注:
我的源码在 
```
/Users/qinshen/go/src/github.com/tsingson/go-ums/
```

测试 register 
```
cd /Users/qinshen/go/src/github.com/tsingson/go-ums/pkg/services

/Users/qinshen/go/src/github.com/tsingson/go-ums/pkg/services $  go test -v .
=== RUN   TestAccountStore_Register
--- PASS: TestAccountStore_Register (0.00s)
=== RUN   TestAccountStore_Register_Error
--- PASS: TestAccountStore_Register_Error (0.00s)
=== RUN   TestAccountStore_Login
--- PASS: TestAccountStore_Login (0.00s)
PASS
ok  	github.com/tsingson/go-ums/pkg/services	0.022s


/Users/qinshen/go/src/github.com/tsingson/go-ums/pkg/services $  go test -bench .
goos: darwin
goarch: amd64
pkg: github.com/tsingson/go-ums/pkg/services
BenchmarkAccountStore_Register-8   	 1000000	      2381 ns/op
BenchmarkAccountStore_Login-8      	 2000000	       908 ns/op
PASS
ok  	github.com/tsingson/go-ums/pkg/services	10.419s

```



测试 web
```
cd /Users/qinshen/go/src/github.com/tsingson/go-ums/pkg/web/fast

/Users/qinshen/go/src/github.com/tsingson/go-ums/pkg/web/fast $  go test -v .
=== RUN   TestHttpSERV_RegisterHandler
--- PASS: TestHttpSERV_RegisterHandler (0.00s)
PASS
ok  	github.com/tsingson/go-ums/pkg/web/fast	0.018s


/Users/qinshen/go/src/github.com/tsingson/go-ums/pkg/web/fast $  go test -bench .
goos: darwin
goarch: amd64
pkg: github.com/tsingson/go-ums/pkg/web/fast
BenchmarkHttpServer_RegisterHandler-8   	  100000	     18683 ns/op
PASS
ok  	github.com/tsingson/go-ums/pkg/web/fast	2.091s

```
## 2.  编译


编译

```
cd /Users/qinshen/go/src/github.com/tsingson/go-ums/cmd/cli
go install ./...
```

## 3. 运行

服务端
```
cd /Users/qinshen/go/bin
./go-ums-cli-server

```


客户端
```
cd /Users/qinshen/go/bin
/Users/qinshen/go/bin $  ./go-ums-cli-client register tsingson@gmail.com passwordddd
-------------------> http request object 
&model.AccountRequest{
  Email: "tsingson@gmail.com",
  Password: "passwordddd",
}
-------------------> http response
200
&model.Account{
  ID: 1557507549010348000,
  Email: "tsingson@gmail.com",
  Password: "passwordddd",
  Role: 0,
  Status: 1,
  CreatedAt: 1557507549010348000,
  UpdatedAt: 1557507549010348000,
  AccessToken: "",
}
/Users/qinshen/go/bin $  ./go-ums-cli-client register tsingson@gmail.com passwordddd
-------------------> http request object 
&model.AccountRequest{
  Email: "tsingson@gmail.com",
  Password: "passwordddd",
}
-------------------> http response
500
"account is existed, duplicated email address"

```

## 4.  fasthttp vs net/http ( gin-gonic/gin )

注1: 
 http client 测试代码完全一样, 服务器端业务逻辑代码完全一样


注2:
 这可能是偶然现象, 哈

### 4.1 fasthttp
map
```
/Users/qinshen/go/src/github.com/tsingson/go-ums/pkg/web/fast $  go test -bench .
goos: darwin
goarch: amd64
pkg: github.com/tsingson/go-ums/pkg/web/fast
BenchmarkHttpServer_RegisterHandler-8   	  100000	     19241 ns/op
PASS
ok  	github.com/tsingson/go-ums/pkg/web/fast	2.140s
```

orcaman/concurrent-map

```
/Users/qinshen/go/src/github.com/tsingson/go-ums/pkg/web $  cd fast/
/Users/qinshen/go/src/github.com/tsingson/go-ums/pkg/web/fast $   go test -bench .

goos: darwin
goarch: amd64
pkg: github.com/tsingson/go-ums/pkg/web/fast
BenchmarkHttpServer_RegisterHandler-8   	  100000	     17584 ns/op
PASS
ok  	github.com/tsingson/go-ums/pkg/web/fast	1.972s
```
### 4.2 gin-gonic/gin 

map 
```
/Users/qinshen/go/src/github.com/tsingson/go-ums/pkg/web/xgin $  go test -bench .
goos: darwin
goarch: amd64
pkg: github.com/tsingson/go-ums/pkg/web/xgin
BenchmarkHttpServer_RegisterHandler-8   	   10000	    104056 ns/op
PASS
ok  	github.com/tsingson/go-ums/pkg/web/xgin	1.076s
```

orcaman/concurrent-map
```
/Users/qinshen/go/src/github.com/tsingson/go-ums/pkg/web/xgin $  go test -bench .
goos: darwin
goarch: amd64
pkg: github.com/tsingson/go-ums/pkg/web/xgin
BenchmarkHttpServer_RegisterHandler-8   	   10000	    111213 ns/op
PASS
ok  	github.com/tsingson/go-ums/pkg/web/xgin	1.148s
```

