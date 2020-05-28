# OriginStamp CLI
Command line client for the [OriginStamp API](https://api.originstamp.com/swagger/swagger-ui.html) written in Go.

## Installation


| Platform   | Installation |
|---|:---|
| MacOS | `brew tap dennis-tra/originstamp-cli && brew install originstamp-cli` |
| Linux | Download the binary from the [latest release](https://github.com/dennis-tra/originstamp-cli/releases). |
| Windows | Download the binary from the [latest release](https://github.com/dennis-tra/originstamp-cli/releases). |

## API-Key

Take your API-Key from the OriginStamp website and set the `ORIGINSTAMP_API_KEY` environment variable in your current shell session:
```
export ORIGINSTAMP_API_KEY=YOUR-API-KEY
``` 
Alternatively put the line above in your `.bashrc` (or equivalent) or provide it on a per-command-basis by providing the `-k` command line flag:
```
stamp -k YOUR-API-KEY filename
```

## Usage

Create a timestamp for a file 
```
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
```
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
```
$ stamp proof filename
https://api.originstamp.com/v3/timestamp/proof/download?token=LONG-TOKEN&name=certificate.Ethereum.file.pdf
```
This will retrieve the `BITCOIN` `SEED` proof. If you want another combination just specify it via command line parameters:
```
$ stamp proof --currency ETHEREUM --proof PDF filename
https://api.originstamp.com/v3/timestamp/proof/download?token=LONG-TOKEN&name=certificate.Ethereum.file.pdf
```
Request your credit usage:
```
$ stamp usage
Consumed credits for the current month:   18.5
Remaining credits for the current month:  31.5
Total number of credits per month:        50.0
---
You have consumed 37.0% of your available credits
```

## Built With

* [urfave/cli](https://github.com/urfave/cli) - The CLI framework

## Versioning

This project uses [SemVer](http://semver.org/) for versioning. For the versions available, see the [releases on this repository](https://github.com/dennis-tra/originstamp-cli/releases). 

## License

This project is licensed under the Apache 2.0 License - see the [LICENSE](LICENSE) file for details
