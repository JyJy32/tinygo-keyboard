# Jy's tinygo keyboard example

This code allows a microcontroller like the raspberry pi pico act as a HID 

To test be sure to have [go](https://go.dev/doc/install) and [tinygo](https://tinygo.org/getting-started/install/) installed

## Status

This is currently very much a W.I.P. but you can see the intended working here  

## Examples
In the example folder you will find code for a basic gpio keyboard

to flash this code to you pico you need to run 
```sh 
tinygo flash -target=pico gpio_keyboard.go
```

## actually use this module

create a new go project and add it to it,
```sh
go mod init
go get github.com/JyJy32/macropad


```
```
```
