<h4 align="center">A Go script to select and apply a DNS server in Linux.</h4>
<p align="center">
  <a href="#installation"><img src="https://img.shields.io/badge/Install-blue?style=for-the-badge" alt="Install"></a>
  <a href="#usage"><img src="https://img.shields.io/badge/Usage-green?style=for-the-badge" alt="Usage"></a>
  <a href="#dns-support"><img src="https://img.shields.io/badge/DNS--Support-orange?style=for-the-badge" alt="DNS-Support"></a>
  <a href="#contributing"><img src="https://img.shields.io/badge/Contributing-yellow?style=for-the-badge" alt="Contributing"></a>
</p>

## Dns-Support

![Dns-changer](https://github.com/user-attachments/assets/af03f351-0922-4fd6-98c2-265781c60c5b)

----

![2024-12-07_15-38](https://github.com/user-attachments/assets/bbb54355-f782-4dc1-9f2e-4eedee699471)


## Installation

```bash
go install github.com/mamad-1999/dns-changer@latest
```
## Usage
   
```bash
dns-changer
```

> [!TIP]
> You can add custom or other DNS entries in the `config.json` file.
> 
> After the first use, `config.json` will be downloaded and saved in `.config/dns-changer` in home directory.
>
> The script first verifies if the DNS is managed by NetworkManager. If so, it disables NetworkManager's automatic rewriting of the `resolv.conf` file at the start of execution.

## Contributing

Contributions to Fallparams are welcome! 

