/*1、将以下函数改为箭头函数：
function multiply(a, b) {
 return a * b;
 } */
const multiply = (a, b) => a * b;

/*2、使用解构赋值,从数组[1，2，3]中取出第一个和第三个值 */
const [first, , third] = [1, 2, 3];
console.log(first);
console.log(third);

/*3、创建一个模块utils.js,导出一个函数double(n)，导入后再main.js中使用 */

/*4、定义一个类Animal,包含构造函数和一个speak方法，再定义一个子类Dog，添加一个bark方法并输出“Woof” */
class Animal {
  constructor(name) {
    name = this.name;
  }

  speak() {
    console.log("你好，我是一只 ${this.name}");
  }
}
const D = new Animal(cat);
D.speak();

class Dog extends Animal {
  constructor(name) {
    super(name);
  }

  speak() {
    console.log("你好，我是一只狗，我的名字叫${this.name}");
  }

  bark() {
    console.log("Woof!");
  }
}

const p = new Dog(billy);
p.speak();
p.bark();
