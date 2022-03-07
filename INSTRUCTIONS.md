```bash
make tfgen
make provider
make build_sdks
cd sdk && go mod tidy && cd -

cp bin/pulumi-resource-checkly $GOPATH/bin
make install_nodejs_sdk


npm install
yarn link @pulumi/checkly
```

### Questions

1. ID type conversion/parsing => string/int
2. The best way to use file content as input (Assets vs Node.js read file)
3. Default values (string, booleans, null)