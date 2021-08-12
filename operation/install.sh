#!/bin/bash
#
# install.sh 1.0.0
#
# Enables the firewall, installs the newest go and the festivals-fileserver and starts it as a service.
#
# (c)2020-2021 Simon Gaus
#

# Move to working directory
#
cd /usr/local || exit

# Enable and configure the firewall.
# 
if command -v ufw > /dev/null; then

  ufw default deny incoming >/dev/null
  ufw default allow outgoing >/dev/null
  ufw allow OpenSSH >/dev/null
  yes | sudo ufw enable >/dev/null
  echo "Enabled ufw"
  sleep 1

  ufw allow 1910/tcp >/dev/null
  echo "Added festivals-fileserver to ufw using port 1910."
  sleep 1

elif ! [ "$(uname -s)" = "Darwin" ]; then
  echo "No firewall detected and not on macOS. Exiting."
  exit 1
fi

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

# Install festivals-fileserver to /usr/local/bin/festivals-fileserver. TODO: Maybe just link to /usr/local/bin?
#
echo "Downloading current festivals-fileserver..."
yes | sudo git clone https://github.com/Festivals-App/festivals-fileserver.git /usr/local/festivals-fileserver > /dev/null;
cd /usr/local/festivals-fileserver || { echo "Failed to access working directory. Exiting." ; exit 1; }
go build main.go
mv main /usr/local/bin/festivals-fileserver || { echo "Failed to install festivals-fileserver binary. Exiting." ; exit 1; }
mv config_template.toml /etc/festivals-fileserver.conf
mkdir -p /srv/festivals-fileserver/images/resized >/dev/null || { echo "Failed to create the image directories. Exiting." ; exit 1; }
mkdir -p /srv/festivals-fileserver/pdf >/dev/null || { echo "Failed to create the pdf directories. Exiting." ; exit 1; }
echo "Installed festivals-fileserver."
sleep 1

# Install systemd service
#
if command -v service > /dev/null; then

  if ! [ -f "/etc/systemd/system/festivals-fileserver.service" ]; then
    mv operation/service_template.service /etc/systemd/system/festivals-fileserver.service
    echo "Created systemd service."
    sleep 1
  fi

  systemctl enable festivals-fileserver > /dev/null
  systemctl start festivals-fileserver > /dev/null
  echo "Enabled systemd service."
  sleep 1

elif ! [ "$(uname -s)" = "Darwin" ]; then
  echo "Systemd is missing and not on macOS. Exiting."
  exit 1
fi

echo "Done."