output "alb_domain_name" {
  description = "Domain name of the ALB"
  value       = aws_lb.main.dns_name
}

output "cluster_name" {
  description = "Name of the ECS cluster"
  value       = aws_ecs_cluster.main.name
}

output "service_names" {
  description = "Names of the ECS services"
  value       = [
    aws_ecs_service.api_gateway.name,
    aws_ecs_service.frontend.name
  ]
}
