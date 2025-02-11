#!/bin/bash
sudo apt-get update
sudo apt-get install -y nginx certbot python3-certbot-nginx

# Copy Nginx configuration files
sudo cp /path/to/docker/nginx/conf.d/*.conf /etc/nginx/conf.d/

# Obtain SSL certificates
sudo certbot --nginx \
  -d real-estate-management.softcelia.com \
  -d db.softcelia.com \
  -d portainer.softcelia.com \
  --non-interactive \
  --agree-tos \
  --email m7firoz@gmail.com

# Reload Nginx
sudo systemctl reload nginx