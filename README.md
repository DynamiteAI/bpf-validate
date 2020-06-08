# Intro
This project contains a simple command line app called bpf_validate that checks the validty of a BPF-style PCAP filter expression.  

Go source there are also two pre-built binaries:

* bpf_validate_full.bin
* bpf_validate - Built with `-ldflags="-s -w"` to remove debugging info and related symbol tables. 

