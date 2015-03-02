# aws-ssh

Generate an SSH Inventory Config for Amazon Web Services

# Installation

Running this simple application will produce entries for your SSH config file for all running instances on EC2 for the given tag. This is easily piped into the SSH config file directly. When combined with auto-completion and a cron-entry you can auto-complete SSH for your instances with ease.

`$ ./aws-ssh --private=true --strict=false --user=username > ~/.ssh/config`

### OSX Users

For auto-complete (tab) to work you will need to add a line similar to this to your /etc/bashrc file.

`$ complete -o default -o nospace -W "$(grep -i -e '^host ' ~/.ssh/config | awk '{print substr($0, index($0,$2))}' ORS=' ')" ssh scp sftp`
