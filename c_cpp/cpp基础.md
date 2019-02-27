[TOC]

# 1 函数类型

1. **返回值类型和实参类型组成函数类型，对于类成员函数来说，类名字本身也是函数类型的一部分。**
eg：
```cpp
double get(const vector<double>& vec,int index);//函数类型：double (const vector<double>&,int)
char& String::operator[](int index);//函数类型：char& String::(int)
```
2. 同名函数
- 若名同参数不同，则编译器负责为每次调用选择匹配度最高的函数。
- 若存在两个可供选择函数并且他们难分优劣，则编译器报错。eg：
```cpp
print(int,double);
print(double,int);
```
调用：print(0,0);
# 2 常用算数类型转换
**表达式中使用的类型转换称为常用类型转换**，目的在于确保表达式以它的运算对象中最高的精度进行求值运算。

# 3 初始化
1. =和{}都是c++的初始化符号。建议多使用{}。=是为了和c兼容。eg：
```cpp
double d1=2.3;
double d2 {2.3};
complex<double> z=1;
complex<doube> z2 {d1,d2};
complex<double> z3 = {1,2}//当使用{}时，通常省略=
vector<int> v {1,2,3,4,5,6};
```
2. 不同：{}不允许发生可能导致信息丢失的类型装换。eg：
```cpp
int i1=7.2；
int i2 {7.2};//错误，试图执行浮点数到整数的类型转换
int i3={7.2}//错误，这里的=是多余的，同上面。
```
3. 在定义一个变量时，若类型可以有初始化符号推导出，则可以用auto代替具体类型，这个在泛型编程中有用。eg：
```cpp
auto i1=123;//int
```
注：**指定了类型时，尽量用{}来初始化；使用了auto时，尽量要=。**
# 4 作用域和生命周期
1. 作用域
- 局部变量的作用域是声明它的块，块的边界用{}表示。
- 命名空间作用域：如果一个名字定义在命名空间内部，同时位于任何函数、lambda、类和enum class的外部，则称改名字为命名空间成员名字。作用域从声明它的空间开始，到命名空间结束为止。
2. 销毁
- 命名空间对象，销毁点在整个程序的末尾。
- 成员对象销毁点依赖于它所属对象的销毁点。
- **用new创建的对象一致存活到delete销毁它为止**。
# 5 常量
1. c++支持两种不变形概念：
- const：大意是“我承诺不改变这个值”，编译器负责确认并执行const的承诺。不管后面赋值声明。
- constexpr：大意是“**在编译时就求值**”（一般的变量是在运行时求值），主要用于说明常量（后面必须接常量），作用是允许把数据置于只读内存中以提升性能。
- eg:
```cpp
const int dmv=17; //一个命名常量
int var=17;	//非常量

constexpr double max1=1.4*square(dmv);	//如果square(17)是常量表达式，则正确
constexpr double max2=1.4*square(var);	//错误：var不是常量表达式。
const double max3=1.4*square(var);		//ok：可以在运行时求值。并保证求值后该值不变。
```
- 如果一个函数被用在常量表达式（constexpr）中，即该表达式在编译时就求值，则这个函数必须被声明为constexpr。
	- 要想声明成constexpr，函数必须非常简单：函数中仅有一条计算某个值得return语句。
	```cpp
	constexpr double square(double x){return x*x;}
	```
	- 也可以用非常量调用constexpr函数，但此时该函数返回值不能用于constexpr环境。
# 6 指针数组和引用
1. range-for语句（类似于go里面的for-range）（c++11新增）
```cpp
int x[]={1,2,3,4,5,6};
for(auto tmp:x){cout<<tmp<<endl;}
```
或者：
```cpp
for(auto x:{1,2,3,4,5,6}){cout<<x<<endl;}
```
当不希望将变量的值拷贝的x中时：
```cpp
for(auto& x:{1,2,3,4,5,6}){cout<<x<<endl;}
```
若既不想拷贝变量，又不想改变变量的值：
```cpp
int sum(const vector<double>&); 
```
- 注：**const修饰数组不能将数组项作为左值，但是可以使用{}或strncpy或memcpy给整个数组赋值。**
- [可以遍历的对象包括](https://www.cnblogs.com/h46incon/archive/2013/06/02/3113737.html)：
  - 数组。（不包括指针）
  - 定义了begin()和end()方法，且返回该方法返回迭代器的类对象。（STL 中所有容器都可以）

2. nullptr
空指针，`while(nullptr)`为false。
# 7 switch case break
类似于c：
```cpp
switch(a){
	case 1:break;
	case 2:break;
	default:;
}
```