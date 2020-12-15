import { Command } from "@oclif/command";

export interface CommandMetaInfo {
  commandName: string;
}

export abstract class ExtendedCommand extends Command {
  static meta: CommandMetaInfo;
}
