echo "Install app..."
sudo cp -rf bin/minibinmac /opt/local/bin
sudo cp -rf bin/AutostartMiniBinMac /opt/local/bin
sudo cp -rf bin/minibinmac.sh /opt/local/bin
sudo cp -rf bin/minibinmac.sh /usr/local/bin
sudo cp -rf bin/AutostartMiniBinMac /usr/local/bin
sudo chmod +x bin/minibinmac.sh
sudo chmod +x /usr/local/bin/minibinmac.sh
sudo cp -rf bin/minibinmac /usr/local/bin
echo "Done..."

# echo 'export="/opt/local/bin"' > ~/.zshrc