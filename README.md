# SFTP Guard 🛡
SFTP Guard is a custom CLI script that checks files in SFTP servers for their updated time. This basic check is important to ensure dependable tools run with the most up to date data.

## Features
- 🔎 Regex match to include/exclude paths
- ⚡ Dynamic regex strings (use {yyyy}, {mm}, {dd} to create dynamic regex strings before regex compilation). {yyyy} gets replaced with the current year, {mm} with current month (with leading 0) and {dd} current day (with leading 0).
```py
# Using the following regex pattern:
"(?si)Dynamic_Filename_{yyyy}{mm}{dd}.*"

#will result in the following regex compilation (on the day 24 Jan 2025)
"(?si)Dynamic_Filename_20240124.*" 
```


## Usage
You will need to create a `clients.json` file with the following properties. This file must be in the same directory as the executable. You can [use this template](https://github.com/joaooliveirapro/sftpguard/blob/main/src/clients.template.json).
```py
[
    "client_name": "Client",
        # If a file has been modified over this treshold, 
        # then it will be highlighted on the table.
        "treshold_hours": 2,  
        "feeds": [
            {
                "feed_name": "some name",
                "host": "files.mysftpserver.com",
                # If the filename is known and static, use this for faster results.
                "filepaths": [
                    "/path/to/my/file.xml"
                ], 
                # If all files in a directory must be considered, add the directory path here. 
                # It's not recursive.
                "directories": [
                    "/list/all/here/",
                    "/list/all/here/too"
                ],
                # For more complex file lookup use a regex match pattern. 
                # Can also use {yyyy}, {mm}, {dd} for dynamic regex string compilation.
                "regex": [
                    {
                        # Directory to apply the regex lookup
                        "directory": "/path/to/dir/",
                        # Regex string to match with each filename
                        "patterns": [
                            "(?si)Dynamic_Filename_{yyyy}{mm}{dd}.*"
                        ]
                    }
                ],
                # 22 for SFTP or 21 for FTP
                "port": 22, 
                "username": "<sftp_username>",
                "password": "<sftp_password>"
            }
        ]
]

```

## Output
A table is shown with 
![example output](https://github.com/joaooliveirapro/sftpguard/blob/main/assets/example1.png)

This is also exported as `data.txt`.


## Installation
Build from source and run it. 
```sh
# Windows
git clone https://github.com/joaooliveirapro/sftpguard.git
cd sftpguard
make build       # (go build -o sftpguard.exe ./src)
.\sftpguard.exe  # launch script (assumes clients.json is configured)
```


### License
The MIT License (MIT)