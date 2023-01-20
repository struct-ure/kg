#!/bin/bash

# fail if any error occurs
set -e

dgraph zero --telemetry "sentry=false" & dgraph alpha --telemetry "sentry=false"
