# env-web-server

Reply to all web requests with the content of an environment variable.

Intended for parking pages either on public sites or during development work internally.

## Using

Directly:

```console
BODY=hi ./env-web-server
```

Docker/Podman:

```console
docker run -p 2000:2000 -e BODY=hi ghcr.io/outdatedversion/env-web-server/server
```
