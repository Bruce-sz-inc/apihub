#!/bin/bash

# see https://github.com/docker/docker/issues/8791

mount_cgroups() {
  # see also https://github.com/tianon/cgroupfs-mount/blob/master/cgroupfs-mount
  if grep -v '^#' /etc/fstab | grep -q cgroup \
          || [ ! -e /proc/cgroups ] \
          || [ ! -d /sys/fs/cgroup ]; then
          return
  fi
  if ! mountpoint -q /sys/fs/cgroup; then
          mount -t tmpfs -o uid=0,gid=0,mode=0755 cgroup /sys/fs/cgroup
  fi
  (
          cd /sys/fs/cgroup
          for sys in $(awk '!/^#/ { if ($4 == 1) print $1 }' /proc/cgroups); do
                  mkdir -p $sys
                  if ! mountpoint -q $sys; then
                          if ! mount -n -t cgroup -o $sys cgroup $sys; then
                                  rmdir $sys || true
                          fi
                  fi
          done
  )
}

start_docker() {
  dockerd -g /tmp/my_graph > /tmp/docker.log 2>&1 &
}
