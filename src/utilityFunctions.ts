export function cleanArr<Q = any>(arr: Q[]) {
  const falseValues = [0, "", false, null, undefined];
  // @ts-ignore
  return arr.filter(item => !falseValues.includes(item));
}
