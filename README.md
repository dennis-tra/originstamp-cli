# OriginStamp CLI
A CLI for the [OriginStamp API](https://api.originstamp.com/swagger/swagger-ui.html) written in Go.

## Installation

### Mac OS
```shell script
# Installation
brew tap dennis-tra/originstamp-cli
brew install originstamp-cli

# timestamp your file
stamp -k YOUR-API-KEY filename
```

## API-Key

Take your API-Key from the OriginStamp website and set the `ORIGINSTAMP_API_KEY` environment variable in your current shell session:
```shell script
export ORIGINSTAMP_API_KEY=YOUR-API-KEY
``` 
Alternatively put this line in your `.bashrc` (or equivalent) or provide it on a per command basis by providing the `-k` command line flag:
```shell script
stamp -k YOUR-API-KEY filename
```

## Usage

Create a timestamp for a file 
```shell script
$ stamp filename
Successfully initiated timestamp creation!
  CURRENCY |       STATUS |                 TIMESTAMP | TRANSACTION
       --- |          --- |                       --- | ---
   BITCOIN |     RECEIVED |                           |
  ETHEREUM |     RECEIVED |                           |
      AION |     RECEIVED |                           |
SUEDKURIER |     RECEIVED |                           | 
```

Retrieve the status of a timestamp:
```shell script
$ stamp status filename
  CURRENCY |       STATUS |                 TIMESTAMP | TRANSACTION
       --- |          --- |                       --- | ---
   BITCOIN |     RECEIVED |                           |
  ETHEREUM | TAMPER_PROOF | 2020-05-11T19:01:16+02:00 | 0xee6df9c7e4a196e7223bde4d15b565dfb215e9833ac3cf6a5eba9b1039424561
      AION | TAMPER_PROOF | 2020-05-11T18:50:15+02:00 | 890bef8ac5d63fb741ee74cba095b69271d782a1206855b45a0b962e3d013d9b
SUEDKURIER |     RECEIVED |                           |
```
Possible status values are:

| Status | Description|
|:---|:---|
|RECEIVED|The hash was not broadcasted yet, but received from our backend|
|BROADCASTED|The hash was included into a transaction and broadcasted to the network, but not included into a block|
|INCLUDED|The transaction was included into the latest block|
|TAMPER_PROOF|The timestamp for your hash was successfully created.|

As soon as the file is in the status of `TAMPER_PROOF` you can request the proof `PDF` or `SEED`
```shell script
$ stamp proof version
https://api.originstamp.com/v3/timestamp/proof/download?token=LONG-TOKEN&name=certificate.Ethereum.file.pdf
```
This will retrieve the `BITCOIN` `SEED` proof. If you want another combination just specify it via command line parameters:
```shell script
$ stamp proof --currency ETHEREUM --proof PDF version
https://api.originstamp.com/v3/timestamp/proof/download?token=LONG-TOKEN&name=certificate.Ethereum.file.pdf
```


### Help Output
```shell script
NAME:
   stamp - create anonymous, tamper-proof timestamps for any digital content

USAGE:
   stamp [global options] command [command options] [arguments...]

VERSION:
   0.1.0

AUTHOR:
   Dennis Trautwein <dennis.trautwein@originstamp.com>

COMMANDS:
   usage    Retrieve information about the current api usage.
   status   Retrieve timestamp information for a certain file
   proof    Retrieve the timestamp proof for a certain file or hash
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --api-key value, -k value  The OriginStamp AG API-Key (also applies to all sub commands) [$ORIGINSTAMP_API_KEY]
   --comment value, -c value  Comment (max. 256 chars) for the timestamp for indexing and searching (public)
   --format value             Go layout of how to format timestamp when printed to the screen. See https://golang.org/pkg/time/#pkg-constants. Defaults to RFC3339. (default: "2006-01-02T15:04:05Z07:00")
   --hash value               Provide the hash string instead of a file
   --help, -h                 show help (default: false)
   --version, -v              print the version (default: false)

COPYRIGHT:
   OriginStamp AG 2020
```

## Built With

* [urfave/cli](https://github.com/urfave/cli) - The CLI framework

## Versioning

This project uses [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/your/project/tags). 

## License

This project is licensed under the Apache 2.0 License - see the [LICENSE](LICENSE) file for details
