# Bootstrap EC2 scripts to install Jenkins 
 #!/bin/bash #specifies the interpreter
sudo yum update
sudo wget -O /etc/yum.repos.d/jenkins.repo https://pkg.jenkins.io/redhat-stable/jenkins.repo 
sudo rpm --import https://pkg.jenkins.io/redhat-stable/jenkins.io-2023.key  #imports the GPG key for the Jenkins repositor
sudo yum upgrade -y #  upgrades packages again, which might be necessary to ensure that all deps are installed
sudo dnf install java-11-amazon-corretto -y  # installs Amazon Corretto 11, a required dep of Jenkins
sudo yum install jenkins -y  # installs jenkins

# enables and starts jenkins service
sudo systemctl enable jenkins  
sudo systemctl start jenkins  