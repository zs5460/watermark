@echo off
if exist watermark.zip (del watermark.zip)
if exist watermark.exe (del watermark.exe)
go build
upx watermark.exe
"c:\program files\7-zip\7z.exe" a watermark.zip watermark.exe watermark.png readme.txt sky.jpg
