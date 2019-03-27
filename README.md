# palmtree

[![Go Lang](http://kavehmz.github.io/static/gopher/gopher-front.svg)](https://golang.org/)
[![Build Status](https://travis-ci.org/kavehmz/palmtree.svg?branch=master)](https://travis-ci.org/kavehmz/palmtree)
[![Coverage Status](https://coveralls.io/repos/github/kavehmz/palmtree/badge.svg?branch=master)](https://coveralls.io/github/kavehmz/palmtree?branch=master)


Palm tree is a connection pool package for go which Closes the connections if it wants to discard them.

Underneath it will keep the idle connection in channel buffer.

If channel buffer which is set by `Buffer` value is full, calling `Put` will close the connection and discard it.

Test show how the package can be used.

