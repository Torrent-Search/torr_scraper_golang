#!/bin/sh
rm -f socks5.txt
wget https://raw.githubusercontent.com/TheSpeedX/PROXY-List/master/socks5.txt
export pr1=$(sed -n 1p socks5.txt)
export pr2=$(sed -n 2p socks5.txt)