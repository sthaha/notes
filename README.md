## installation

### Create a conda env


```
cd $(git rev-parse --show-toplevel)
conda create -n notes python=3.6
source activate notes
```

### Update env

```
conda env update
```

### Install golang kernel

Follow documentation of [gophernotes][]

```sh

mkdir -p go/{src,pkg,bin}
export GOPATH=$PWD/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN
go get -u github.com/gopherdata/gophernotes

mkdir -p $(jupyter --data-dir)/kernels/gophernotes

cp $GOPATH/src/github.com/gopherdata/gophernotes/kernel/* !$
```


### Running


```
cd $(git rev-parse --show-toplevel)
source activate notes

export GOPATH=$PWD/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN

jupyter notebook
```
