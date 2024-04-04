![logo](img/logo-entropy.jpeg)

Induce irregularity into your k8s cluster to test resilience and integrity.

## Set-up

### Clone Repository
- Clone with ```git clone https://github.com/hari166/entropy```
- ```cd entropy```
- Run ```./entropy [command]```

### Docker
- Coming soon



> [!WARNING] 
> Proceed only with adequate testing and rollback procedures. Do not deploy in production environment.

## Usage
> [!NOTE]
> Requires glibc v2.35 and above.
 ```inject``` command is work in progress
> 
Usage:
```
  entropy [command]
```

Available Commands:
```
    artifact    Experiment with config maps and secrets
    cord        Cordon a node
    help        Help about any command
    inject      Exit shell script with status code 1
    killRandom  Kill a random pod
    scale       Scale a deployment
    service     Terminate a service
```
### Help
Use ```entropy [command] --help``` to know more about a particular command.

### Example
You can terminate a random pod with:
```entropy  killRandom --ns NAME_SPACE``` 

## Documentation
 - [client-go](https://pkg.go.dev/k8s.io/client-go/kubernetes)
 - [cobra-cli](https://cobra.dev/)
 - [Golang Official Documentation](https://go.dev/doc/)