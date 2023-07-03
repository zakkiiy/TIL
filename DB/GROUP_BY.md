## GROUP BY
- 複数のデータごとに集計したいときに使用する。

### 使用するデータ

|  name  |  created_day  |  channel  |  age  |
| ------ | ------------- | --------- | ------|
|  高橋  |  2021-02-13  |  web |  27  |
|  山田  |  2021-02-13  |  ad  |  27  |
|  美山  |  2021-02-15  |  ad  |  27  |
|  佐藤  |  2021-02-15  |  ad  |  33  |
|  丸岡  |  2021-02-16  |  web |  24  |

### 01.今週の日毎の会員登録数
```sql
SELECT created_day, count(name)
FROM members
GROUP BY created_day;
```

|  created_day  |  count(name)  | 
| ------------- | --------- |
|  2021-02-13  |  2 |
|  2021-02-15  |  2 |
|  2021-02-16  |  1 |

### 02.日毎の会員登録数をチャンネルごとに出す。
```sql
SELECT created_day, channel, count(name)
FROM members
GROUP BY created_day;
```

- エラーになる。
- adが入るかwebが入るか不明のためエラー発生。

|  created_day  |  channel  | 
| ------------- | --------- |
|  2021-02-13  |  ? |
|  2021-02-15  |  ad |
|  2021-02-16  |  web |

- 軸になるカラムを全てGROUP BYで指定する。
```sql
SELECT created_day, channel, count(name)
FROM members
GROUP BY created_day channel;
```
|  created_day  |  channel  |  channel  | 
| ------------- | --------- | --------- |
|  2021-02-13  |  web |  1 |
|  2021-02-13  |  ad |  2 |
|  2021-02-15  |  ad |  1 |
|  2021-02-16  |  web |  1 |

### 03.日毎の会員登録者の平均年齢と最大年齢を出す。

```sql
SELECT created_day, AVG(age), MAX(age)
FROM members
GROUP BY created_day;
```
|  created_day  |  avg(age)  |  max(age)  | 
| ------------- | --------- | --------- |
|  2021-02-13  |  27 |  27 |
|  2021-02-15  |  30 |  33 |
|  2021-02-16  |  24 |  24 |
