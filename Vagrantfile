# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure(2) do |config|
  config.vm.box = "debian/jessie64"
  config.vm.network "forwarded_port", guest: 4242, host: 4242

  config.vm.provision "shell", inline: <<-SHELL
end
