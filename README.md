dummy-file-generator
====================
Simple tool for generating dummy file from 1 byte to N GiBs (Gibibyte).

### Usage
Generating a dummy file of size 1.19 GB (1,283,508,224 bytes) =  
**(1 * 1024 * 1024 * 1024) Gibibyte + (200 * 1024 * 1024) Mebibyte + (50 * 1024) kibibyte + 0 byte**
```batch
dummy-file-generator --g=1 --m=200 --k=50
```


### Reference
[Kilobyte - Wikipedia](https://en.wikipedia.org/wiki/Kilobyte)
