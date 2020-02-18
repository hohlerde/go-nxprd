# Important

**This repository has moved to https://codeberg.org/hohlerde/go-nxprd**

# go-nxprd

go-nxprd is a very *simple* wrapper written in GO for NXP's NFC Reader Library
(NxpRdLib) in order to interface with the [NXP EXPLORE-NFC](http://www.nxp.com/products/identification-and-security/nfc-and-reader-ics/nfc-frontend-solutions/explore-nfc-exclusive-from-element14:PNEV512R)
expansion board for the [Raspberry Pi](http://www.raspberrypi.org).

Right now only Mifare Ultralight cards are supported for reading and writing
bytes. This may change in the future, but currently I only have the Ultralight
card that was shipped with the NFC board.  

This wrapper was heavily inspired by svvitale's [nxppy](https://github.com/svvitale/nxppy) python wrapper. I use a slightly modified version of his nice bash script ([get_nxpRdLib.sh](https://github.com/hohlerde/go-nxprd/blob/master/get_nxpRdLib.sh)).

## License
All files in this repository are distributed under the MIT license.

### External components
This wrapper is heavily based on NXP's example code (mainly MIFAREClassic and
BasicDiscoveryLoop). The example code was reorganized and adjusted to be used
from a GO wrapper. NXP still retains full copyright and ownership of the example
code and the NFC Reader Library. The license conditions that need to be accepted
in order to use this project in conjunction with the NFC Reader Library can be
found in the document
[NXP_Library_License.pdf](https://github.com/hohlerde/go-nxprd/blob/master/NXP_Library_License.pdf).

## Compatibility
Tested with [GO 1.6.3](https://storage.googleapis.com/golang/go1.6.3.linux-armv6l.tar.gz)
on Raspian Jessie.

## Requirements
The EXPLORE-NFC board relies on SPI being enabled. Please enable SPI using
raspi-config prior to using go-nxprd. You may follow this
[tutorial](http://www.raspberrypi-spy.co.uk/2014/08/enabling-the-spi-interface-on-the-raspberry-pi/)
from raspberrypi-spy.

As go-nxprd is a GO wrapper for NXP's C based NFC Reader Library, you need to
install and configure GO on your Raspberry Pi. Instructions can be found at
[https://golang.org/doc/install](https://golang.org/doc/install).

## Installation
go-nxprd uses cgo to interface the NFC Reader Library, so before using the wrapper
please make sure you have a working gcc and cmake (used for building the NFC
Reader Library) installed.

The easiest way is to use the script `get_nxpRdLib.sh` in the repository.
It will download the Reader Library sources from DropBox, extract them to the
folder `nxp` and run `cmake` to build the library for the EXPLORE-NFC board.
Beforehand it will check whether the debian packages build-essential and cmake
are installed and if not, will install them. Therefore, the script should be run
with root privileges.

cgo relies on the original folder structure of the Reader Library in the `nxp`
folder, so better don't change it. You'll find the compiled NFC Reader Library
in the `nxp/build/linux` folder.

If the NFC Reader Library was build successfully, you can try to build the
GO based test application for Mifare Ultralight cards. It is located under
`cmd/gomful` in the repository. As the NFC Reader Library will be statically
linked to the wrapper, you need to pass some extra parameters to GO's build
command:

```
go build --ldflags '-extldflags -static' .
```

This will build the test application in the current folder. You'll find the
build command in the script [build.sh](https://github.com/hohlerde/blob/master/cmd/gomful/build.sh) in the `gomful` folder.
After a successfull build you can start the application simply by using:

```
./gomful
```

The output may be similar to the following:

```
Mifare Ultralight Test

Place a card near the reader and press <Enter>


Card            : MifareUL
Tag type        : 2
Technology type : A
UID             : 0x04 0x90 0x98 0xC2 0x37 0x30 0x81
ATQ(A)          : 0x44 0x00
SAK             : 0x00

Trying to read memory pages 0-15:

Page 00: 0x04 0x90 0x98 0x84
Page 01: 0xC2 0x37 0x30 0x81
Page 02: 0x44 0x48 0x00 0x00
Page 03: 0x00 0x00 0x00 0x00
Page 04: 0x00 0x00 0x00 0x00
Page 05: 0x00 0x00 0x00 0x00
Page 06: 0x00 0x00 0x00 0x00
Page 07: 0x00 0x00 0x00 0x00
Page 08: 0x00 0x00 0x00 0x00
Page 09: 0x00 0x00 0x00 0x00
Page 10: 0x00 0x00 0x00 0x00
Page 11: 0x00 0x00 0x00 0x00
Page 12: 0x00 0x00 0x00 0x00
Page 13: 0x00 0x00 0x00 0x00
Page 14: 0x00 0x00 0x00 0x00
Page 15: 0x00 0x00 0x00 0x00
```

## Usage

To use go-nxprd in your GO application simply use the following import
statement:

```go
import "github.com/hohlerde/go-nxprd"
```

The main package of the wrapper is `nxprd`. In order to detect and read a
card/tag the wrapper and also the NFC Reader Library need to be initialized.
This is done by calling `nxprd.Init()`:

```go
err := nxprd.Init()
if err != nil {
    // handle error
}
```

Now you are ready to select a card/tag by calling `Discover()`. Please note that
the Discover function blocks the caller. It has an optional timeout parameter,
which can be used to specify how long the function should try to detect a
card/tag before returning. The default timeout is 1000ms.

```go
dev, err := nxprd.Discover(5000) // timeout = 5s
if err != nil {
    // handle error
}
```

If a card/tag could be detected, it can be accessed through the `Device`
structure returned by `Discover()`.

Parameters of the card/tag can be simply accessed by the `Params` member.

```go
// Prints the UID of the card/tag
fmt.Printf("%+v", dev.Params.UID)
```

And if we have a Mifare Ultralight card, we can even read and write data.

```go
// Right now only Mifare Ultralight cards are supported
if dev.Params.DevType == MifareUL {
    // Read all bytes (4) of page/block 10
    buffer, err := dev.Reader.ReadBlock(10)
    if err != nil {
        // handle error
    }

    buffer[0] = 66 // Change the first byte of page/block 10

    // and write it back to page 10.
    err = dev.Writer.WriteBlock(10, buffer)
    if err != nil {
        // handle error
    }
}
```

At the end we need to cleanup the wrapper and the NFC Reader Library by calling `DeInit()`.

```go
nxprd.DeInit()
```

As an example you also may want to take a look at the [test application](https://github.com/hohlerde/go-nxprd/blob/master/cmd/gomful/main.go).

## Feedback
Feedback and PRs are always welcome!
