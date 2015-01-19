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

In the example above - the original docker command without wait-for-weave was - the containers entrypoint is `/bin/backup.sh`:

```bash
$ docker run binocarlos/database-backup --server 10.255.0.1
```

The job of identifying and modifying the container entrypoint is not in the scope of this repository.

## License

MIT
