import * as fs from "fs";
import { execFile } from "child_process";
import { cleanArr } from "./utilityFunctions";

interface OptionalArgumentsRunCommand {
  args?: string[];
}

export class Git {
  runCommand(
    commandName = "",
    optional: OptionalArgumentsRunCommand = {}
  ): Promise<string> {
    optional.args = optional.args || [];

    const _args = cleanArr([commandName].concat(optional.args));

    return new Promise((resolve, reject) => {
      execFile("git", _args, (error, stdout, stderr) => {
        if (error) {
          reject(stderr);
        }
        resolve(stdout || stderr);
      });
    });
  }

  initGitRepo(directoryPath = "."): Promise<string> {
    console.log(directoryPath);
    return new Promise((resolve, reject) => {
      this.runCommand("init", { args: [directoryPath] })
        .then(msg => {
          console.log(msg);
          resolve("done");
        })
        .catch(reject);
    });
  }

  isAGitRepo(directoryPath = "."): Promise<any> {
    return new Promise((resolve, reject) => {
      this.runCommand("rev-parse", {
        args: ["-C", directoryPath],
      })
        .then(() => {
          resolve(true);
        })
        .catch((error: string) => {
          if (error.includes("not a git repo")) {
            resolve(false);
          } else reject(error);
        });
    });
  }

  cloneRepo(gitRepoLink: string, projectDir: string): Promise<string> {
    return new Promise((resolve, reject) => {
      this.runCommand("clone", { args: [gitRepoLink, projectDir] })
        .then(msg => {
          if (msg.includes("done")) {
            // means cloning finished successfully
            resolve("done");
          }
        })
        .catch(error => {
          reject(error);
        });
    });
  }

  /**
   * @description [DANGER] Resets the git repository by deleting the '.git' folder
   */
  resetGitRepository(directoryPath = "."): Promise<string> {
    return new Promise((resolve, reject) => {
      // delete directoryPath/.git

      try {
        // @ts-ignore
        fs.rmdirSync(`${directoryPath}/.git`, { recursive: true });
      } catch (error) {
        reject(error);
      }

      // reinitialize git repo
      this.initGitRepo(directoryPath)
        .then(() => {
          resolve("done");
        })
        .catch(error => {
          reject(error);
        });

      resolve("done");
    });
  }
}
