terraform {
  required_version = ">= 1.0.0"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }

  backend "s3" {
    bucket         = "vcd-simple-blog-terraform-state"
    key            = "terraform.tfstate"
    region         = "us-east-1"
    dynamodb_table = "vcd-simple-blog-terraform-locks"
    encrypt        = true
  }
}

provider "aws" {
  region = var.aws_region

  default_tags {
    tags = {
      Project     = "vcd-simple-blog"
      Environment = var.environment
      ManagedBy   = "terraform"
    }
  }
}

module "vpc" {
  source = "./modules/vpc"

  environment = var.environment
  vpc_cidr    = var.vpc_cidr
  azs         = var.availability_zones
}

module "rds" {
  source = "./modules/rds"

  environment      = var.environment
  vpc_id           = module.vpc.vpc_id
  subnet_ids       = module.vpc.private_subnet_ids
  postgres_version = var.postgres_version
  instance_class   = var.db_instance_class
  db_name          = var.db_name
  db_username      = var.db_username
  db_password      = var.db_password
}

module "ecr" {
  source = "./modules/ecr"

  environment = var.environment
  repositories = [
    "api-gateway",
    "auth-service",
    "blog-service",
    "user-service",
    "frontend-shell"
  ]
}

module "ecs" {
  source = "./modules/ecs"

  environment         = var.environment
  vpc_id              = module.vpc.vpc_id
  public_subnet_ids   = module.vpc.public_subnet_ids
  private_subnet_ids  = module.vpc.private_subnet_ids
  ecr_repository_urls = module.ecr.repository_urls
  db_host             = module.rds.db_endpoint
  db_name             = var.db_name
  db_username         = var.db_username
  db_password         = var.db_password
}

module "cloudfront" {
  source = "./modules/cloudfront"

  environment      = var.environment
  alb_domain_name  = module.ecs.alb_domain_name
  certificate_arn  = var.certificate_arn
  domain_name      = var.domain_name
}

module "monitoring" {
  source = "./modules/monitoring"

  environment     = var.environment
  vpc_id          = module.vpc.vpc_id
  private_subnet_ids = module.vpc.private_subnet_ids
  ecs_cluster_name = module.ecs.cluster_name
  ecs_service_names = module.ecs.service_names
  rds_identifier  = module.rds.db_identifier
}
