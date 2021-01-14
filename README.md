# sttp
Simple Transfer Toolkit base on HTTP Protocol

sttp can be used as alternative to *scp*, *rsync*, *lzrz*.

## not goals
* security : this tool pays no attention to data security, DO NOT use it in sensitive environment.

## Features
* upload/download/list/delete file
* simple token verification
* cli tool for client and server
* web browser (auto) support
* ignore config
* absolute/relative path mode 
* multi-user 

## Configuration

config file locate at *.sttp*

## Usage
List files in directory *path/to/directory*
```
curl server:port/path/to/directory
```
Get/Download file at *path/to/file*
```
curl server:port/path/to/file > file
```
Put/Upload file *data.file* to *path/to/save*
```
curl -XPOST --data-binary @data.file server:port/path/to/save
```

## CLI
client
```
sttp get/list path
sttp delete/del path/to/file
sttp down/download  path/to/file
sttp exec mkdir/ls
```

server
```
sttp run -h localhost -p 8008 -root /data 
```

## scp mode
```
sttp server:port/path
```


