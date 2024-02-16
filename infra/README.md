# Infrastructure

## [REMOVED SUPPORT] Experimentation with provisioning Jenkins on the cloud

> Reason for removal: ec2 upkeep cost exceeding free-tier budget

Automated provision of cloud resources using Packer + Terraform

### Features

- Custom AMIs baked using Packer (groovy configurations + bash)
- Jenkins Master-Workers Architecture
  - Declarative configuration
  - Automated deployment of cluster and scaling with ASG

### Baking Custom Jenkins Enabled AMIs with Packer

**Credit**: The AMIs used in this project were built following the specifications in https://github.com/tailwarden/virtual-workshops/tree/master/jenkins-cluster-on-aws/packer

Hashicorp's Packer utility bakes custom Amazon Linux AMI to enable the automated deployment of Master and Workers Jenkins EC2 instances.
