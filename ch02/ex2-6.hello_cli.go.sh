# Listing 2.6  Hello World CLI: hello_cli.go

$ hello_cli
Hello World!
$ hello_cli --name Inigo
Hello Inigo!
$ hello_cli -n Inigo
Hello Inigo!
$ hello_cli -help
NAME:

USAGE:
   hello_cli [global options] command [command options] [arguments...]

VERSION:
   0.0.0

COMMANDS:
   help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --name, -n 'World'      Who to say hello to.
   --help, -h              show help
   --version, -v           print the version
