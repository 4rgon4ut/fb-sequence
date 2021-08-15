
# Debug

I'm personally use VSCode with dlv-dap debugger
* https://github.com/golang/vscode-go/blob/master/docs/dlv-dap.md


### launch.json
```json
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch Package",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "debugAdapter": "dlv-dap",
            "trace": "verbose",
            "program": "${workspaceFolder}/cmd/server/",
            "envFile": "${workspaceFolder}/.env"
        }
    ]
}
```
