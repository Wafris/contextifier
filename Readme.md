## Contextifier

### What is this?

If you are investigating an attack on a website you spend a lot of time searching through log files to find the nuggets of useful information from the sea of noise.

This tool extracts IP addresses and frequency from any type of structured or unstructured log file. 

### Usage

```bash
./contextifier [path_to_log_file] [path_to_exclusion_file]
```

#### Example
  
```bash
./contextifier /var/log/apache2/access.log /usr/local/etc/exclude_ranges.txt
```

### Output

The output is a list of IP addresses and the number of times they appear in the log file.  For example:

```bash
66.248.202.37        153        https://wafris.org/ip-lookup/66.248.202.37
185.93.228.37        94         https://wafris.org/ip-lookup/185.93.228.37
185.93.230.37        75         https://wafris.org/ip-lookup/185.93.230.37
192.88.134.37        55         https://wafris.org/ip-lookup/192.88.134.37
185.93.229.37        38         https://wafris.org/ip-lookup/185.93.229.37
208.109.1.37         22         https://wafris.org/ip-lookup/208.109.1.37
66.248.203.37        20         https://wafris.org/ip-lookup/66.248.203.37
66.248.200.37        16         https://wafris.org/ip-lookup/66.248.200.37
192.88.135.37        16         https://wafris.org/ip-lookup/192.88.135.37
109.147.188.246      12         https://wafris.org/ip-lookup/109.147.188.246
97.134.99.54         11         https://wafris.org/ip-lookup/97.134.99.54
209.85.238.171       10         https://wafris.org/ip-lookup/209.85.238.171
```


### Excluding Ranges

The `exclude_ranges.txt` file is a list of CIDR ranges that you want to exclude from the output.  This is useful for removing known good IP addresses from the output, proxies, internal networks, etc.

Add your own CIDR ranges to be excluded one per line.  For example:

#### Akamai IP Ranges

```2.16.0.0/13  
23.0.0.0/12  
23.192.0.0/11  
23.32.0.0/11  
23.64.0.0/14  
23.72.0.0/13  
69.192.0.0/16  
72.246.0.0/15  
88.221.0.0/16  
92.122.0.0/15
95.100.0.0/15  
96.16.0.0/15  
96.6.0.0/15  
104.64.0.0/10  
118.214.0.0/16  
173.222.0.0/15  
184.24.0.0/13  
184.50.0.0/15  
184.84.0.0/14
2a02:26f0::/32  
2600:1400::/24  
2405:9600::/32
```

#### Cloudflare IP Ranges

```173.245.48.0/20
103.21.244.0/22
103.22.200.0/22
103.31.4.0/22
141.101.64.0/18
108.162.192.0/18
190.93.240.0/20
188.114.96.0/20
197.234.240.0/22
198.41.128.0/17
162.158.0.0/15
104.16.0.0/13
104.24.0.0/14
172.64.0.0/13
131.0.72.0/22
2400:cb00::/32
2606:4700::/32
2803:f800::/32
2405:b500::/32
2405:8100::/32
2a06:98c0::/29
2c0f:f248::/32
```


#### Sucuri IP Ranges

```192.88.134.0/23 
185.93.228.0/22 
66.248.200.0/22 
2a02:fe80::/29 
208.109.0.0/22
```

#### Private IPv4 Ranges

```10.0.0.0/8
172.16.0.0/12
192.168.0.0/16
```


### IP Lookup Integration

Each IP in the report is linked to a comprehensive reputation report on [IP Lookup](https://wafris.org/ip-lookup/) for further investigation.



/*

MIT License

Copyright (c) [2024] [Michael J Buckbee]

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

*/
