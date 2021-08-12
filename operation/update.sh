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
cd /usr/local || exit

# Stop the festivals-server
#
systemctl stop festivals-fileserver
echo "Stopped festivals-fileserver"
sleep 1

# Install go if needed.
# Binaries linked to /usr/local/bin
#
if ! command -v go > /dev/null; then
  echo "Installing go..."
  apt-get install golang -y > /dev/null;
fi

# Install git if needed.
#
if ! command -v git > /dev/null; then
  echo "Installing git..."
  apt-get install git -y > /dev/null;
fi

# Updating festivals-fileserver to the newest version
#
echo "Downloading current festivals-fileserver..."
yes | sudo git clone https://github.com/Festivals-App/festivals-fileserver.git /usr/local/festivals-fileserver > /dev/null;
cd /usr/local/festivals-fileserver || { echo "Failed to access working directory. Exiting." ; exit 1; }
go build main.go
mv main /usr/local/bin/festivals-fileserver || { echo "Failed to install festivals-fileserver binary. Exiting." ; exit 1; }
echo "Installed festivals-fileserver."
sleep 1

# Updating go to the newest version
#
systemctl start festivals-fileserver
echo "Started festivals-fileserver"
sleep 1

# Removing unused files
#
echo "Cleanup..."
cd /usr/local || exit
rm -R /usr/local/festivals-fileserver
sleep 1

echo "Done!"
sleep 1
