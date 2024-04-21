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

- `each_char`
レシーバのそれぞれの文字をyieldする。文字は1文字の文字列で表現される。
yieldするとは？　ブロック内にあるブロック変数に値を渡すことを指すらしい。(今回の場合？)
```ruby
s = "abcde"
p s.each_char.map { |i|
  i * 2
}
# => ["aa", "bb", "cc", "dd", "ee"]
```

- `chars`
charsはレシーバのすべての文字列を配列で返す。
```ruby
p "cocoa".chars
["c", "o", "c", "o", "a"]
=> ["c", "o", "c", "o", "a"]
```

- `tally`
```ruby
 p "cocoa".chars.tally
{"c"=>2, "o"=>2, "a"=>1}
=> {"c"=>2, "o"=>2, "a"=>1}
```

- `shift` `pop` `push`
shiftは配列の最初の要素を削除し、その要素を返します。（先入れ先出し）
popは配列の最後の要素を削除し、その要素を返します。（後入れ先出し）
pushは配列の末尾に指定された要素を追加します。

- `select`,`filter`
Array#selectとArray#filterは与えられたブロックが真の値を返す要素の配列を返します。

- `reject`
各要素に対してブロックを評価し、その値が偽であった要素を集めた新しい配列を返します。条件を反転させた select です。

- `slice`
```ruby
a = [120, 40, 20, 80, 160, 60, 180]
a.slice(1,-1)
=> nil
irb(main):199> 
irb(main):200> a.slice(1..-1)
=> [40, 20, 80, 160, 60, 180]
```

要素を順番にブロックに渡して評価し、その結果が真になった要素をすべて削除します。 delete_if は常に self を返しますが、reject! は要素が 1 つ以上削除されれば self を、 1 つも削除されなければ nil を返します。
ブロックが与えられなかった場合は、自身と reject! から生成した Enumerator オブジェクトを返します。返された Enumerator オブジェクトの each メソッドには、もとの配列に対して副作用があることに注意してください。
```ruby
a = [0, 1, 2, 3, 4, 5]
a.delete_if{|x| x % 2 == 0}
p a #=> [1, 3, 5]

a = [0, 1, 2, 3, 4, 5]
e = a.reject!
e.each{|i| i % 2 == 0}
p a                    #=> [1, 3, 5]  もとの配列から削除されていることに注意。
```

- `concat`
```ruby
p ["apple", "banana"].concat ["banana", "carrot"]
["apple", "banana", "banana", "carrot"]
=> ["apple", "banana", "banana", "carrot"]

 p ["apple", "banana"] | ["banana", "carrot"]
["apple", "banana", "carrot"]
=> ["apple", "banana", "carrot"]

irb(main):220> p ["apple", "banana"] || ["banana", "carrot"]
["apple", "banana"]
=> ["apple", "banana"]
irb(main):221> 

irb(main):222> p ["apple", "banana"] & ["banana", "carrot"]
["banana"]
=> ["banana"]
```
