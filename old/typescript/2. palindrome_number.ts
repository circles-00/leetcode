const isPalindrome = (x: number): boolean => {
  return x.toString() === `${x}`.split("").reverse().join("");
};
