import * as fsExtra from "fs-extra";
import { config } from "./settings";

export const isHoneyTreeBoilerplateTemplate = (
  directoryPath: string
): Promise<boolean> => {
  return new Promise((resolve, reject) => {
    fsExtra.stat(`${directoryPath}/${config.configFileName}`, err => {
      if (err === null) resolve(true);
      reject(err);
    });
  });
};

export function cleanArr<Q = any>(arr: Q[]) {
  const falseValues = [0, "", false, null, undefined];
  // @ts-ignore
  return arr.filter(item => !falseValues.includes(item));
}
