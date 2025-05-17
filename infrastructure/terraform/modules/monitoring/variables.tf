variable "environment" {
  description = "Environment name (e.g., dev, sit, uat, stg, prd)"
  type        = string
}

variable "vpc_id" {
  description = "ID of the VPC"
  type        = string
}

variable "private_subnet_ids" {
  description = "List of private subnet IDs"
  type        = list(string)
}

variable "ecs_cluster_name" {
  description = "Name of the ECS cluster"
  type        = string
}

variable "ecs_service_names" {
  description = "Names of the ECS services"
  type        = list(string)
}

variable "rds_identifier" {
  description = "Identifier of the RDS instance"
  type        = string
}

variable "aws_region" {
  description = "AWS region"
  type        = string
  default     = "us-east-1"
}
