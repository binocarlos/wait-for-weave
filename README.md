# wait-for-weave

A small binary that waits for a weave network to be connected before running the cli arguments as a command.

Useful when overriding a docker container entrypoint such that the container will wait for the weave network to be ready before running.

## usage

wait-for-weave is not designed to be run as a stand alone container.

It is intended to have it's volumes mounted onto another container using `--volumes-from`.

First - run the wait-for-weave container ensuring you provide a `--name` parameter and a `WAIT_FOR_WEAVE_QUIT=yes` environment variable.

```bash
$ docker run --name weavetools \
  -e "WAIT_FOR_WEAVE_QUIT=yes" \
  binocarlos/wait-for-weave /home/weavetools/wait-for-weave
```

The `WAIT_FOR_WEAVE_QUIT` variable enables us to run the binary but without waiting - otherwise docker complains that we have no given it a command to run.

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

## example usage

The command I ran without weave (where wait-for-weave lived in /srv/projects):

```bash
$ docker run \
  -v /srv/projects/wait-for-weave/stage/wait-for-weave:/bin/wait-for-weave \
  --entrypoint="/bin/wait-for-weave" \
  ubuntu bash -c "while true; do echo hello && sleep 1; done"
```

This printed `interface ethwe not found after 10 seconds`

The command I ran with weave:

```bash
$ sudo weave run 10.255.0.14/8 \
  -v /srv/projects/wait-for-weave/stage/wait-for-weave:/bin/wait-for-weave \
  --entrypoint="/bin/wait-for-weave" \
  ubuntu bash -c "while true; do echo hello && sleep 1; done"
```

This printed the container id and when I `docker logs $ID` it was printing hello lots of times (i.e. it found ethwe and ran the entrypoint)

## licence

Copyright 2015 Kai Davenport & Contributors

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.  You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.  See the License for the specific language governing permissions and limitations under the License.