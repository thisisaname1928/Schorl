#include <dirent.h>
#include <stdbool.h>
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

void hlt() {
  for (;;) {
  }
}

int test(int n, char **argv) {
  for (int i = 0; i < n; i++) {
    printf("%d: |%s|\n", i, argv[i]);
  }

  return 0;
}

int ls(int n, char **argv) {
  DIR *d = opendir(".");
  struct dirent *entry = readdir(d);
  while (entry != NULL) {
    if ((entry->d_type & DT_DIR) == DT_DIR)
      printf("\e[0;34m%s\e[0m\n", entry->d_name);
    else
      printf("%s\n", entry->d_name);
    entry = readdir(d);
  }

  closedir(d);
  return 0;
}

int cd(int n, char **argv) {
  if (n > 1) {
    chdir(argv[1]);
  } else {
    chdir("/");
  }

  return 0;
}

int clear(int n, char **argv) {
  printf("\033[2J\033[H");
  return 0;
}

int shellinfo(int n, char **argv) {
  printf("this is an internal Schorl init shell!\n");
  return 0;
}

int help(int n, char **argv);

typedef struct {
  int (*funcPtr)(int, char **);
  const char *cmd;
} command;

const command cmds[] = {{test, "test"},   {ls, "ls"},          {cd, "cd"},
                        {clear, "clear"}, {shellinfo, "info"}, {help, "help"}};
const uint64_t numberOfCommands = sizeof(cmds) / sizeof(command);

int help(int n, char **argv) {
  printf("Available commands: ");
  for (uint64_t i = 0; i < numberOfCommands; i++) {
    printf("%s ", cmds[i].cmd);
  }
  printf("\n");
  return 0;
}

extern DIR *currentDir;
void executeShell(const char *cmd) {
  char *command = malloc(strlen(cmd));
  if (command == NULL) {
    printf("OUT OF MEMORY!");
    hlt();
  }
  strcpy(command, cmd);

  uint64_t i = 0, n = 0;
  char **argsList = malloc(sizeof(char *));
  while (command[i] != 0) {
    // pass space
    while (command[i] == ' ' && command[i] != 0)
      i++;

    if (command[i] == 0)
      break;

    // we meet a token
    if (command[i] == '"') {
      i++;
      if (command[i] == 0)
        break;

      argsList = realloc(argsList, (n + 1) * sizeof(char *));
      n++;
      argsList[n - 1] = &command[i];

      while (command[i] != '"' && command[i] != 0)
        i++;

      if (command[i] == 0)
        break;

    } else if (command[i] != 0) {
      argsList = realloc(argsList, (n + 1) * sizeof(char *));
      n++;
      argsList[n - 1] = &command[i];

      while (command[i] != ' ' && command[i] != 0)
        i++;

      if (command[i] == 0)
        break;
    }
    // set end of token as zero
    // only do this if we meet a space, not end of string
    command[i] = 0;
    i++;
  }

  bool found = false;
  if (n > 0) {
    for (uint64_t i = 0; i < numberOfCommands; i++) {
      if (strcmp(argsList[0], cmds[i].cmd) == 0) {
        cmds[i].funcPtr(n, argsList);
        found = true;
        break;
      }
    }

    if (!found) {
      printf("command not found!\n");
    }
  }

  free(command);
}

extern int console;

void shell() {
  char buffer[1025];
  while (true) {
    int i = 0;
    buffer[0] = '\n';

    printf("shell> ");
    fflush(stdout);
    scanf("%1022[^\n]s", buffer);
    getchar();

    while (buffer[i] != '\n' && i < 1024)
      i++;

    buffer[i] = 0;

    executeShell(buffer);
  }
}