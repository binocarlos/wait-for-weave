FROM scratch
MAINTAINER Kai Davenport <kaiyadavenport@gmail.com>
ADD ./stage/wait-for-weave /home/weavetools/wait-for-weave
VOLUME /home/weavetools