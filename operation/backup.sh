#!/bin/bash
#
# backup.sh - FestivalsApp File Server Backup Script
#
# Create a backup of all images in the callers home folder. 
#
# (c)2025 Simon Gaus
#

# ─────────────────────────────────────────────────────────────────────────────
# 🖥  Archive and zip image folder
# ─────────────────────────────────────────────────────────────────────────────
tar -zcvf ~/backup.tar.gz /srv/festivals-fileserver/images/