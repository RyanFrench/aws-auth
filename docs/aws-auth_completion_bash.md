## aws-auth completion bash

generate the autocompletion script for bash

### Synopsis


Generate the autocompletion script for the bash shell.

This script depends on the 'bash-completion' package.
If it is not installed already, you can install it via your OS's package manager.

To load completions in your current shell session:
$ source <(aws-auth completion bash)

To load completions for every new session, execute once:
Linux:
  $ aws-auth completion bash > /etc/bash_completion.d/aws-auth
MacOS:
  $ aws-auth completion bash > /usr/local/etc/bash_completion.d/aws-auth

You will need to start a new shell for this setup to take effect.
  

```
aws-auth completion bash
```

### Options

```
  -h, --help              help for bash
      --no-descriptions   disable completion descriptions
```

### SEE ALSO

* [aws-auth completion](aws-auth_completion.md)	 - generate the autocompletion script for the specified shell

###### Auto generated by spf13/cobra on 30-Nov-2021