FROM debian:stable-slim

COPY dealer /usr/local/bin/dealer
COPY lotus /usr/local/bin/lotus
COPY ipfs /usr/local/bin/ipfs
COPY Dealer.viper.yaml /etc/Dealer/Dealer.viper.yaml
ENV ENV_DATA_DIR=/data \
    ENV_LOG_DIR=/var/log \
    ENV_LOG_LEVEL=debug \
    ENV_ETH_ENDPOINT="" \
    ENV_INVERSTOR_KEY=""
RUN mkdir $ENV_DATA_DIR

RUN apt update -y
RUN apt install -y libhwloc-dev
RUN ln -s /usr/lib/x86_64-linux-gnu/libhwloc.so /usr/lib/x86_64-linux-gnu/libhwloc.so.5
RUN apt install -y --no-install-recommends ca-certificates curl

WORKDIR /usr/local/bin/

CMD ./dealer run
