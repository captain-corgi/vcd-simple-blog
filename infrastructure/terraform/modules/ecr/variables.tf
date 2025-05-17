variable "environment" {
  description = "Environment name (e.g., dev, sit, uat, stg, prd)"
  type        = string
}

variable "repositories" {
  description = "List of ECR repository names to create"
  type        = list(string)
}
