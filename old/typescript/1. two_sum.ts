const twoSum = (numbers: number[], target: number): number[] => {
  const previousValues: Map<number, number> = new Map<number, number>();
  for (let i = 0; i < numbers.length; i++) {
    const wantedNumber = target - numbers[i];
    if (previousValues.has(wantedNumber)) {
      return [previousValues.get(wantedNumber) as number, i];
    }
    previousValues.set(numbers[i], i);
  }

  return [];
};
