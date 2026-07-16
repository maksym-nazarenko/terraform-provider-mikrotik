MikroTik resource code generation
=================================

This tool allows generating MikroTik resources for API client.

## MikroTik client resource

You can generate MikroTik resource and save it to file using remote RouterOS instance (requires valid credentials in environment)
```sh
$ go run ./client/cmd/generate resource -query -basePath /ip/arp -outFile client/arp.go
```

or using offline definition file created with [inspection tool](../inspect/):
```sh
$ go run ./client/cmd/generate resource -basePath /ip/arp -definitionFile arp.json -outFile client/arp.go
```

If you have already generated resource in `arp.go` file, you can also generate a basic test file for it:
```sh
$ go run ./client/cmd/generate test -sourceFile client/arp.go -outFile client/arp_test.go
```

carefully review the test file, as codegen tool is still experimental.
