# Task : Hello Protobuf

- 1. Install required tools : Go & Protocol Buffers compiler

- 2. For IDE, use any tools you want (e.g. Visual Studio Code)

- 3. Create a protobuf schema definition containing message Hello, with only one field : name (string).

- 4. Generate Go source code from protobuf definition

```sh
protoc \
--go_out=. --go_opt=paths=source_relative \
./hello.proto
```

- 5. Write Go application to create new Hello and display it
