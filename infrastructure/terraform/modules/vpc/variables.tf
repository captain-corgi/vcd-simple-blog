variable "environment" {
  description = "Environment name (e.g., dev, sit, uat, stg, prd)"
  type        = string
}

variable "vpc_cidr" {
  description = "CIDR block for the VPC"
  type        = string
}

variable "azs" {
  description = "List of availability zones to use"
  type        = list(string)
}
