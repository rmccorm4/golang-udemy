# Container From Scratch

## Mounting a File System for the Container

`container.go` attempts to `chroot` into a file system for the container to live inside.
To do this, we will need a filesystem. Instructions below:

### Creating a File System with LXC (untested)

1. Install LXC

    ```bash
    # On Ubuntu, for example
    sudo apt-get install lxc
    ```

2. Create a container with `lxc-create` and use this filesystem

    ```bash
    lxc-create -t ubuntu -n my-container
    # Now use /var/cache/lxc/my-container/rootfs-amd64 as the directory to `chroot` into in `container.go`
    ```

### Creating a (Fake) File System 

If you're already inside an **unprivileged** container (Like running Linux on ChromeOS), then you don't
have the permission to create an lxc container. So my work around was to try and re-create a minimal file
system for the purpose of this example by copying things from the root `/` directory until it worked.

Here's what worked for me:

```bash
# Create directory for your fake rootfs
mkdir fake-rootfs

# These seem to be the bare-minimum to get this example working
cp -r /bin /lib /lib64 fake-rootfs/

# To get processes unique to the container from the `ps` command
mkdir -p fake-rootfs/proc
mkdir -p fake-rootfs/sys/fs/cgroup/pids

# Test that it works - This should bring you to a minimal bash prompt
sudo chroot fake-rootfs
```

## Running the Container

```bash
# go run container.go <cmd> <args>
go run container.go run echo "hello container-world"
go run container.go run /bin/bash
```

If you get a permission error, you can try this as well:

```bash
go build container.go
# sudo ./container <cmd> <args>
sudo ./container run echo "hello container-world"
sudo ./container run /bin/bash
```