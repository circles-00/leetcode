const romanValues: Record<string, number> = {
  I: 1,
  V: 5,
  X: 10,
  L: 50,
  C: 100,
  D: 500,
  M: 1000,
};

const romanToInt = (s: string): number => {
  const romanDigits = s.split("");
  let sum = 0;
  for (let i = 0; i < romanDigits.length; i++) {
    switch (`${romanDigits[i]}${romanDigits[i + 1]}`) {
      case "IV":
        sum += 4;
        ++i;
        break;
      case "IX":
        sum += 9;
        ++i;
        break;
      case "XL":
        sum += 40;
        ++i;
        break;
      case "XC":
        sum += 90;
        ++i;
        break;
      case "CD":
        sum += 400;
        ++i;
        break;
      case "CM":
        sum += 900;
        ++i;
        break;
      default:
        sum += romanValues[romanDigits[i]];
        break;
    }
  }

  return sum;
};
