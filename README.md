# pps-client

A basic client for retrieving and updating Pleasant Password Server credentials

## Usage

To build:

```
make build
```

The binary will be placed in `bin/pps-client`

A valid pleasant password server API token and endpoint are required.  These are configured via environment variables `LOCKBOX` and `LOCKBOX_URL`

```
export LOCKBOX=your_api_token
export LOCKBOX_URL="https://pleasant.password.server"
```

To retrieve a password:

```
pps-client --item e70a3d6a-d596-11e9-9c22-438976518285
```

To retrieve an attachment:

```
pps-client --item e70a3d6a-d596-11e9-9c22-438976518285 --attachment ee960c38-e85b-4275-b421-72029a9dd050
```

like `--item`, `--attachment` returns the raw json from the API.  The key `FileData` holds the file in bas64 format.

If you wish to save the file to disk, extract and convert the data:

```
pps-client --item e70a3d6a-d596-11e9-9c22-438976518285 --attachment ee960c38-e85b-4275-b421-72029a9dd050 | jq -r .FileData | base64 -d > file
```

To update a password:

```
pps-client --item f7f9aa7a-d596-11e9-98e1-5be96b1732e6 --update "new_password"
```

It is assumed that `--update` is json-injection safe.

Attachment updates are not yet supported.
