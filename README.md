# Alice-LG - Your friendly looking glass
__"No, no! The adventures first, explanations take such a dreadful time."__  
_Lewis Carroll, Alice's Adventures in Wonderland & Through the Looking-Glass_

Take a look at an Alice-LG production example:
- https://lg.ecix.net/
- https://lg.ecix.net/api/config
- https://lg.ecix.net/api/routeservers
- https://lg.ecix.net/api/routeservers/0/status
- https://lg.ecix.net/api/routeservers/0/neighbours
- https://lg.ecix.net/api/routeservers/0/neighbours/ID109_AS31078/routes
- https://lg.ecix.net/api/routeservers/0/lookup/prefix?q=217.115.15.0

# Explanations
Alice-LG is a BGP looking glass which gets its data from external APIs.

Currently Alice-LG supports the following APIs:
- [birdwatcher API](https://github.com/ecix/birdwatcher) for [BIRD](http://bird.network.cz/)

Normally you would first install the [birdwatcher API](https://github.com/ecix/birdwatcher) directly on the machine(s) where you run [BIRD](http://bird.network.cz/) on
and then install Alice-LG on a seperate public facing server and point her to the afore mentioned [birdwatcher API](https://github.com/ecix/birdwatcher).

This project was a direct result of the [RIPE IXP Tools Hackathon](https://atlas.ripe.net/hackathon/ixp-tools/) 
just prior to [RIPE73](https://ripe73.ripe.net/) in Madrid, Spain.

Major thanks to Barry O'Donovan who built the original [INEX Bird's Eye](https://github.com/inex/birdseye) BIRD API of which Alice-LG is a spinnoff

## Building Alice-LG from scratch
__These examples include setting up your Go environment, if you already have set that up then you can obviously skip that__

### CentOS 7:
First add the following lines at the end of your ~/.bash_profile:
```
GOPATH=$HOME/go
export GOPATH
PATH=$PATH:$GOPATH/bin
export PATH
```
Now run:
```
source ~/.bash_profile
sudo yum install golang npm
sudo npm install --global gulp-cli
go get github.com/GeertJohan/go.rice
go get github.com/GeertJohan/go.rice/rice
mkdir -p ~/go/bin ~/go/pkg ~/go/src
cd ~/go/src/
git clone git@github.com:ecix/alice-lg.git
cd alice-lg/client
make
cd ..
make
```
Your Alice_LG source will now be located at `~/go/src/alice-lg` and your alice-LG executable should be at `~/go/src/alice-lg/bin/alice-lg-linux-amd64`

## Configuration

An example configuration can be found at 
[etc/alicelg/alice.example.conf](https://github.com/ecix/alice-lg/blob/readme_update/etc/alicelg/alice.example.conf).

You can copy it to any of the following locations:
```
    etc/alicelg/alice.conf        # local
    etc/alicelg/alice.local.conf  # local
    /etc/alicelg/alice.conf       # global
```

You will have to edit the configuration file as you need to point Alice-LG to the correct [APIs](https://github.com/ecix/birdwatcher):
```
    [source.0]
    name = rs1.example.com (IPv4)
    [source.0.birdwatcher]
    api = http://rs1.example.com:29184/
    # show_last_reboot = true
    # timezone = UTC

    [source.1]
    name = rs1.example.com (IPv6)
    [source.1.birdwatcher]
    api = http://rs1.example.com:29186/
```

## Running

Launch the server by running

    ./bin/alice-lg-linux-amd64


## Deployment

We added a `Makefile` for packaging Alice as an RPM using [fpm](https://github.com/jordansissel/fpm).

If you have all tools available locally, you can just type:

    make rpm

If you want to build the package on a remote machine, just use

    make remote_rpm BUILD_SERVER=my-rpm-building-server.example.com

which will copy the dist to the remote server and executes fpm via ssh.

You can specify which system integration to use:
Set the `SYSTEM_INIT` variable to `upstart` or `systemd` (default)
prior to building the RPM.

    make remote_rpm BUILD_SERVER=rpmbuild.example.com SYSTEM_INIT=upstart



## Hacking

The client is a Single Page React Application.
All sources are available in `client/`. 

Install build tools as needed:

    npm install -g gulp-cli


Create a fresh UI build with

    cd client/
    make client

This will install all dependencies and run `gulp`.

While working on the UI you might want to use `make watch`,
which will keep the `gulp watch` task up and running.



