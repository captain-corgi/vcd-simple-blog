output "repository_urls" {
  description = "URLs of the ECR repositories"
  value = {
    for name, repo in aws_ecr_repository.repo : name => repo.repository_url
  }
}
