# WeGYB
## Pre image
1. Copy files (`cp -n ..`) to rootfs
1. remove resize `init=/usr/lib/raspi-config/init_resize.sh` from cmdline.txt
1. Resize image to 2.7GB

## Run in new image
1. Set Wifi Country in raspi-config
1. Update and Install
```bash
sudo apt-get update && sudo apt-get install -y hostapd dnsmasq
```
1. Configure static IP
```bash
sudo bash -c "cat >> /etc/dhcpcd.conf << EOF

interface wlan0
static ip_address=192.168.5.10/24
denyinterfaces eth0
denyinterfaces wlan0
EOF"
```
1. Configure the DHCP server
```bash
sudo mv /etc/dnsmasq.conf /etc/dnsmasq.conf.orig
sudo bash -c "cat >> /etc/dnsmasq.conf << EOF
interface=wlan0
dhcp-range=192.168.5.11,192.168.5.30,255.255.255.0,24h

domain-needed

no-resolv

local=/app/
domain=app
EOF"
sudo bash -c "echo '192.168.5.10	wegyb.app' >> /etc/hosts"
```

1. Configure the access point host software
```bash
sudo bash -c "cat > /etc/hostapd/hostapd.conf << EOF
interface=wlan0
hw_mode=g
channel=7
wmm_enabled=1
macaddr_acl=0
auth_algs=1
ignore_broadcast_ssid=0
wpa=2
wpa_key_mgmt=WPA-PSK
wpa_pairwise=TKIP
rsn_pairwise=CCMP
ssid=WeGotYourBack
wpa_passphrase=
EOF"
sudo systemctl unmask hostapd
sudo systemctl disable wpa_supplicant.service
```
1. Setup wegyb app folder (use env file so we can update env without changing service file)
```bash
sudo mkdir -p /var/lib/wegyb/video
sudo bash -c "cat > /var/lib/wegyb/conf.env << EOF
WEGYB_HOST=0.0.0.0
WEGYB_OUTPUT=/var/lib/wegyb/video
EOF"
```
1. Setup wegyb
```bash
sudo bash -c "cat > /etc/systemd/system/wegyb.service << EOF
[Unit]
Description=WeGYB
After=network.target

[Service]
Type=simple
Restart=always
RestartSec=3
User=root
EnvironmentFile=/var/lib/wegyb/conf.env
ExecStart=/usr/bin/wegyb s

[Install]
WantedBy=multi-user.target
EOF"
sudo systemctl --system daemon-reload
sudo systemctl enable wegyb.service
```

## Backup
1. Backup image
```sudo dd if=/dev/sdb of=wegyb.img bs=512 count=5273601
$ sudo fdisk -l /dev/sdb
Disk /dev/sdb: 59.7 GiB, 64088965120 bytes, 125173760 sectors
Disk model: STORAGE DEVICE  
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 512 bytes
I/O size (minimum/optimal): 512 bytes / 512 bytes
Disklabel type: dos
Disk identifier: 0x7d5a2870

Device     Boot  Start     End Sectors  Size Id Type
/dev/sdb1         8192  532479  524288  256M  c W95 FAT32 (LBA)
/dev/sdb2       532480 5806079 5273600  2.5G 83 Linux
$ sudo dd if=/dev/sdb of=wegyb.img bs=512 count=5806080
```
2. Resize image
https://thepi.io/how-to-use-your-raspberry-pi-as-a-wireless-access-point/
