# 注意点

1. ReadLine能读取的行大小有限制，超过行大小时只会读取部分行，下次接着读取。

# 解决方法

1. 循环读取，直到一行末尾

```go
// ReadLongLine 调用ReadLine来读取一行，解决ReadLine的长行截断问题
// 外部使用需要判断`io.EOF`和普通error
func ReadLongLine(bufreader *bufio.Reader) (line []byte, err error) {
	var rst bytes.Buffer
	err = nil
	isPrefix := true
	for err == nil && isPrefix {
		line, isPrefix, err = bufreader.ReadLine()
		rst.Write(line)
	}

	if err != nil {
		if err == io.EOF {
			return rst.Bytes(), err
		}
		beego.Warn("ReadLongLine failed for: ", err, " in ", FileLineFunc(1))
		return nil, err
	}

	return rst.Bytes(), nil
}
```

注意检测`io.EOF`

# 参考资料

- [https://golang.org/pkg/bufio/#Reader.ReadLine](https://golang.org/pkg/bufio/#Reader.ReadLine)