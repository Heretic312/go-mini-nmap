# Go Mini Nmap

A **mini Nmap scanner in Go** with:

- TCP & UDP scanning
- Banner grabbing for TCP
- Multi-host scanning
- Concurrency-safe worker pool
- CSV & JSON output

## Installation

```bash
git clone https://github.com/Heretic312/go-mini-nmap.git
cd go-mini-nmap
go build
````

## Usage

```bash
./go-mini-nmap -hosts 192.168.1.10,192.168.1.11 -start 1 -end 1024 -tcp=true -udp=false -csv=results.csv -json=results.json
```

## Features

* Color-coded console output
* Configurable concurrency
* Exportable results
* Portfolio-ready structure

## Examples

Check `examples/scan_results.csv` and `examples/scan_results.json` for sample outputs.

## License

MIT

```

---