# linear regression(线性回归)

## 一些约定

1. m: the number of training examples
2. x: the input variables/features
3. y：output/target variable
4. (x,y): one training example
5. (x(i),y(i)): ith training example
6. h：hypathesis（假说）
7. represention of linear regretion
- ![clsss02_linear_regression_represation.png](clsss02_linear_regression_represation.png)

8. 训练的目标

- ![class02_target_of_training.png](class02_target_of_training.png)

9. partial derivative：偏导数
10. `:=`表示赋值，`=`表示等于判断

# gradient descent（梯度下降）

1. 用来最小化差异衡量函数J, 就是从初始点开始，每一步选下降最快的方向来移动，最后会得到一个局部最优解。不同初始点，得到的局部最优解不一定相同。
2. 在梯度下降的每一轮迭代中更新siouta i的方法：siouta i等于siouta i减去alpha乘以J在siouta i方向上的偏导数
3. 单个训练集时，siouta i的更方法(<https://open.163.com/movie/2008/1/B/O/M6SGF6VB4_M6SGHJ9BO.html> 34:46秒)：

- ![class02update_method_of_siouta_i_in_one_point_example_in_gradient_descent](class02update_method_of_siouta_i_in_one_point_example_in_gradient_descent.png)

# normal equation（正规方程组）

