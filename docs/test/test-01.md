# Testing and Tunning  测试与调优



## 0. 环境准备



## 1. 测试

```
cd ./pkg/web/fast

go test -memprofile mem.out -memprofilerate 1 -cpuprofile cpu.prof -bench .
```



查看 cpu

```
go tool pprof -http=0.0.0.0:4231 cpu.prof
go tool pprof -http=0.0.0.0:4232 mem.out
```

