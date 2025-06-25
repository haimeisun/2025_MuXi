/*题目：请你编写一个函数，它接收一个函数数组 [f1, f2, f3，…， fn] ，并返回一个新的函数 fn ，它是函数数组的 复合函数 。

[f(x)， g(x)， h(x)] 的 复合函数 为 fn(x) = f(g(h(x))) 。

一个空函数列表的 复合函数 是 恒等函数 f(x) = x 。

你可以假设数组中的每个函数接受一个整型参数作为输入，并返回一个整型作为输出。*/
var compose = function (functions) {
  //特殊情况：如果给的函数长度为0（即空函数），返回x就行
  if (functions.length === 0) {
    return (x) => x;
  }
  return function (x) {
    let result = x;
    //一层一层往回算，先算最外层即functions[n-1]
    for (let i = functions.length - 1; i >= 0; i--) {
      const fuc = functions[i];
      result = fuc(result);
    }
    return result;
  };
};
