## テーブル結合
#### 複数のテーブルをまとめること

## 内部結合（INNER JOIN）
### 両方のテーブルにデータがあり、合体できるデータのみを取り出す。

- usersテーブル

|  id  |  name  |  company_id  |
| ---- | ------ | ------------ |
|  1  |  高橋  |  1  |
|  2  |  山田  |  2  |
|  3  |  美山  |  4  |

- companiesテーブル

|  id  |  name  |
| ---- | ------ |
|  1  |  パナ  |
|  2  |  ソニ  |
|  3  |  日立山  |

- 内部結合結果

```sql
SELECT * FROM users
INNER JOIN companies
      ON users.company_id = companies.id;
```

|  id  |  name  |  company_id  |  name  |  name  |
| ---- | ---- | ---- | ---- | ---- |
|  1  | 高橋  |  1  |  1  |  パナ  |
|  2  | 山田  |  2  |  2  |  ソニ  |

## 外部結合（OUTER JOIN） どちらかのテーブルにデータがあれば取り出す
- まず比較して一緒のものを取り出す。比較してなかったものも取り出す。

### LEFT OUTER JOIN  左側のテーブルが軸
- まずusersテーブルのキーを取ってきて、合致するcompaniesのidも取ってくる。
- キーが同じデータがない場合でもusersは取り出す。ない箇所には、nullが入って結合される。
```sql
SELECT *
FROM users
LEFT OUTER JOIN companies
      ON users.company_id = companies.id;
```
|  id  |  name  |  company_id  |  name  |  name  |
| ---- | ---- | ---- | ---- | ---- |
|  1  | 高橋  |  1  |  1  |  パナ  |
|  2  | 山田  |  2  |  2  |  ソニ  |
|  3  | 美山  |  4  |  null  |  null  |


### RIGHT OUTER JOIN 右側のテーブルが軸
- 右側のテーブルを軸にデータを結合する。
```sql
SELECT *
FROM users
RIGHT OUTER JOIN companies
      ON users.company_id = companies.id;
```
|  id  |  name  |  company_id  |  name  |  name  |
| ---- | ---- | ---- | ---- | ---- |
|  1  | 高橋  |  1  |  1  |  パナ  |
|  2  | 山田  |  2  |  2  |  ソニ  |
|  null  | null  |  null  |  3  |  日立山  |





