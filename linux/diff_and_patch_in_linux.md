# 单个文件差异

- 得到patch：

```shell
diff -up a.file a.file.new > a.patch
```

- 应用patch：将a.file更新为a.file.new的样子

```shell
cd /path/to/a.file/
patch -p0 < a.patch
```

# 文件夹差异
```shell
diff -uprN linux-2.6.28.8.orig/net/sunrpc/ linux-2.6.28.8/net/sunrpc/ > patch
```
- 可以用来对比线上版本差异


# 参考资料

- [Linux下生成patch和打patch](https://www.cnblogs.com/aaronLinux/p/5860552.html)