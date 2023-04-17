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

## Run The Server

To run the server, we first need to start our database. A docker compose file has been created to make this simple. Run:

~~~
docker compose up postgres
~~~

to run the database.

Then we can run the server:

~~~
task run #or
task dev # for hot reloading
~~~

Even once the server has been killed, the docker container running postgres will continue to run. Stop it with:

~~~
docker compose down
~~~
