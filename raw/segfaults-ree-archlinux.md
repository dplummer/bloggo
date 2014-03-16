If you're like me and develop on archlinux, but your Rails apps use REE
still, you may have had some problems recently. Specifically, GCC 4.7
breaks REE, producing segfaults when you want to run code.

I recently realized the solution, which is pretty simple. Install GCC
4.6, then install REE with that instead.

```shell
$ sudo pacman -S gcc-4.6
$ rvm reinstall ree --with-gcc=/usr/bin/gcc-4.6
```
