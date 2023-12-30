# Infrastructure
Automated provision of cloud resources using Packer + Terraform

## Features
- Custom AMIs baked using Packer (groovy configurations + bash)
- Jenkins Master-Workers Architecture
    - Declarative configuration
    - Automated deployment of cluster and scaling with ASG

## Baking Custom Jenkins Enabled AMIs with Packer

**Credit**: The setup scripts and security groovy configurations for the Packer section are referenced from https://github.com/tailwarden/virtual-workshops/tree/master/jenkins-cluster-on-aws/packer

> Limitations: Only support creating AMIs for deployment on AWS (i.e not cloud agnostic)

Hashicorp's Packer utility bakes custom Amazon Linux AMI to enable the automated deployment of Master and Workers Jenkins EC2 instances.
