# Installation

    git clone https://github.com/bestpilotingalaxy/fb-sequence.git
#
    
    docker-compose up

# Project layout
![image](https://user-images.githubusercontent.com/59182467/129530759-976a3279-fcf8-44b7-95c5-4d9373d9139a.png)



# Debug

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
