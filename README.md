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

Item uuids are accepted through stdin

To retrieve a password:

```
pps-client <<< e70a3d6a-d596-11e9-9c22-438976518285
```

To update a password:

```
pps-client -update -pass=foo <<< f7f9aa7a-d596-11e9-98e1-5be96b1732e6
```
