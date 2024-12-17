# Jy's tinygo keyboard example

This code allows a microcontroller like the raspberry pi pico act as a keyboard or MIDI controller

To use this code be sure to have [go](https://go.dev/doc/install) and [tinygo](https://tinygo.org/getting-started/install/) installed

## Status

The examples work 

## Examples
In the example folder you will find code for a basic gpio keyboard

To flash this code to you pico you need to run 
```sh 
tinygo flash -target=pico gpio_keyboard.go
```

## actually use this module

Create a new go project and add it to it,
```sh
go mod init
go get github.com/JyJy32/tinygo-keyboard
```
or copy the files into your project 
