env "local" {
  url = env("DB_URL")
}

migration {
  dir = "db/migrations"
  format = sql
}
