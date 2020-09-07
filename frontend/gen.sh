#!/bin/bash
fyne bundle -package frontend ./resources/curve.png > bundled.go
fyne bundle -package frontend -append ./resources/logo.png >> bundled.go

