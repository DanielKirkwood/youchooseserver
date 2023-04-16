#!/usr/bin/env bash
# Script installs [TASK](https://taskfile.dev) which is a Make alternative.

# bool function to test if the user is root or not
is_user_root () { [ ${EUID:-$(id -u)} -eq 0 ]; }

TASK_PATH=$(which task)
if [ -z "$TASK_PATH" ]
then
  if is_user_root;
  then
    sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b ~/.local/bin
    echo "task binary added to \$PATH"
  else
    echo "you need to be a sudo to add the binary to \$PATH"
  fi
else
  echo "Task has already been installed"
fi
