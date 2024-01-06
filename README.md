# Interoduction
This tool is inspair and based on [olliebarrick:](https://github.com/olliebarrick/go-k8s-portforward) project

When working on a large microservice development it is often need to port-forward so you will be able to connect to a database and other services you depend on.

However it is tedious define all those manually so this tool come to rescue.
It use a configuration including a defintions of all the ports forward you need to create in order to test you service locally 

# Gettting Start
To use the tool you will a config file which contain a defintion of all the services 

For example
````yaml
- name: httpd
  pod: httpd-deployment
  namespace: httpd
  ports:
    src: 8080
    dst: 80
- name: forward
  pod: httpd-forward
  namespace: tunnels
  ports:
    src: 8081
    dst: 80
````
This define two services you want to port forward to.
> * `pod`: the tool will look at the first service match the prefix define in the pod attrinute
> * `ports.src` is you local port
> * `ports.dst` is the remote port you want to connect

# Configuration Search
The tool looks for config file in four steps:

1. using `-f` in the command line
2. using `$PORT_FORWARD_CONFIGFILE` environment variable
3. using `pf.yaml` | `pf.yml` in the local directory
4. using `pf.yaml` | `pf.yml` in ~/.config/port-forward

# Build
Run make build to create the binary file
````bash
make build
````
Binary file create under `./bin` directory

# Usage
if `-n <some-service-to-forward-to` the tool will port foward just to that endpoint.
otherwise it will read the all the entries in the file and create the appropriate port forwards

Command line argumants
* pf -h
* pf -n 'prefix-pod-name' -f 'config-file'
* pf -f 'config-file'
* pf

# Kubeconfig

By default, it will load a Kubernetes configuration file from `~/.kube/config` or `$KUBECONFIG`.

