## Comamands

```bash
brew services start postgresql@15
psql -U postgres
\c greenlight
CREATE ROLE greenlight WITH LOGIN PASSWORD '<password>';
CREATE EXTENSION IF NOT EXISTS citext;
psql --host=localhost --dbname=greenlight --username=greenlight

```

## Set up Env Variables

- Create ~/.profile
- Add:

  ```
  export GREENLIGHT_DB_DSN='postgres://greenlight:pa55word@localhost/ greenlight'
  ```

- Run: `source ~/.profile`
- Check: `echo $GREENLIGHT_DB_DSN`
