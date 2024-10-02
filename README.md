<h4 align="center">A Go script to select and apply a DNS server in Linux.</h4>
<p align="center">
  <a href="#installation"><img src="https://img.shields.io/badge/Install-blue?style=for-the-badge" alt="Install"></a>
  <a href="#usage"><img src="https://img.shields.io/badge/Usage-green?style=for-the-badge" alt="Usage"></a>
  <a href="#dns-support"><img src="https://img.shields.io/badge/DNS--Support-orange?style=for-the-badge" alt="DNS-Support"></a>
  <a href="#contributing"><img src="https://img.shields.io/badge/Contributing-yellow?style=for-the-badge" alt="Contributing"></a>
</p>

## Dns-Support

![Dns-changer](https://github.com/user-attachments/assets/af03f351-0922-4fd6-98c2-265781c60c5b)

## Installation

```bash
go install github.com/mamad-1999/dns-changer@latest
```
> [!IMPORTANT]
> Backup: It's a good idea to back up your existing `/etc/resolv.conf` before running the script:
> ```bash
> sudo cp /etc/resolv.conf /etc/resolv.conf.backup
> ```

## Usage
   
```bash
dns-changer
```

> [!TIP]
> You can add custom or other DNS entries in the `config.json` file.
> 
> After the first use, `config.json` will be downloaded and saved in your home directory.

> [!NOTE]
> **Network Manager Configuration**: If you find that your changes are being overwritten, you may need to configure your network manager to stop updating /etc/resolv.conf.
For example:
#### Option 1:
- NetworkManager: Edit `/etc/NetworkManager/NetworkManager.conf` and set `dns=none` under the `[main]` section.
- systemd-resolved: You might need to disable it or configure it to use the DNS servers you want.

After making changes to network manager configurations, restart the service:

```bash
sudo systemctl restart NetworkManager
```
or
```bash
sudo systemctl restart systemd-resolved
```
#### Option 2
- Just do this:
```bash
unlink /etc/resolv.conf 
```
## Contributing

Contributions to Fallparams are welcome! 

