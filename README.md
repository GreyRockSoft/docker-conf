docker-conf
===========

docker-conf is intended to be used with a wrapper script inside of a docker container to create configuration files from environment variables before the main docker application runs.

Usage
=====

Inside of a wrapper script you would use docker-conf this way:

```bash
#!/bin/bash
docker-conf -t /opt/templates

run_your_app_here
```

The `-t` option tells docker-conf where to look for templates that it needs to fill out for your application.

