## installation

### Create a conda env


```
cd $(git rev-parse --show-toplevel)
conda create -n notes python=3.6
conda activate notes
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

### Install nodejs kernel

See: https://github.com/n-riesco/ijavascript

```shell
# after adding nodejs
conda env update

npm install -g ijavascript
ijsinstall
```


### Install typescript

Followed instructions on https://github.com/nearbydelta/itypescript to install
typescript kernel.

NOTE: First I installed nodejs kernel and then the itypescript

```shell
conda activate notes
npm install -g itypescript
its --ts-install=local
```

### Running

```
conda activate notes
./start.sh
```
Also note that the script exits if it fails to detect that he active conda env isn't `notes`

## Theme

The one I use is

```
jt -t onedork -fs 95 -altp -tfs 11 -nfs 115 -cellw 88% -T -vim
```
