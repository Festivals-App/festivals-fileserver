#!/bin/bash
#
# install.sh - FestivalsApp File Server Installer Script
#
# Enables the firewall, installs the latest version of the FestialsApp File Server, starts it as a service.
#
# (c)2020-2025 Simon Gaus
#

# ─────────────────────────────────────────────────────────────────────────────
# 🔍 Detect Web Server User
# ─────────────────────────────────────────────────────────────────────────────
WEB_USER="www-data"
if ! id -u "$WEB_USER" &>/dev/null; then
    WEB_USER="www"
    if ! id -u "$WEB_USER" &>/dev/null; then
        echo -e "\n\033[1;31m❌  ERROR: Web server user not found! Exiting.\033[0m\n"
        exit 1
    fi
fi

echo -e "\n👤  Web server user detected: \e[1;34m$WEB_USER\e[0m"
sleep 1

# ─────────────────────────────────────────────────────────────────────────────
# 📁 Setup Working Directory
# ─────────────────────────────────────────────────────────────────────────────
WORK_DIR="/usr/local/festivals-fileserver/install"
mkdir -p "$WORK_DIR" && cd "$WORK_DIR" || { echo -e "\n\033[1;31m❌  ERROR: Failed to create/access working directory!\033[0m\n"; exit 1; }

echo -e "\n📂  Working directory set to \e[1;34m$WORK_DIR\e[0m\n"
sleep 1

# ─────────────────────────────────────────────────────────────────────────────
# 🖥  Detect System OS and Architecture
# ─────────────────────────────────────────────────────────────────────────────

echo -e "\n\n\n🔍  Detecting system OS and architecture..."
sleep 1

if [ "$(uname -s)" = "Darwin" ]; then
    os="darwin"
elif [ "$(uname -s)" = "Linux" ]; then
    os="linux"
else
    echo -e "\n🚨  ERROR: Unsupported OS. Exiting.\n"
    exit 1
fi

if [ "$(uname -m)" = "x86_64" ]; then
    arch="amd64"
elif [ "$(uname -m)" = "arm64" ]; then
    arch="arm64"
else
    echo -e "\n🚨  ERROR: Unsupported CPU architecture. Exiting.\n"
    exit 1
fi

echo -e "\n✅  Detected OS: \e[1;34m$os\e[0m, Architecture: \e[1;34m$arch\e[0m."
sleep 1

# ─────────────────────────────────────────────────────────────────────────────
# 📦 Install FestivalsApp File Server
# ─────────────────────────────────────────────────────────────────────────────

file_url="https://github.com/Festivals-App/festivals-fileserver/releases/latest/download/festivals-fileserver-$os-$arch.tar.gz"

echo -e "\n📥  Downloading latest FestivalsApp File Server binary..."
curl --progress-bar -L "$file_url" -o festivals-fileserver.tar.gz

echo -e "\n📦  Extracting binary..."
tar -xf festivals-fileserver.tar.gz

mv festivals-fileserver /usr/local/bin/festivals-fileserver || {
    echo -e "\n🚨  ERROR: Failed to install FestivalsApp File Server binary. Exiting.\n"
    exit 1
}

echo -e "\n✅  Installed FestivalsApp File Server to \e[1;34m/usr/local/bin/festivals-fileserver\e[0m.\n"
sleep 1

# ─────────────────────────────────────────────────────────────────────────────
# 🛠  Install Server Configuration File
# ─────────────────────────────────────────────────────────────────────────────

echo -e "\n\n\n📂  Moving default configuration file..."
mv config_template.toml /etc/festivals-fileserver.conf

if [ -f "/etc/festivals-fileserver.conf" ]; then
    echo -e "\n✅  Configuration file moved to \e[1;34m/etc/festivals-fileserver.conf\e[0m.\n"
else
    echo -e "\n🚨  ERROR: Failed to move configuration file. Exiting.\n"
    exit 1
fi
sleep 1

# ─────────────────────────────────────────────────────────────────────────────
# 📂  Prepare Log Directory
# ─────────────────────────────────────────────────────────────────────────────

echo -e "\n\n\n📁  Creating log directory..."
mkdir -p /var/log/festivals-fileserver || {
    echo -e "\n🚨  ERROR: Failed to create log directory. Exiting.\n"
    exit 1
}

echo -e "\n✅  Log directory created at \e[1;34m/var/log/festivals-fileserver\e[0m.\n"
sleep 1

# ─────────────────────────────────────────────────────────────────────────────
# 📂  Prepare File Storage Directories
# ─────────────────────────────────────────────────────────────────────────────

