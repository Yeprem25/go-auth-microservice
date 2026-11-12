data "external_schema" "gorm" {
  program = [
    "go", "run", "-mod=mod", "./auth-service/cmd/atlas-loader",
  ]
}

env "local" {
  src = data.external_schema.gorm.url
  dev = "docker://postgres/17/dev?search_path=public"

  migration {
    dir = "file://auth-service/migrations?format=golang-migrate"
  }
}