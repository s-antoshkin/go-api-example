```bash
Example:
$ psql -p 5432 -h localhost -U postgres -c "CREATE DATABASE phonebook"
$ psql -h localhost -p 5432 -d phonebook -U postgres -f phonebook.sql
```