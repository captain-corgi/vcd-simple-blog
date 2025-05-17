variable "environment" {
  description = "Environment name (e.g., dev, sit, uat, stg, prd)"
  type        = string
}

variable "alb_domain_name" {
  description = "Domain name of the ALB"
  type        = string
}

variable "certificate_arn" {
  description = "ARN of the SSL certificate for CloudFront"
  type        = string
  default     = ""
}

variable "domain_name" {
  description = "Domain name for the application"
  type        = string
  default     = ""
}
