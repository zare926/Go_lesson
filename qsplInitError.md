# postgresを導入して起動したかったけどできなかった

macなのでbrewでインストール
```
$ brew install postgresql
```
文字コードをUTF-8でデータベースの初期化
```
$ initdb /usr/local/var/postgres --encoding=UTF-8 --locale=ja_JP.UTF-8

The files belonging to this database system will be owned by user "username".
This user must also own the server process.

The database cluster will be initialized with locale "ja_JP.UTF-8".
initdb: could not find suitable text search configuration for locale "ja_JP.UTF-8"
The default text search configuration will be set to "simple".

Data page checksums are disabled.

initdb: directory "/usr/local/var/postgres" exists but is not empty
If you want to create a new database system, either remove or empty
the directory "/usr/local/var/postgres" or run initdb
with an argument other than "/usr/local/var/postgres".

```
で、ここからDBの起動をしてエンコードを確認したかったのにこれ
```
$ postgres -D /usr/local/var/postgres
2021-12-28 08:55:23.530 JST [63749] FATAL:  lock file "postmaster.pid" already exists
2021-12-28 08:55:23.530 JST [63749] HINT:  Is another postmaster (PID 63717) running in data directory "/usr/local/var/postgres"?
```
https://zenn.dev/tetsuya/articles/d5e0aef6f7144f<br>
ここを参照してみたものの治らず

```
$ psql 
psql: error: connection to server on socket "/tmp/.s.PGSQL.5432" failed: FATAL:  database "username" does not exist
```

と思ったらデータベース作ってないので、最初は上記の文言が出るようです。
```
$ psql -l
```
でエンコードも確認ができました。