/*请你编写一个函数，检查给定的值是否是给定类或超类的实例。
可以传递给函数的数据类型没有限制。例如，值或类可能是  undefined */
var checkIfInstanceOf = function (obj, classFunction) {
  //特殊情况obj的值为null或者undefined，或者classFunction不为函数，直接返回false
  if (
    obj === null ||
    obj === undefined ||
    typeof classFunction !== "function"
  ) {
    return false;
  }

  //原始链起点
  let prototype = Object.getPrototypeOf(obj);
  //遍历原始链
  while (prototype !== null) {
    if (prototype === classFunction.prototype) {
      return true;
    }
    prototype = Object.getPrototypeOf(prototype);
  }
  //遍历完仍没有匹配，则直接返回false
  return false;
};
