Развернуть БД локально:
```bash
docker stack deploy -c docker-compose.local.yml -c env.development.yml alef
```

Переменные окружения в файле env.{environment}.yml:

```yml
version: '3.8'

services:
  alefdb:
    environment:
      - POSTGRES_USER={user_name}
      - POSTGRES_PASSWORD={password}

  server:
    environment:
      - APP_ENV={environment}
      - DB_HOST={database_host}
      - DB_USER={database_user}
      - DB_PASSWORD={database_password}
      - DB_NAME={database_name}
```