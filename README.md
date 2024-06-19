# native-devenv

## Set started
1. Install direnv `curl -sfL https://direnv.net/install.sh | bash` and hook it to your shell rc
2. Install devbox `curl -fsSL https://get.jetify.com/devbox | bash`
3. Run `devbox install` to download and install all the dependencies
4. Run `devbox up` to bring up all services

## To do list
* Multiple redis instances for classic and platform
* Update Nginx reverse proxy
* Unleash
* SSL proxy and cert gen https://www.nixhub.io/packages/ssl-proxy
* ...

## TBD
* dnsmasq for supporting wildcard subdomain to avoid editing `/etc/hosts` file
