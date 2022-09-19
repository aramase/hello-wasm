# hello-wasm

Trying Go Plugin System over WebAssembly using https://github.com/knqyf263/go-plugin.

## Commands

### Generating SDK

```bash
(cd greeting && protoc --go-plugin_out=. --go-plugin_opt=paths=source_relative greeting.proto)
```

### Building

```bash
tinygo build -o plugin/plugin.wasm -scheduler=none -target=wasi --no-debug plugin/plugin.go
```

### Running

```bash
go run main.go
```
