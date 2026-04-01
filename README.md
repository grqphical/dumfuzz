# dumfuzz
A CLI tool to quickly generate dummy data for testing applications or creating example data.

## Usage
```bash
dumfuzz -n <count> -param name=type -output-format <format>
```

`-n` is the number of data entries you want generated

`-param name=type` defines a field with the given name and generated data type. This flag can be used multiple times

`-output-format` defines how the data should be output, currently the only output supported is CSV

## Currently Supported Fuzzers (Data types)
- `randint` A random integer between 0 and 128k
- `randfloat` A random real number between 0 and 1
- `randbigint` A large random integer

## License
dumfuzz is licensed under the Apache 2.0 license
