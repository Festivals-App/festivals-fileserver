#!/bin/bash
#
# install.sh 1.0.0
#
# Enables the firewall, installs the newest go and the festivals-fileserver and starts it as a service.
#
# (c)2020 Simon Gaus
#

# Move to working directory
#
cd /usr/local || exit

# Enable and configure the firewall.
# Supported firewalls: ufw and firewalld
# This step is skipped under macOS.
#
if command -v firewalld > /dev/null; then

  systemctl enable firewalld >/dev/null
  systemctl start firewalld >/dev/null
  echo "Enabled firewalld"
  sleep 1

  firewall-cmd --permanent --new-service=festivals-fileserver >/dev/null
  firewall-cmd --permanent --service=festivals-fileserver --set-description="A live and lightweight go server app providing static files for the FestivalsAPI." >/dev/null
  firewall-cmd --permanent --service=festivals-fileserver --add-port=1910/tcp >/dev/null
  firewall-cmd --permanent --add-service=festivals-fileserver >/dev/null
  firewall-cmd --reload >/dev/null
  echo "Added festivals-fileserver.service to firewalld"
  sleep 1

elif command -v ufw > /dev/null; then

  ufw default deny incoming >/dev/null
  ufw default allow outgoing >/dev/null
  ufw allow OpenSSH >/dev/null
  yes | sudo ufw enable >/dev/null
  echo "Enabled ufw"
  sleep 1

  ufw allow 1919/tcp >/dev/null
  echo "Added festivals-fileserver to ufw"
  sleep 1

elif ! [ "$(uname -s)" = "Darwin" ]; then
  echo "No firewall detected and not on macOS. Exiting."
  exit 1
fi

# Install go to /usr/local/go if needed.
# Binaries linked to /usr/local/bin
#
if ! command -v go > /dev/null; then
  echo "Downloading current go version..."
  goVersion="$(curl --silent "https://golang.org/VERSION?m=text")"
  currentGo="$goVersion.linux-amd64.tar.gz"
  goURL="https://dl.google.com/go/$currentGo"
  goOut=/var/cache/festivals-fileserver/$currentGo

  if ! [ -f "$goOut" ]; then
    mkdir -p /var/cache/festivals-fileserver >/dev/null || { echo "Failed to create cache directory. Exiting." ; exit 1; }
    curl --progress-bar -o "$goOut" "$goURL" || { echo "Failed to download go. Exiting." ; exit 1; }
  else
    echo "Using cached go package at $goOut"
    sleep 1
  fi

  tar -C /usr/local -xf "$goOut"
  ln -sf /usr/local/go/bin/* /usr/local/bin
  echo "Installed go ($currentGo)"
  sleep 1
fi

# Install git if needed.
#
if ! command -v git > /dev/null; then
  if command -v dnf > /dev/null; then
    echo "Installing git"
    dnf install git -y > /dev/null;
  elif command -v apt > /dev/null; then
    echo "Installing git"
    apt install git -y > /dev/null;
  else
    echo "Unable to install git. Exiting."
    sleep 1
    exit 1
  fi
else
  echo "Already installed git"
fi

# Install festivals-fileserver to /usr/local/bin/festivals-fileserver. TODO: Maybe just link to /usr/local/bin?
#
echo "Downloading current festivals-fileserver..."
yes | sudo git clone https://github.com/Festivals-App/festivals-fileserver.git /usr/local/festivals-fileserver > /dev/null;
cd /usr/local/festivals-fileserver || { echo "Failed to access working directory. Exiting." ; exit 1; }
/usr/local/bin/go build main.go
mv main /usr/local/bin/festivals-fileserver || { echo "Failed to install festivals-fileserver binary. Exiting." ; exit 1; }
if command -v restorecon > /dev/null; then
  restorecon -v /usr/local/bin/festivals-fileserver >/dev/null
fi
mv config_template.toml /etc/festivals-fileserver.conf
mkdir -p /srv/festivals-fileserver/images/resized >/dev/null || { echo "Failed to create the image directories. Exiting." ; exit 1; }
echo "Installed festivals-fileserver."
sleep 1

# Install systemd service
#
if command -v systemctl > /dev/null; then

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