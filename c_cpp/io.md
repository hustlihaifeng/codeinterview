1. [C++中Txt文件读取和写入(方法汇总)](https://blog.csdn.net/lz20120808/article/details/49622787)
2. [C++文件读写详解（ofstream,ifstream,fstream）](https://blog.csdn.net/kingstar158/article/details/6859379)
3. [Which C I/O library should be used in C++ code? closed](https://stackoverflow.com/questions/119098/which-c-i-o-library-should-be-used-in-c-code)
4. [http://www.cplusplus.com/reference/fstream/fstream/open/](http://www.cplusplus.com/reference/fstream/fstream/open/)这里open有可能`may throw ios_base::failure if that state flag was registered using member exceptions`)

# stdio.h

1. `printf %02x`  的输出结果
	- 以16进制输出，右对齐，最少输出两位宽度，不足补0，超过的全部输出。[【printf,%02X和%x有什么区别】](https://blog.csdn.net/kebu12345678/article/details/78119917)
	- **对高位，有符号数进行符号扩展，无符号数进行0扩展**。[【printf("%02x\n", c) 之坑------浪费0.5小时】](https://blog.csdn.net/stpeace/article/details/51794230) [【printf格式化输出%x时的分析】](https://blog.csdn.net/jmh1996/article/details/53884246)