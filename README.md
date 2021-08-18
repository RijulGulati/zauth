# zauth

zauth is a 2FA (Two-Factor Authentication) application for terminal written in Go.


---


![zauth](/assets/zauth.gif)


## Features
- Supports both TOTP and HOTP codes.
- Add new entries directly from CLI.
    - support setting custom digits (default: 6)
    - support setting a custom period (TOTP) (default: 30)
    - support SHA1, SHA256 and SHA512 algorithms (TOTP)
- Import/Export [andOTP](https://github.com/andOTP/andOTP) backups (encrypted files supported).
- More upcoming features in [What's next](https://github.com/grijul/zauth#whats-next)

*If you would like any other app to be supported, please [create an issue](https://github.com/grijul/zauth/issues) and (if possible) provide an unencrypted sample backup file. Of course I am accepting pull requests as well :)*



---



## Installation
    $ go install github.com/grijul/zauth@latest

By default, zauth stores it entries in `$HOME/.zauth` directory.

### Using Docker

zauth can be installed using docker as well. Running the following command pulls zauth image and runs `zauth -h` command.


    $ docker run grijul/zauth:latest zauth -h


You can bind container's `/root/.zauth` directory to your host's `$HOME/.zauth` directory to use zauth.json from your host system. Something like this should work:


    $ docker run -v $HOME/.zauth:/root/.zauth zauth:latest zauth



**Important Note:** There is only 1 docker image with `latest` tag on docker hub. Since there is no release cycle (as of now), I manually have to update the docker image whenever there are new commits. So the image is subject to be outdated and may not contain latest changes/fixes. I will try to update the image as frequently as possible.

If latest changes are desired, you can [build docker image from source]() (it's easier than it sounds).



---



## Building from source
* Clone repository and cd into dir

        $ git clone https://github.com/grijul/zauth.git && cd zauth

* Build using `go build` command
        
        $ go build .

### Building docker image from source
1. Clone repository and cd into dir

        $ git clone https://github.com/grijul/zauth.git && cd zauth


2. Build docker image
    
        $ docker build -t zauth:latest .

3. Run docker image
        
        $ docker run zauth:latest zauth -h



---



## Examples

**Print OTP**
    
    $ zauth


If zauth.json file exists, corresponding entries will be printed. Else the above command will give a `file not found` error.

This will simply print zauth entries with OTP and exit.
If you wish to watch zauth entries update every second, you can use `watch` command.

    $ watch -n1 zauth


---


**Add new entry**

    $ zauth entry -new

A prompt will be displayed to capture necessary details (secret, issuer, etc..).


---


**List entries**

    $ zauth entry -list


---


**Import decrypted file**
    
    $ zauth import -file <import_file> -type <import_type>

`-file` flag tells zauth which file to import

`-type` flag tells zauth what type of file is being imported ([supported files](https://github.com/grijul/zauth#supported-app-files-for-import))


---


**Import encrypted file**

    $ zauth import -file <import_file> -type <import_type> -decrypt



`-decrypt` flag tells zauth that import file is encrypted, and prompts user for decryption password. If not provided, files are assumed to be decrypted.


---


**Import file (entries are overwritten)**

    $ zauth import -file <import_file> -type <import_type> -overwrite

`-overwrite` flag overwrites existing entries with new entries. If not provided, entries are appended.


---


**Export file**

    $ zauth export -type <export_type> -encrypt

`-encrypt` flag tells zauth that exported file should be encrypted. If not provided, exported file is decrypted.

The file exported (encrypted/decrypted) is compatible with `export_type` app. This means user should be able to import this exported file back to `export_type` app.


---


### Supported app files for import
- [andOTP](https://github.com/andOTP/andOTP) - supports both encrypted/decrypted file. [`-type=andotp`]

### Supported app files for export
- [andOTP](https://github.com/andOTP/andOTP) - supports both encrypted/decrypted file. [`-type=andotp`]


### What's next
- zauth uses json file to store it's entries. At this moment, this json file is unencrypted. It'd be better we could have encrypted file instead.
- Edit/Delete entries from CLI.

## Contact
Feel free to get in touch with me via [Twitter](https://twitter.com/grijul) or [Email](mailto:grijul@protonmail.ch).


## License
[MIT](https://github.com/grijul/zauth/blob/main/LICENSE)

