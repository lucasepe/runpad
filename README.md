# RunPad

RunPad is a tool designed to execute commands by utilizing a configuration file named `runpad.rc` located in the current working directory.

The `runpad.rc` file follows a simple structure:

- each line adheres to the format: `label: command to run`
- comments can be added by prefixing a line with the `#` symbol
- empty lines are ignored

Example `runpad.rc` content:

```text
# Create a kind cluster
kind-up: kind create cluster

# Delete the cluster
kind-down: kind delete cluster

# List all CRDs
list-crd: kubectl get crds
```

Upon parsing the `runpad.rc` file, this tool generates a window with a menu. Each menu item corresponds to a label in the configuration file. When a specific item is selected from the menu, `runpad` executes the associated command, providing a straightforward and efficient way to manage and execute predefined commands.

