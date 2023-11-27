#!/bin/bash
#
# install.sh 1.0.0
#
# Enables the firewall, installs the newest festivals-fileserver and starts it as a service.
#
# (c)2020-2022 Simon Gaus
#

# Test for web server user
#
WEB_USER="www-data"
id -u "$WEB_USER" &>/dev/null;
if [ $? -ne 0 ]; then
  WEB_USER="www"
  if [ $? -ne 0 ]; then
    echo "Failed to find user to run web server. Exiting."
    exit 1
  fi
fi

# Move to working dir
#
mkdir -p /usr/local/festivals-fileserver/install || { echo "Failed to create working directory. Exiting." ; exit 1; }
cd /usr/local/festivals-fileserver/install || { echo "Failed to access working directory. Exiting." ; exit 1; }

echo "Installing festivals-fileserver using port 1910."
sleep 1

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

# Install festivals-fileserver to /usr/local/bin/festivals-fileserver. TODO: Maybe just link to /usr/local/bin?
#
echo "Downloading newest festivals-fileserver binary release..."
curl -L "$file_url" -o festivals-fileserver.tar.gz
tar -xf festivals-fileserver.tar.gz
mv festivals-fileserver /usr/local/bin/festivals-fileserver || { echo "Failed to install festivals-fileserver binary. Exiting." ; exit 1; }
echo "Installed the festivals-fileserver binary to '/usr/local/bin/festivals-fileserver'."
sleep 1

## Install server config file
mv config_template.toml /etc/festivals-fileserver.conf
echo "Moved default festivals-server config to '/etc/festivals-fileserver.conf'."
sleep 1

## Prepare log directory
mkdir /var/log/festivals-fileserver || { echo "Failed to create log directory. Exiting." ; exit 1; }
echo "Create log directory at '/var/log/festivals-fileserver'."

## Prepare file directories
mkdir -p /srv/festivals-fileserver/images/resized >/dev/null || { echo "Failed to create the image directories. Exiting." ; exit 1; }
mkdir -p /srv/festivals-fileserver/pdf >/dev/null || { echo "Failed to create the pdf directories. Exiting." ; exit 1; }
echo "Created folders to hold uploaded files at '/srv/festivals-fileserver'."
sleep 1

## Prepare server update workflow
mv update.sh /usr/local/festivals-fileserver/update.sh
chmod +x /usr/local/festivals-fileserver/update.sh
cp /etc/sudoers /tmp/sudoers.bak
echo "$WEB_USER ALL = (ALL) NOPASSWD: /usr/local/festivals-fileserver/update.sh" >> /tmp/sudoers.bak
# Check syntax of the backup file to make sure it is correct.
visudo -cf /tmp/sudoers.bak
if [ $? -eq 0 ]; then
  # Replace the sudoers file with the new only if syntax is correct.
  sudo cp /tmp/sudoers.bak /etc/sudoers
else
  echo "Could not modify /etc/sudoers file. Please do this manually." ; exit 1;
fi

# Enable and configure the firewall.
# 
if command -v ufw > /dev/null; then

  ufw allow 1910/tcp >/dev/null
  echo "Added festivals-fileserver to ufw using port 1910."
  sleep 1

elif ! [ "$(uname -s)" = "Darwin" ]; then
  echo "No firewall detected and not on macOS. Exiting."
  exit 1
fi

# Install systemd service
#
if command -v service > /dev/null; then

  if ! [ -f "/etc/systemd/system/festivals-fileserver.service" ]; then
    mv service_template.service /etc/systemd/system/festivals-fileserver.service
    echo "Created systemd service."
    sleep 1
  fi

  systemctl enable festivals-fileserver > /dev/null
  echo "Enabled systemd service."
  sleep 1

elif ! [ "$(uname -s)" = "Darwin" ]; then
  echo "Systemd is missing and not on macOS. Exiting."
  exit 1
fi

## Set appropriate permissions
#
chown -R "$WEB_USER":"$WEB_USER" /usr/local/festivals-fileserver
chown -R "$WEB_USER":"$WEB_USER" /var/log/festivals-fileserver
chown -R "$WEB_USER":"$WEB_USER" /srv/festivals-fileserver
chown "$WEB_USER":"$WEB_USER" /etc/festivals-identity-server.conf
echo "Seting appropriate permissions..."
sleep 1

# Removing unused files
#
echo "Cleanup..."
cd /usr/local/festivals-fileserver || exit
rm -R /usr/local/festivals-fileserver/install
sleep 1

echo "Done!"
sleep 1

echo "You can start the server manually by running 'systemctl start festivals-fileserver' after you updated the configuration file at '/etc/festivals-fileserver.conf'"
sleep 1