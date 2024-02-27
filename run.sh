#!/bin/bash

# It's a sample on how to build the project in order to be compatible with plugins mode.
# Here "-gcflags" options allows using the debugger while using plugins.
# Without this options an error is thrown saying that the build in incompatible.
#
go run .