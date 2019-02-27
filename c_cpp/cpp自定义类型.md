[TOC]

# 1 内置类型

由基本类型（bool、char、int、unsigned、double）、const修饰符和声明运算符（&、`*`、[]）构造出来的类型成为**内置类型**。
# 2 new
new运算符从**堆**（heap，或者动态内存：dynamic memory）区域中分配内存，**堆里面分配的内存独立于它所在的作用域，会一直存活到delete运算符销毁它为止**。
# 3 struct
```cpp
struct Vector{
	int size;
	double* elem;
};
Vector v;//注意，不同于c，这里不需要再加struct了
Vector& rv=v;
Vector* pv=&v;
v.size;
rv.size;
pv->size;
```
# 4 class
```cpp
class Vector{
public:
    Vector(int s):elem{new double[s]},sz{s} {}//使用初始值列表来初始化Vector的成员
    double& operator[](int i){return elem[i];}//注意这里返回引用，但并没有在return的时候有其他的说明。
    int size(){return sz;}
private:
    int sz;
    double* elem;
};
```
注意：**c++中struct和class没有本质区别，唯一不同的是struct的成员默认是public的，也可以为struct定义构造函数和其他成员函数。**
# 5 union
union是特殊的struct，只不过所有的成员被分配在同一块内存区域。c++规定由程序员而非编译器追踪union中实际存储的值。
```cpp
enmu{str,num};
union Value{
    int i;
    char* s;
};
struct Entry{
    Type t;
    Value v;
};
void f(Entry* p){
    if(p->t==str){//通过这个来追踪
    	//p->v.s
    }
}
```
基于这种联合标记（tagged union）的使用比裸的union常见。
# 6 enum
## enum class
```cpp
enum class Color{red,blue,green};
enum class Traffic_light{green,yellow,red};
Color col=Color::red;
Traffic_light light=Traffic_light::green;
```
1. enum class的作用于在该enum class内部，所以不能混用不同的enum class或整数。
如下面的都错：
```cpp
Color col=Traffic_light::red;//错误
Color col=1;//错误
```
2. **默认情况下，enum class只定义了赋值、初始化和比较（如== <）**，然而，既然枚举类型是一种用户自定义类型，那么我们也可以为它定义别的运算符：
```cpp
Traffic_light& operator++(Traffic_light& t){
    switch(t){
        case Traffic_light::green: return t=Traffic_light::yellow;
        case Traffic_light::yellow: return t=Traffic_light::red;
        case Traffic_light::gre: return t=Traffic_light::green;
    }
}
```
## 普通的enum(避免使用)
**若想将enum值与int混用，则去掉class即可，与c一致**
```cpp
enum Color {red,yellow,green};
int col=green;
```
