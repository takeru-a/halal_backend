# halal_backend

# SetUp
## /src直下に.env追加を追加
``` 
POSTGRES_USER=hackU
POSTGRES_PASSWORD=yahoo
POSTGRES_DB=halal-db
```
## /src直下に sqlboiler.tomlを追加
``` 
output   = "./db/models"
wipe     = true
no-tests = true
add-enum-types = true

[psql]
  dbname = "halal-db"
  host = "halalDB"
  port = 5432
  user = "hackU"
  pass = "yahoo"
  sslmode = "disable"
  blacklist = [""]
```



# Docker
build　(初回) 
```
docker-compose build
```
起動
```
docker-compose up -d 
```

# Migration
DBにsqlファイルを基にテーブルを作成
```
sql-migrate up
```

# SQLBoiler
テーブルの情報を元に構造体が作成
```
sqlboiler psql
```

# API
- /shops  GET
  - すべての店の情報を取得
- /shops   POST
  - 店の情報を登録
- /shops/id  GET
  - IDがidの店の情報を取得



# 参考資料
- https://zenn.dev/10inoino/articles/57e137f0914330
  - Ginのチュートリアルから派生させて、PostgreSQLへの保存まで
- SQL-Migrate
  - https://github.com/rubenv/sql-migrate
- SQL Boiler
  - https://github.com/volatiletech/sqlboiler
