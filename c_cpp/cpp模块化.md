# 1. 申明和实现分离 
将声明放进.h文件，和实现分离，来实现模块化。
# 2. 命名空间
- 命名可以用来避免名字冲突，明智的做法是将我们自己的代码放在自己的命名空间中，避免与标准库中的函数冲突。
- main函数属于全局命名空间，也就是说，不属于任何命名空间。
# 3. 异常
- 内部抛出，外部捕获
```cpp
double& Vector::operator[](int i){
    if(i<0 || i>=size()){
    	throw out_of_range{"Vector::operater[]"};	//注意这种用法，out_of_range像是初始化的
    }
    return elem[i];
}
void f(Vector& v){
    try{
    	v[v.size()]=7;
    }catch(out_of_range){
    	//在这里处理异常越界
    }
}
```
- 可以将一个永远不会抛出异常的函数声明成 noexpect
如：
```cpp
void user(int sz)noexpect
{
    Vector v(sz);
    iota(&v[0],&v[sz],1);//为V赋值1,2,3,4
}
```
**一旦user发生了错误，函数user还是会抛出异常，此时标准库函数terminate立即终止当前程序的执行**。
# 4. 不变式

不变式是设计类的关键，而前置条件也在设计函数的过程中起到类似的作用。不变式能：
- 帮助我们准确的理解想要什么。
- 强制我们具体而准确的描述设计，有助于代码正确。
如前置条件，检查输入参数：
```cpp
Vector::Vector(int s){
    if(s<0){//前置条件检查
    	throw length_error{};
    }
    elem=new double[s];
    sz=s;
}
void test(){
    try{
    	Vector v(-27);
    }catch(std::length_error){//std::length_error报告元素数目为非正常的错误
        cout<<"test faild:length error\n"<<endl;
        throw;//简单处理后，继续抛出异常
    }catch(std::bad_alloc){//std::bad_alloc 如果new找不到可分配的内存
    	std::terminate();//终止程序
    }
}
```
- 不变式是c++中由构造函数和析构函数支撑资源管理概念的基础
# 5. 静态断言
```cpp
static_assert(A,s);//在编译阶段判断，当常量表达式A不为true时，把s作为一条编译器错误信息输出
```
静态断言常用来在泛型编程中作为形参的类型设置断言。
# 6. 建议
- 不要在头文件中定义非内联结构
- 不要在头文件中使用using指令
- 在设计阶段就想好错误处理策略
- 用专门设计的用户自定义类型作为异常类型
