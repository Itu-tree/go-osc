# build

```
go build OSCSend.go
```

- GUIを表示したくない場合
```
go build -ldflags="-H windowsgui" OSCSend.go
```
タスクスケジューラーを使う場合にcmdが表示されなくなる


# Run on Windows

```
OSCSend.go
```