## 01.Dockerコマンドについて
### Dockerコンテナの起動と停止
```docker:docker.rb
$ docker compose up
$ docker compose down
```
### コンテナ内部に入る
```docker:docker.rb
$ docker compose exec サービス名 bash
```

## 02.MySQLに接続する
### コンテナ起動及びコンテナ内部に入る
```docker:docker.rb
$ docker compose up -d && docker compose exec db bash
# mysql -u root -p
```

### コマンド一覧
```sql:sql.rb
# DB一覧を確認する
show databases;

# DBを作成する
create database runteq_study;
# railsでの作成の場合
rails db:create

# 作成したDBを選択して使用する
use runteq_study;

# テーブルの確認
show tables

# テーブルの作成
CREATE TABLE user(
    id INT(11) AUTO_INCREMENT NOT NULL,
    first_name VARCHAR(30) NOT NULL,
    last_name VARCHAR(30) NOT NULL,
    PRIMARY KEY (id));
# railsの場合
$ rails g model user first_name:string last_name:string
$ rails db:migrate

# insert文でのデータ作成
insert into user (first_name, last_name) values ('rantekun', 'shibuya');
# selectで確認
select * from user;
# railsの場合
user = User.new(first_name: 'rantekun', last_name: 'shibuya')
user.save

# ２つ目
CREATE TABLE board(
    id INT(11) AUTO_INCREMENT NOT NULL,
    title VARCHAR(30) NOT NULL,
    body VARCHAR(30) NOT NULL,
    user_id INT(11) NOT NULL,
    PRIMARY KEY (id),
    INDEX (user_id),
    FOREIGN KEY (user_id)
      REFERENCES user(id));

# whereで条件を指定して絞り込む
select * from board where user_id = 1;
# rails
Board.where(user_id: 1)

# 結合（join）
9レコード（3 * 3 = 9）
select * from user inner join board;

# on（結合条件を指定する）
select * from user inner join board on user.id = board.user_id;

# update
update user set last_name="japan" where id = 3;

# delete
delete from user where id = 3;




```
