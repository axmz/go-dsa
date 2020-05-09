var maxInt = 2 ** 31 - 1
var minInt = -(2 ** 31)

var reverse = function (x) {

  let result = 0
  let pop = 0
  while (x != 0) {
    pop = x % 10
    if (result > maxInt / 10 || result == maxInt / 10 && pop > 7) {
      return 0
    }
    if (result < minInt / 10 || result == minInt / 10 && pop < 8) {
      return 0
    }
    x /= 10
    x = parseInt(x)
    result = result * 10 + pop
  }

  return result 
};

console.log("result", reverse(-1230))
