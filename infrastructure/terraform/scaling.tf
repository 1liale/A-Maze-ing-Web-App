
# Auto-scaling
resource "aws_autoscaling_group" "jenkins_workers" {
  name                 = "jenkins-workers-asg"
  launch_configuration = aws_launch_configuration.jenkins_workers_launch_conf.name
  
  vpc_zone_identifier  = [aws_subnet.public_subnet.id]

  min_size             = 2
  max_size             = 5

  depends_on = [aws_instance.jenkins_master]

  lifecycle {
    create_before_destroy = true
  }

  tag {
    key                 = "Name"
    value               = "jenkins_worker"
    propagate_at_launch = true
  }
}