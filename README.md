# Alchemist Service

## Local Setup

These are the tools I use:

- [Docker Desktop](https://www.docker.com/products/docker-desktop)
- [Git](https://git-scm.com)
- [Go](https://go.dev/)
- [VS Code](https://code.visualstudio.com)

### 1. Clone Code

Clone the GitHub repository locally.

```bash
git clone https://github.com/coding-shenanigans/alchemist-service.git
```

### 2. Create Database

Create a PostgreSQL Docker container.

```bash
docker run --name local-postgres -e POSTGRES_PASSWORD=your_password -p 5432:5432 -d postgres
```

You can set up the database via CLI or the pgAdmin UI. There's a section for each method below.

#### CLI

1. Connect to the PostgreSQL container.

```bash
docker exec -it local-postgres /bin/bash
```

2. Connect to the PostgreSQL instance running in the container.

```bash
psql -U postgres -h localhost -d postgres
```

3. Create a database for the project.

```sql
CREATE DATABASE alchemist;
```

4. Connect to the new database.

```sql
\c alchemist
```

5. Create all the database objects by running the queries in [database.sql](https://github.com/coding-shenanigans/alchemist-service/blob/main/internal/database/database.sql)

#### pgAdmin

1. Create a pgAdmin Docker container.

```bash
docker run --name local-pgadmin4 -e PGADMIN_DEFAULT_EMAIL=your_fake_email -e PGADMIN_DEFAULT_PASSWORD=your_fake_email_password -p 8080:80 -d dpage/pgadmin4
```

2. Log in to the pgAdmin UI.

- http://localhost:8080
- Use the fake email and password used when creating the pdAgmin Docker container.

3. Connect to the PostgreSQL container.

- Right click `Servers` > `Register` > `server...`
- In the `Connection` tab, enter the information for your `local-postgres` container.
  - You may need to set the `Host name/address` field to `host.docker.internal` when using Docker Desktop.

4. Create a database for the project.

- Right click `Databases` > `Create` > `Database...`
  - Enter `alchemist`.

5. Set up database.
   - Right click `alchemist` > `Query Tool`
   - Create all the database objects by running the queries in [database.sql](https://github.com/coding-shenanigans/alchemist-service/blob/main/internal/database/database.sql)

### 3. Run the Service

1. Open the folder download from GitHub in VS Code.
2. Go to the `Run and Debug` tab in VS Code.
3. Select `Alchemist Service Local` and click the play button.

## Run All Tests

```bash
go test ./...
```
