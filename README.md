# youchooseserver

## Tools

This repo uses mutiple tools that should be installed.

Firstly, Task is used as a Make replacement. Install it with:

~~~
sudo ./scripts/install-task.sh
~~~

This installs the `task` to `/usr/local/bin/task` so `sudo` is needed.

> On macOS and Windows, `~/.local/bin` and `~/bin` are not added to $PATH by default.

`Task` tasks are defined inside `Taskfile.yml` file. A list of tasks availible can be views with:

~~~
task -l #or
task list
~~~
