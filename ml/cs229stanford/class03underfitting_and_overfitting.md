1. underfitting：欠拟合（没有很好的反应样本数据的情况） overfitting：过拟合（一般仅仅精确的反映了样本数据的情况）
2. parametric learning algorithm has a fixed, finite number of parameters (the θi’s), which are fit to the data. The term “non-parametric” (roughly) refers to the fact that the amount of stuff **we need to keep** grows linearly with the size of the training set in order to represent the hypothesis h.
3. Local Weight Regression：不需要我们很关心参数的选择。但是从公式来看，只是不是很关心测试集的数据量，w(y-θx), θ要和x相乘，其维度必须和x匹配。（θ：theta）
4. 常见数学符号输入（实际上是希腊字母）：无法拼出来时，搜狗输入法特殊符号的希腊字母里面有全部的希腊字母。

- θ：theta
- τ：tao
- σ：sigma
- α：alpha
- β：beta
- γ：gamma
- δ：delta
- κ：kappa
- λ：lamda
- π：pi
- ρ：rou
- φ：fai
- ω：omega

5. sigmoid function/logictic function：1/(1+e^(-z))，z很大的时候是1，z很小的时候是0.
6. 让学生下课看着讲义把推导弄懂，然后不看讲义，自己进行推导。