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

## Installation
    $ go install github.com/grijul/zauth@latest

By default, zauth stores it entries in `$HOME/.zauth` directory.

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
[MIT]()

