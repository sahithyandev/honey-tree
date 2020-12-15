import { CommandMetaInfo, ExtendedCommand } from "./../types";

export class Init extends ExtendedCommand {
  static meta: CommandMetaInfo = {
    commandName: "init",
  };
  static args = [
    {
      name: "template-name",
      description: "Name of the template",
      required: true,
    },
  ];
  static description = "Initialize a honey-tree template project";
  static requiredArgsCount = Init.args.filter((argObj) => argObj.required)
    .length;

  async run() {
    if (this.argv.length < Init.requiredArgsCount) {
      this.error(new Error("Required arguments not given"));
    }
  }
}
