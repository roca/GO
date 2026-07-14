# Notes

- [Preview: ranging over functions in Go](https://eli.thegreenplace.net/2023/preview-ranging-over-functions-in-go/)
- - Example-510541

## Trying proposals
You can play with any proposed changes by building Go with the patches implementing the proposal. Using the gotip tool, this is much easier than it sounds.

First of all, install gotip:

```sh
go install golang.org/dl/gotip@latest
```

Now, ask gotip to download and build Go at a special CL number (this is the CL stack implementing the proposal):

```sh
gotip download 510541
```

Once this step is done (it can take a minute or two), you're ready to run the examples; just use gotip instead of go; as in, gotip run, gotip build and so on.
