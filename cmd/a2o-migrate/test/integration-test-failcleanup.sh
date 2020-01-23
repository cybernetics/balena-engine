#!/bin/sh

cat /etc/os-release
balena-engine info || exit 1

# ls -lR /var/lib/balena-engine/

./a2o-migrate -version
./a2o-migrate -debug -migrate

# we mess up the files here intentionally to test the
# rollback path
rm -rf /var/lib/balena-engine/image/overlay/l

./a2o-migrate -debug -rollback

# ls -lR /var/lib/balena-engine/
