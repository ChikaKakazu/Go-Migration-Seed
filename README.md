# Go-Migration-Seed

- [golang-migrate](https://github.com/golang-migrate/migrate)
- [SQLBoiler](https://github.com/volatiletech/sqlboiler)

## Migration

**マイグレーションファイルの作成**

```
# migrate create -ext sql -dir ファイルの生成先 -seq create_テーブル名

docker exec -ti app migrate create -ext sql -dir database/migrations -seq create_users
```

**テーブルの作成**

```
# migrate -database="接続先" -path=マイグレーションファイルのパス up

docker exec -ti app migrate -database="mysql://user:password@tcp(mysql:3306)/local" -path=database/migrations/ up
```

**テーブルの削除**

```
# migrate -database="接続先" -path=マイグレーションファイルのパス down

docker exec -ti app migrate -database="mysql://user:password@tcp(mysql:3306)/local" -path=database/migrations/ down
```

## SQLBoiler (ORM)

**コードの自動生成**

```
docker exec -ti app sqlboiler mysql -c sqlboiler.toml -o models -p models --no-tests
```

## Seed

1. `database/seeds/csv`配下に入れ込みたいデータをcsvファイルで作成する
2. `database/seeds/seedModels`配下に作成したテーブルのファイルを作成し、そこに`Create()`と`Delete()`を作成する
3. `database/seeds/seedModels`配下のseed.goに追加する
4. 以下コマンドを実行する
    ```
    docker exec -ti app go run database/seeds/seeder.go
    ```