# 1. `__init__.py`的作用

- <https://stackoverflow.com/questions/448271/what-is-init-py-for>

- python有regular package和namespace package两类包

  - regular package 一般是里面包含`__init__.py`的目录，当regular package被导入时，`__init__.py`被隐藏执行，

  > Importing `parent.one` will implicitly execute `parent/__init__.py` and `parent/one/__init__.py`

  - namespace package一般用来管理不同目录下分散的包，详见[10.5 利用命名空间导入目录分散的代码](https://python3-cookbook.readthedocs.io/zh_CN/latest/c10/p05_separate_directories_import_by_namespace.html)

# 2. `#!/usr/bin/env python`

- <https://askubuntu.com/questions/88257/what-type-of-path-in-shebang-is-more-preferable>
- `#!/bin/bash `固定目录，`#!/usr/bin/env bash`找env输出的`$PATH`中第一个匹配到的`bash`

# 3. `# -*- coding: utf-8 -*-`指定编码

# 4. import

```python
import sys
import os, sys, time # 导入多个
import sys as system # 导入时重命名
import urllib.error # 某些子模块必须要使用点标记法才能导入，urllib和error的__init__.py都会执行
from functools import lru_cache # 上面这行代码可以让你直接调用lru_cache。如果你按常规方式导入functools，那么你就必须像这样调用lru_cache：functools.lru_cache(*args)
from os import * # 你可能定义了一个与导入模块中名称相同的变量或函数，这时如果你试图使用os模块中的同名变量或函数，实际使用的将是你自己定义的内容
from os import path, walk, unlink
```

- 可选导入

```python
try:
    # For Python 3
    from http.client import responses
except ImportError:  # For Python 2.5-2.7
    try:
        from httplib import responses  # NOQA
    except ImportError:  # For Python 2.4
        from BaseHTTPServer import BaseHTTPRequestHandler as _BHRH
        responses = dict([(k, v[0]) for k, v in _BHRH.responses.items()])
```

- 局部导入

```python
def square_root(a):
    # This import is into the square_root functions local scope
    import math
    return math.sqrt(a)
```

- 导入常见错误

  - 循环导入：a导入b，b导入a
  - 当你创建的模块与标准库中的模块同名时，如果你导入这个模块，就会出现覆盖导入，也即去导入我们的模块，而非标准库模块。例如：创建一个math.py

  ```python
  import math
  
  def square_root(number):
      return math.sqrt(number)
  
  square_root(72)
  ```

  现在打开终端，试着运行这个文件，你会得到以下回溯信息（traceback）：

  ```shell
  Traceback (most recent call last):
    File "math.py", line 1, in <module>
      import math
    File "/Users/michael/Desktop/math.py", line 6, in <module>
      square_root(72)
    File "/Users/michael/Desktop/math.py", line 4, in square_root
      return math.sqrt(number)
  AttributeError: module 'math' has no attribute 'sqrt'
  ```

  这到底是怎么回事？其实，你运行这个文件的时候，Python解释器首先在当前运行脚本所处的的文件夹中查找名叫`math`的模块。在这个例子中，解释器找到了我们正在执行的模块，试图导入它。但是我们的模块中并没有叫`sqrt`的函数或属性，所以就抛出了`AttributeError`。

# 5. def 函数

- <https://www.tutorialspoint.com/python/python_functions.htm>

1. def关键在开头，后面是函数名和`():`, `()`里面可以有参数，后面的`:`标识 函数代码块开始，代码块需要对齐。
2. 函数的第一行是一个可选的对函数的说明字符串。
3. return语句可以有一个返回值，没有返回值的与`return None`效果相同。

```python
def functionname( parameters ):
   "function_docstring"
   function_suite
   return [expression]
```

## passed by reference

4. **All parameters (arguments)** in python are **passed by reference**. 在函数里面对参数进行修改，在函数外面可见。
- passed by reference也passed了，只是参数内容被指向原地址，变量地址改变了。所以下面这种对参数重新赋值的情况在函数外并不可见：

```python
#!/usr/bin/python

# Function definition is here
def changeme( mylist ):
   "This changes a passed list into this function"
   mylist = [1,2,3,4]; # This would assig new reference in mylist
   print "Values inside the function: ", mylist
   return

# Now you can call changeme function
mylist = [10,20,30];
changeme( mylist );
print "Values outside the function: ", mylist
```

输出：

```shell
Values inside the function:  [1, 2, 3, 4]
Values outside the function:  [10, 20, 30]
```

## 函数的参数

1. 函数的参数

- 正常传参：按照参数顺序，依次给出值
- 关键字传参：`paramName=paramValue`的方式传参
- 参数默认值：

```python
#!/usr/bin/python

# Function definition is here
def printinfo( name, age = 35 ):
   "This prints a passed info into this function"
   print "Name: ", name
   print "Age ", age
   return;

# Now you can call printinfo function
printinfo( age=50, name="miki" )
printinfo( name="miki" )
```

- 变长参数

  - 格式

  ```python
  def functionname([formal_args,] *var_args_tuple ):
     "function_docstring"
     function_suite
     return [expression]
  ```

  - 例子

  ```python
  #!/usr/bin/python
  
  # Function definition is here
  def printinfo( arg1, *vartuple ):
     "This prints a variable passed arguments"
     print "Output is: "
     print arg1
     for var in vartuple:
        print var
     return;
  
  # Now you can call printinfo function
  printinfo( 10 )
  printinfo( 70, 60, 50 )
  ```

## 匿名函数：lambda表达式

1. 匿名函数

- 不是以`def`定义的，而是以`lambda`定义的
- 格式：`lambda [arg1 [,arg2,.....argn]]:expression`， 如：

```python
#!/usr/bin/python

# Function definition is here
sum = lambda arg1, arg2: arg1 + arg2;

# Now you can call sum as a function
print "Value of total : ", sum( 10, 20 )
print "Value of total : ", sum( 20, 20 )
```

- 有点儿类似于c中的inline，参数个数任意，但是返回值只有一个（TODO：看起来更像是只有一个返回值表达式）。不能包含其他命令或者多个表达式。
- lambda表达式只能使用参数列表和the global namespace中的参数。

## 变量作用域

1. 变量作用域

- 有global和local两种作用域，global在函数定义之外，local的在函数定义之内。global变量能被所有函数使用
- 似乎local优先级更高，如下：

```python
#!/usr/bin/python

total = 0; # This is global variable.
# Function definition is here
def sum( arg1, arg2 ):
   # Add both the parameters and return them."
   total = arg1 + arg2; # Here total is local variable.
   print "Inside the function local total : ", total
   return total;

# Now you can call sum function
sum( 10, 20 );
print "Outside the function global total : ", total 
```

- 那么有声明变量的说法么，或者说，上面那里，如果要更改global的total，该怎么做。TODO：？

# 6. 进程初始化

- <https://pythonhow.com/how-a-flask-app-works/>

1. `APP = Flask(__name__)`

- create an instance of the *Flask* class for our web app
- `__name__ `is a special variable that gets as value the string `"__main__"` when you’re executing the script

2. 常见

```python
if __name__ == '__main__':
    app.run(debug=True)
```

