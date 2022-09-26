# dungeons

## In progress tasks

* Adding multithreads in fuzzer command

# Problems with max open files (to much sockets opened)

```sh
ulimit -n 1000000
```

# GO MAX PROCESSORS (by deault is 1)

Is necessary?
```sh
export GOMAXPROCS=2
```

# Tunning sysctl parameters

https://rtcamp.com/tutorials/linux/sysctl-conf/  
http://www.linux-admins.net/2010/09/linux-tcp-tuning.html