echo -e "\n\n\n📁  Creating file storage directories..."
mkdir -p /srv/festivals-fileserver/images/resized || {
    echo -e "\n🚨  ERROR: Failed to create image directory. Exiting.\n"
    exit 1
}
mkdir -p /srv/festivals-fileserver/pdf || {
    echo -e "\n🚨  ERROR: Failed to create pdf directory. Exiting.\n"
    exit 1
}
+
echo -e "\n✅  File storage directories created at \e[1;34m/srv/festivals-fileserver\e[0m.\n"
sleep 1

# ─────────────────────────────────────────────────────────────────────────────
# 🔄 Prepare Remote Update Workflow
# ─────────────────────────────────────────────────────────────────────────────

echo -e "\n\n\n⚙️  Preparing remote update workflow..."
sleep 1

mv update.sh /usr/local/festivals-fileserver/update.sh
chmod +x /usr/local/festivals-fileserver/update.sh

cp /etc/sudoers /tmp/sudoers.bak
echo "$WEB_USER ALL = (ALL) NOPASSWD: /usr/local/festivals-fileserver/update.sh" >> /tmp/sudoers.bak

# Validate and replace sudoers file if syntax is correct
if visudo -cf /tmp/sudoers.bak &>/dev/null; then
    sudo cp /tmp/sudoers.bak /etc/sudoers
    echo -e "\n✅  Updated sudoers file successfully."
else
    echo -e "\n🚨  ERROR: Could not modify /etc/sudoers file. Please do this manually. Exiting.\n"
    exit 1
fi
sleep 1

# ─────────────────────────────────────────────────────────────────────────────
# 🔥 Enable and Configure Firewall
# ─────────────────────────────────────────────────────────────────────────────

if command -v ufw > /dev/null; then
    echo -e "\n\n\n🚀  Configuring UFW firewall..."
    mv ufw_app_profile /etc/ufw/applications.d/festivals-fileserver
    ufw allow festivals-fileserver >/dev/null
    echo -e "\n✅  Added festivals-fileserver to UFW with port 1910."
    sleep 1
elif ! [ "$(uname -s)" = "Darwin" ]; then
    echo -e "\n🚨  ERROR: No firewall detected and not on macOS. Exiting.\n"
    exit 1
fi

# ─────────────────────────────────────────────────────────────────────────────
# ⚙️  Install Systemd Service
# ─────────────────────────────────────────────────────────────────────────────

if command -v service > /dev/null; then
    echo -e "\n\n\n🚀  Configuring systemd service..."
    if ! [ -f "/etc/systemd/system/festivals-fileserver.service" ]; then
        mv service_template.service /etc/systemd/system/festivals-fileserver.service
        echo -e "\n✅  Created systemd service configuration."
        sleep 1
    fi
    systemctl enable festivals-fileserver > /dev/null
    echo -e "\n✅  Enabled systemd service for FestivalsApp File Server."
    sleep 1
elif ! [ "$(uname -s)" = "Darwin" ]; then
    echo -e "\n🚨  ERROR: Systemd is missing and not on macOS. Exiting.\n"
    exit 1
fi

# ─────────────────────────────────────────────────────────────────────────────
# 🔑 Set Appropriate Permissions
# ─────────────────────────────────────────────────────────────────────────────

echo -e "\n\n\n🔑  Setting appropriate permissions..."
sleep 1

chown -R "$WEB_USER":"$WEB_USER" /usr/local/festivals-fileserver
chown -R "$WEB_USER":"$WEB_USER" /var/log/festivals-fileserver
chown -R "$WEB_USER":"$WEB_USER" /srv/festivals-fileserver
chown "$WEB_USER":"$WEB_USER" /etc/festivals-fileserver.conf

echo -e "\n✅  Set Appropriate Permissions.\n"
sleep 1

# ─────────────────────────────────────────────────────────────────────────────
# 🧹 Cleanup Installation Files
# ─────────────────────────────────────────────────────────────────────────────

echo -e "\n🧹  Cleaning up installation files..."
cd /usr/local/festivals-fileserver || exit
rm -R /usr/local/festivals-fileserver/install
sleep 1

# ─────────────────────────────────────────────────────────────────────────────
# 🎉 Final Message
# ─────────────────────────────────────────────────────────────────────────────

echo -e "\n\n\n\n\033[1;32m══════════════════════════════════════════════════════════════════════════\033[0m"
echo -e "\033[1;32m✅  INSTALLATION COMPLETE! 🚀\033[0m"
echo -e "\033[1;32m══════════════════════════════════════════════════════════════════════════\033[0m"

echo -e "\n🔹 \033[1;34mTo start the server manually, run:\033[0m"
echo -e "\n   \033[1;32msudo systemctl start festivals-fileserver\033[0m"

echo -e "\n📂 \033[1;34mBefore starting, update the configuration file at:\033[0m"
echo -e "\n   \033[1;34m/etc/festivals-fileserver.conf\033[0m"

echo -e "\n\033[1;32m══════════════════════════════════════════════════════════════════════════\033[0m\n"
sleep 1
