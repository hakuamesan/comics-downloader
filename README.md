# Comics Downloader

[![CircleCI](https://circleci.com/gh/Girbons/comics-downloader/tree/master.svg?style=svg)](https://circleci.com/gh/Girbons/comics-downloader/tree/master)
[![Coverage Status](https://img.shields.io/coveralls/github/Girbons/comics-downloader.svg?style=flat-square)](https://coveralls.io/github/Girbons/comics-downloader?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/Girbons/comics-downloader)](https://goreportcard.com/report/github.com/Girbons/comics-downloader)
[![Github All Releases](https://img.shields.io/github/downloads/Girbons/comics-downloader/total.svg?style=flat-square)]()
[![Release](https://img.shields.io/github/release/Girbons/comics-downloader.svg?style=flat-square)](https://github.com/Girbons/comics-downlowader/releases/latest)
[![License](https://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](LICENSE)

## Supported Sites

- http://www.comicextra.com/
- https://mangarock.com/
- https://www.mangareader.net/

## Getting Started

### Installing

Download the latest release:

- [Linux](https://github.com/Girbons/comics-downloader/releases/download/v0.7.1/comics-downloader)
- [OSX](https://github.com/Girbons/comics-downloader/releases/download/v0.7.1/comics-downloader-osx)
- [Windows](https://github.com/Girbons/comics-downloader/releases/download/v0.7.1/comics-downloader.exe)

Put the script under a folder.

## Usage

<img src="img/usage.gif?raw=true" />

### Multiple URLs

```bash
./comics-downloader -url=url,url2,url3
```

### Specify the format output

available formats:

- pdf
- epub
- cbr
- cbz

Default format is __pdf__.

example:

```bash
./comics-downloader -url=[your url] -format=epub
```

## Built With

- [go](https://github.com/golang/go)
- [gofpdf](https://github.com/jung-kurt/gofpdf)
- [go-epub](http://github.com/bmaupin/go-epub)
- [soup](https://github.com/anaskhan96/soup)
- [progressbar](https://github.com/schollz/progressbar)
- [logrus](https://github.com/sirupsen/logrus)
- [mri](https://github.com/BakeRolls/mri/blob/master/mri.go)
- [archiver](https://github.com/mholt/archiver)

## Contribuiting

Feel free to submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
