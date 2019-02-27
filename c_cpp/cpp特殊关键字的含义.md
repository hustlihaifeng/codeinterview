

|关键字|含义|连接|开始版本|
|-|-|-|-|
|explicit|函数声明前面加上explicit，表示阻止隐式类型转换。普通类型的隐式类型转换类似与c，用户自定义的类型中，单参数构造函数可以进行隐式类型转换，包括只需要指定一个参数的的构造函数（有可能是多参数，但是其他的参数有默认值）|[https://en.cppreference.com/w/cpp/language/explicit](https://en.cppreference.com/w/cpp/language/explicit) ||
|volatile|volatile 变量每次需要时都会从其地址读取最新值，告诉编译器不要优化器访问，因为该值可能会被一些编译器所未知的行为修改|[https://en.cppreference.com/w/cpp/language/cv](https://en.cppreference.com/w/cpp/language/cv)||
|noexpect|可以将一个永远不会抛出异常的函数声明成 noexpect|||
|std::initializer_list|std::initializer_list是一种标准库类型，编译器可以识别：当使用{}列表初始化对象时，编译器会自动创建一个std::initializer_list对象并将其提供给程序。|[如果类是一个容器，给它一个初始值列表构造函数](https://note.youdao.com/web/#/file/WEBcd302329e9e7068fb57532daae71e195/note/WEBcf37c75451aec4a745de83ca00520c47/)||
|virtual|虚函数|运行时多态，基类提供默认实现。编译器常见的做法是**将虚函数的名字转换成函数指针表中对应的索引值**，然后每个类有个虚函数表，有虚函数的类的对象有个虚函数表指针，编译时决定这个虚函数表指针绑定基类还是子类，调用时才去这个指针指向的虚表取得真正的虚函数地址。[类对象的大小=各非静态数据成员（包括父类的非静态数据成员但都不包括所有的成员函数）的总和+ vfptr指针(多继承下可能不止一个)+vbptr指针(多继承下可能不止一个)+编译器额外增加的字节。](https://blog.csdn.net/zzwdkxx/article/details/53635173)||
|virtual...=0|纯虚函数，抽象类，不能实例化，派生类必须实现这个函数|||
|override|显示覆盖父类虚函数|在类函数声明后面加override就显式的指明该函数必须覆盖某父类的同名函数（条件：同名同参数，函数在父类中有virtual关键字），防止程序员不小心写错函数名。||
|dynamic_cast|从父类向子类转换|如果想调用某个派生类里面的特有的函数，则可以使用dynamic_cast运算符，如果dynamic_cast的参数既不是期望类也不是期望类的派生类，则dynamic_cast返回**nullptr**||
|std::auto_ptr |独有型智能指针，ownership实现，违反会在执行时coredump|智能指针应该避免删除非堆内存(因为会自动释放，这样只能指针析构函数里面的delete就会出错)。容器算法禁止使用auto_ptr。[浅谈智能指针auto_ptr/shared_ptr/unique_ptr](https://blog.csdn.net/derkampf/article/details/72654883)|c++98提出，c++11抛弃|
|std::unique_ptr|独有型只能指针，ownership实现，违反会在编译时出错|容器算法允许使用auto_ptr。[unique_ptr不是不允许赋值，它允许源unique_ptr是个临时右值](https://blog.csdn.net/derkampf/article/details/72654883) std::move()可以将一个智能指针对象赋给另一个.[浅谈智能指针auto_ptr/shared_ptr/unique_ptr](https://blog.csdn.net/derkampf/article/details/72654883)|c++11新增。`g++ xx.cpp -o xx -std=c++11`|
|std::shared_ptr|共享型只能指针，reference counting(引用计数)实现|[浅谈智能指针auto_ptr/shared_ptr/unique_ptr](https://blog.csdn.net/derkampf/article/details/72654883)|c++11新增. `g++ xx.cpp -o xx -std=c++11`|
|&&|右值引用|&&指“右值引用”，右值是我们无法为其赋值的值（一般是临时对象），与左值相对。要尽量保证移动构造函数 不发生异常，可以通过noexcept关键字。右值引用不能绑定左值：int a; int &&c = a;   这样是不行的。|c++11新增|
|=default|程序员只需在函数声明后加上“=default;”，就可将该函数声明为 "=default"函数，**编译器将为显式声明的 "=default"函数自动生成函数体,提高代码的执行效率**|[c++11 类默认函数的控制："=default" 和 "=delete"函数](https://www.cnblogs.com/lsgxeva/p/7787438.html)|C++11新增|
|=delete|程序员只需在函数声明后上“=delete;”，就可将该函数禁用|[c++11 类默认函数的控制："=default" 和 "=delete"函数](https://www.cnblogs.com/lsgxeva/p/7787438.html)|c++11新增|

