FROM ubuntu
MAINTAINER Kai Davenport <kaiyadavenport@gmail.com>

RUN mkdir -p /home/weavetools
ADD ./stage/wait-for-weave /home/weavetools/wait-for-weave
VOLUME /home/weavetools