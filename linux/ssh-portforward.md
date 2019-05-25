
# Enable port forwarding
Listen on local port and forward everything on that port to some

```shell
 ssh -L <local>:<remote-ip>:<remote-port> <ssh-host>
 ```

To foward local 8000 to reddit.com:80 via ssh-host

 ```shell
 ssh -L 8000:reddit.com:80 ssh-host
 ```

 Convert that into ssh_config

```
Host reddit-via-host
  Hostname ssh-host
  LocalForward 8443 reddit.com:80
```
