# Overview

Android do not have the exact total download number. We have to download the csv and count them

This program is written by Go

# How to run

Please move your report csv file to the same location with .go file

Then name it `all-countries.csv`

|----sum_downloads.go
|----sum_downloads
|----all-countries.csv

Please note that the program only works with the same csv pattern to `all-countries-example.csv`

## With Go

If you already installed Go language, you can easy run this by this command:

```
go run sum_downloads.go
```

If you want to build to MacOs or Window program for the non-Go users run:

```
go build sum_downloads.go
```

## Without Go

You can run the complied program by

#### MacOS/Linux

```
./sum_downloads
```
