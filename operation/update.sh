#!/bin/bash
#
# update.sh 1.0.0
#
# Updates the festivals-fileserver and restarts it.
#
# (c)2020-2021 Simon Gaus
#

# Move to working dir
#
mkdir /usr/local/festivals-fileserver/install || { echo "Failed to create working directory. Exiting." ; exit 1; }
cd /usr/local/festivals-fileserver/install || { echo "Failed to access working directory. Exiting." ; exit 1; }

# Get system os
#
if [ "$(uname -s)" = "Darwin" ]; then
  os="darwin"
elif [ "$(uname -s)" = "Linux" ]; then
  os="linux"
else
  echo "System is not Darwin or Linux. Exiting."
  exit 1
fi

# Get systems cpu architecture
#
if [ "$(uname -m)" = "x86_64" ]; then
  arch="amd64"
elif [ "$(uname -m)" = "arm64" ]; then
  arch="arm64"
else
  echo "System is not x86_64 or arm64. Exiting."
  exit 1
fi

# Build url to latest binary for the given system
#
file_url="https://github.com/Festivals-App/festivals-fileserver/releases/latest/download/festivals-fileserver-$os-$arch.tar.gz"
echo "The system is $os on $arch."
sleep 1

# Updating festivals-fileserver to the newest binary release
#
echo "Downloading newest festivals-fileserver binary release..."
curl -L "$file_url" -o festivals-fileserver.tar.gz
tar -xf festivals-fileserver.tar.gz
mv festivals-fileserver /usr/local/bin/festivals-fileserver || { echo "Failed to install festivals-fileserver binary. Exiting." ; exit 1; }
echo "Updated festivals-fileserver binary."
sleep 1

# Removing unused files
#
echo "Cleanup..."
cd /usr/local/festivals-fileserver || { echo "Failed to access home directory. Exiting." ; exit 1; }
rm -r /usr/local/festivals-fileserver/install
sleep 1

# Restart the festivals-fileserver
#
systemctl restart festivals-fileserver
echo "Restarted the festivals-fileserver"
sleep 1

echo "Done!"
sleep 1