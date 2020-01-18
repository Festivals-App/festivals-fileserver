# VPS Settings
	
Name:       - na -
IP4:        - na -
IP6:        - na -
    
Username:   - na -
Password:   - na -

### Commands:    
``
    sudo reboot
    ufw status // check firewall status
    top // displays system resources
``


## GO Settings

GOPATH:     - na -


## NGINX Settings
	
Config:     /etc/nginx/sites-available/...

### Commands:	
``
    sudo nginx -t // test configuration
    sudo service nginx start
    sudo service nginx stop
    sudo service nginx restart
    sudo service nginx status
``


## EVENTUSFILESERVER Settings

Port:       - na -
Host:		128.0.0.1

### Commands:    
``
    sudo nano /lib/systemd/system/eventusfileserver.service // create service
    sudo service eventusfileserver start
    sudo service eventusfileserver stop
    sudo service eventusfileserver restart
    sudo service eventusfileserver status
``


## Resources

Setup Server&SSH:
- https://www.digitalocean.com/community/tutorials/initial-server-setup-with-ubuntu-18-04
Setup NGINX
- https://www.digitalocean.com/community/tutorials/how-to-install-nginx-on-ubuntu-18-04
Setup SSL:
- https://www.digitalocean.com/community/tutorials/how-to-secure-nginx-with-let-s-encrypt-on-ubuntu-18-04
Setup GO:
- https://www.digitalocean.com/community/tutorials/how-to-install-go-and-set-up-a-local-programming-environment-on-ubuntu-18-04


## Remarks:

	- Remove default NGINX config from /etc/nginx/sites-enabled
	- 
