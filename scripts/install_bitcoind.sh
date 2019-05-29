#!/usr/bin/env bash

set -ev

export PEERCOIND_VERSION=0.8.0
# delete when it becomes full release
export RC=rc9

if sudo cp ~/peercoin/peercoin-$PEERCOIND_VERSION/bin/peercoind /usr/local/bin/peercoind
then
        echo "found cached peecoind"
else
        mkdir -p ~/peercoin && \
        pushd ~/peercoin && \
        wget https://github.com/peercoin/peercoin/releases/download/v${PEERCOIND_VERSION}ppc.${RC}/peercoin-${PEERCOIND_VERSION}-x86_64-linux-gnu.tar.gz && \
        tar xvfz peercoin-$PEERCOIND_VERSION-x86_64-linux-gnu.tar.gz && \
        sudo cp ./peercoin-$PEERCOIND_VERSION/bin/peercoind /usr/local/bin/peercoind && \
        popd
fi

