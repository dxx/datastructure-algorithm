/**
 * 递归
 * 递归是指在程序运行过程中调用本身的编程技巧
 */
function factorial(n) {
  if (n > 0) {
      return n * factorial(n - 1);
  }
  return 1;
}
function main() {
  let res = factorial(5);
  console.log(res); // 120
}

main();
