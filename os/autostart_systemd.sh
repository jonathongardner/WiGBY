#!/bin/bash
cp wegyb@pi.service /etc/systemd/system/wegyb@${SUDO_USER:-${USER}}.service
systemctl --system daemon-reload
systemctl enable wegyb@${SUDO_USER:-${USER}}.service
