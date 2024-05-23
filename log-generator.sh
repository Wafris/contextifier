#!/bin/bash

# List of IP addresses
ips=(23.92.20.209 87.9.123.227 95.251.141.109 45.26.76.85 71.223.88.183 174.108.33.53 94.253.88.115 104.222.16.99 94.13.230.164 176.105.216.44 51.89.138.158 181.117.13.51 66.63.167.173 177.54.123.165 142.162.46.76 37.187.151.123 187.95.220.144 89.245.183.140 63.76.251.51 138.64.226.142 77.77.129.246 212.205.0.30 189.34.247.144 86.209.92.40 80.187.119.235)

# Methods array
methods=(GET POST DELETE PUT)

# URIs array
uris=("/home" "/api/data" "/user/profile" "/about")

# Protocols array
protocols=("HTTP/1.1" "HTTP/2.0")

# Status codes array
statuses=(200 301 404 500)

# Log file
logfile="http_log.txt"

# Check if logfile can be written
touch $logfile
if [ $? -ne 0 ]; then
    echo "Failed to create log file. Check permissions or path."
    exit 1
fi

# Generate 100 log entries
for i in {1..100}; do
  ip=${ips[$RANDOM % ${#ips[@]}]}
  method=${methods[$RANDOM % ${#methods[@]}]}
  uri=${uris[$RANDOM % ${#uris[@]}]}
  protocol=${protocols[$RANDOM % ${#protocols[@]}]}
  status=${statuses[$RANDOM % ${#statuses[@]}]}
  size=$((RANDOM % 5000 + 100))
  timestamp=$(date +"%d/%b/%Y:%H:%M:%S +0000")

  # Format: IP user-identifier userid [timestamp] "method URI protocol" status size
  echo "$ip - - [$timestamp] \"$method $uri $protocol\" $status $size" >> $logfile
done

echo "Generated HTTP log data in $logfile"
