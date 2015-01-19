# wait-for-weave

A small binary that waits for a weave network to be connected before running the cli arguments as a command.

Useful when overriding a docker container entrypoint such that the container will wait for the weave network to be ready before running.

## usage

wait-for-weave is not designed to be run as a stand alone container.

It is intended to have it's volumes mounted onto another container using `--volunes-from`.

First - run the wait-for-weave container ensuring you provide a `--name` parameter:

```bash
$ docker run --name weavetools binocarlos/wait-for-weave
```

Then - you can run another container with a `--volumes-from=weavetools` option.

This enables you change the `--entrypoint` to `/home/weavetools/wait-for-weave`.

Pass the original entrypoint as arguments and it will run only once the weave network is up and running:

```bash
$ docker run \
    --volumes-from=weavetools \
    --entrypoint="/home/weavetools/wait-for-weave" \
    binocarlos/database-backup /bin/backup.sh --server 10.255.0.1
```

In the example above - the containers entrypoint is `/bin/backup.sh` and the original docker command without wait-for-weave was:

```bash
$ docker run binocarlos/database-backup --server 10.255.0.1
```

The job of identifying and modifying the container entrypoint and adding the `--volumes-from` flag is not in the scope of this project.  This is much better handled by a docker proxy or similar tool.

## exit code

If the `ethwe` interface is not found within 10 seconds then `wait-for-weave` will NOT execute the entrypoint, will print a message to `stderr` and exit with a non-zero exit code.

## License

MIT
