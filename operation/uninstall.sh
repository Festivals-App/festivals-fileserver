#!/bin/bash
#
# uninstall.sh 1.0.0
#
# Removes the firewall configuration, uninstalls go, git and the festivals-fileserver and stops and removes it as a service.
#
# (c)2020 Simon Gaus
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
# Supported firewalls: ufw and firewalld
# This step is skipped under macOS.
#
if command -v firewalld > /dev/null; then

  firewall-cmd --permanent --remove-service=festivals-fileserver >/dev/null
  rm -f /etc/firewalld/services/festivals-fileserver.xml >/dev/null
  rm -f /etc/firewalld/services/festivals-fileserver.xml.old >/dev/null
  firewall-cmd --reload >/dev/null
  echo "Removed firewalld configuration"
  sleep 1

elif command -v ufw > /dev/null; then

  ufw delete allow 1919/tcp >/dev/null
  echo "Removed ufw configuration"
  sleep 1

elif ! [ "$(uname -s)" = "Darwin" ]; then
  echo "No firewall detected and not on macOS. Exiting."
  exit 1
fi

# Remove go
#
rm -R /usr/local/go
rm /usr/local/bin/go
rm /usr/local/bin/gofmt
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
