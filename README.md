## ytask
CLI todo manager

## Usage

Download the binary in the repo [here](ytask?raw=true) and place it in your `$PATH`

If you have Go installed, you can clone the repo in your GOPATH and build the binary from the source code.

---

There are 3 commands that it accepts - `list` , `add` and `do`

```
$ ytask 
ytask is a CLI task manager

Usage:
  ytask [command]

Available Commands:
  add         add a task to your list
  do          mark a task completed
  help        Help about any command
  list        list all tasks

Flags:
  -h, --help   help for ytask

Use "ytask [command] --help" for more information about a command.
```


```
$ ytask list
You have no tasks to complete.
```
```
$ ytask add foo
added "foo" to the task list
$ ytask add bar
added "bar" to the task list
$ ytask add baz
added "baz" to the task list
$ ytask list
You have the following tasks:
1. foo 
2. bar 
3. baz
```

```
$ ytask do 1
marked as done task "1"
$ ytask list
You have the following tasks:
1. bar 
2. baz 
$ ytask do 1 2
marked as done task "1"
marked as done task "2"
$ ytask list
You have no tasks to complete.
```

## Dependencies

ytask uses [boltdb](https://github.com/boltdb/bolt) for on-disk database and [cobra](https://github.com/spf13/cobra) for CLI interactions.
