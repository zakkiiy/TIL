## Ruby 問題集まとめ
- `empty?`
レシーバが空の場合、true。値が入っていればfalseとなる。
```ruby
[].empty?
=> true
irb(main):039> {}.empty?
=> true
irb(main):040> "文字列".empty?
=> false
irb(main):041> "".empty?
=> true
```
int型等の数値だとエラー
```ruby
1.empty?
(irb):43:in `<main>': undefined method `empty?' for 1:Integer (NoMethodError)
```

