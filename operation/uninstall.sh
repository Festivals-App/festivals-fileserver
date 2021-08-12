#!/bin/bash
#
# uninstall.sh 1.0.0
#
# Removes the firewall configuration, uninstalls go, git and the festivals-fileserver and stops and removes it as a service.
#
# (c)2020-2021 Simon Gaus
#

# Move to working directory
#
cd /usr/local || exit

# Stop the service
#
systemctl stop festivals-fileserver >/dev/null
echo "Stopped festivals-fileserver"
sleep 1

# Remove systemd configuration
#
systemctl disable festivals-fileserver >/dev/null
rm /etc/systemd/system/festivals-fileserver.service
echo "Removed systemd service"
sleep 1

# Remove the firewall configuration.
# This step is skipped under macOS.
#
if command -v ufw > /dev/null; then

  ufw delete allow 1910/tcp >/dev/null
  echo "Removed ufw configuration"
  sleep 1

elif ! [ "$(uname -s)" = "Darwin" ]; then
  echo "No firewall detected and not on macOS. Exiting."
  exit 1
fi

# Remove go
#
apt-get --purge remove golang -y
apt autoremove -y
echo "Removed go"
sleep 1

# Remove festivals-server
#
rm /usr/local/bin/festivals-fileserver
rm /etc/festivals-fileserver.conf
rm -R /var/log/festivals-fileserver
rm -R /srv/festivals-fileserver
echo "Removed festivals-server"
sleep 1

echo "Done"
