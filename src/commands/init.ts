import { CommandMetaInfo, ExtendedCommand } from "../types";
import { Git } from "../git";

export class Init extends ExtendedCommand {
  static meta: CommandMetaInfo = {
    commandName: "init",
  };

  static args = [
    {
      name: "template-url",
      description: "Name of the template",
      required: true,
    },
    {
      name: "project-dir",
      description: "Directory of the project",
      required: true,
    },
  ];

  static description =
    "Initialize a project from a honey-tree template project";

  static requiredArgsCount = Init.args.filter(argObj => argObj.required).length;

  git = new Git();

  async run() {
    const { args } = this.parse(Init);

    if (this.argv.length < Init.requiredArgsCount) {
      throw new Error("Required arguments not given");
    }

    const templateUrl = args["template-url"];
    const projectDir = args["project-dir"];

    this.log(`Initiating project @ ${projectDir} using ${templateUrl}`);

    try {
      await this.git.isAGitRepo(templateUrl);
    } catch (error) {
      console.error(`${templateUrl} is not a git repository\n`, error);
      return;
    }

    try {
      await this.git.cloneRepo(templateUrl, projectDir);
    } catch (error) {
      console.error(`Cloning failed\n ${error}`);
      return;
    }

    try {
      await this.git.resetGitRepository(projectDir);
    } catch (error) {
      console.error(`Error occured\n ${error}`);
    }
  }
}
