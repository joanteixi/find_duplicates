# Find duplicates in your FileSystem
This is a copy of a code finded in https://www.socketloop.com/tutorials/golang-find-duplicate-files-with-filepath-walk just to learn GoLang. 

It just search inside a folder all the files (and files inside subfolders), store a sha512 of the file checksum in memory and compares files by this hash. 

## Take care that it deletes the duplicate folder that find. 

