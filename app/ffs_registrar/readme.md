# FunctionFS Registrar

If you’re cross-compiling from x86_64 → arm64 u need these:

```bash
export GOOS=linux
export GOARCH=arm64
export CGO_ENABLED=0

go build -trimpath -ldflags="-s -w -buildid=" -o ./tools/builder/assets/ffs_registrar ./app/ffs_registrar
```