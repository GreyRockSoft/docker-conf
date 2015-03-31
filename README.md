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

Template Format
===============

docker-conf templates use the following format for creating templates.  Template files must end with `.tmpl`  The `.tmpl` will get removed when processing occurs so if you have a file named `nginx.conf.tmpl` the resulting file will be named `nginx.conf`

You can inject values into the template files with `{{ env name }}` where `env` is telling docker-conf to look for the enviornment variable `name`
