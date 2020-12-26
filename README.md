# go-benq-remote

It's a small go program which allows you to control some basic aspects of your benq beamer(Tested only with Benq TH682ST).

Currently supported:
* Power off/Power on
* Volume Up
* Volume Down
* Brightness Down
* Brightness Up

# How to run

Compile it and run it with: 
```
packr2
go build
./go-benq-remote /dev/ttyUSB0 115200 ":4001"
```

First parameter is your serial interface, second baudrate and last one the listen address of your webserver.

Attention: In the above configuration the server will listen on all interfaces! No auth supported

# Used ressources

* https://business-display.benq.com/content/dam/bb/en/product/projector/interactive-room-projector/lw890ust/quick-start-guide/lw890ust-rs232-control-guide-0-windows7-windows8-winxp.pdf