# Intro
This project contains a simple command line app called bpf_validate that checks the validty of a BPF-style PCAP filter expression.  

In addition to the Go source there are also two pre-built binaries:

* bpf_validate_full.bin - No build mods
* bpf_validate - Built with `-ldflags="-s -w"` to remove debugging info and related symbol tables. Reduces the binary size by about 1MB. 

