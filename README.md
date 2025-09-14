## Comamands

```bash
brew install postgresql@15 
brew services start postgresql@15
psql postgres
CREATE DATABASE greenlight;
\c greenlight
CREATE ROLE greenlight WITH LOGIN PASSWORD '<password>';
CREATE EXTENSION IF NOT EXISTS citext;
psql --host=localhost --dbname=greenlight --username=greenlight

```

## Set up Env Variables

- Create ~/.profile
- Add:

  ```
  export GREENLIGHT_DB_DSN='postgres://greenlight:pa55word@localhost/ greenlight?sslmode=disable'
  ```

- Run: `source ~/.profile`
- Check: `echo $GREENLIGHT_DB_DSN`
- Run to start db: `go run ./cmd/api -db-max-open-conns=50 -db-max-idle-conns=50 -db-max-idle-time=2h30m`

## Issues
- If you get the following error: `error: pq: permission denied for schema public`
  - Run:
    ```sql
    ALTER DATABASE greenlight OWNER TO greenlight;
    GRANT CREATE ON DATABASE greenlight TO greenlight;    
    ```

## Set up SQL migration

- Run:
  ```bash
  brew install golang-migrate
  ```
- Next:
  - `migrate create -seq -ext=.sql -dir=./migrations create_movies_table`
  - `migrate create -seq -ext=.sql -dir=./migrations add_movies_check_constraints`

- Add to `migrations/000001_create_movies_table.up.sql`
  ```sql
  CREATE TABLE IF NOT EXISTS movies (
      id bigserial PRIMARY KEY,  
      created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
      title text NOT NULL,
      year integer NOT NULL,
      runtime integer NOT NULL,
      genres text[] NOT NULL,
      version integer NOT NULL DEFAULT 1
  );  
  ```

- Add to `migrations/000001_create_movies_table.down.sql`
  ```sql
  DROP TABLE IF EXISTS movies;
  ```

- Add to `migrations/000002_add_movies_check_constraints.up.sql`
  ```sql
  ALTER TABLE movies ADD CONSTRAINT movies_runtime_check CHECK (runtime >= 0);

  ALTER TABLE movies ADD CONSTRAINT movies_year_check CHECK (year BETWEEN 1888 AND date_part('year', now()));

  ALTER TABLE movies ADD CONSTRAINT genres_length_check CHECK (array_length(genres, 1) BETWEEN 1 AND 5);
  ```

- Add to `migrations/000002_add_movies_check_constraints.down.sql`
  ```sql
  ALTER TABLE movies DROP CONSTRAINT IF EXISTS movies_runtime_check;

  ALTER TABLE movies DROP CONSTRAINT IF EXISTS movies_year_check;

  ALTER TABLE movies DROP CONSTRAINT IF EXISTS genres_length_check;
  ```

- Run: `migrate -path=./migrations -database=$GREENLIGHT_DB_DSN up`

- Check: `SELECT * FROM schema_migrations;`
  - `version` column indicates the migration files up to (and including) number shown have been executed against the database.
  - `dirty` indicates whether the migration files were cleanly executed without any errors and the SQL statements they contain were successfully applied in full.