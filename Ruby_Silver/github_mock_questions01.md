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

シンボルの配列を簡単に作るためのリテラル記法だ
```ruby
p %i(x1 x2 x3)
[:x1, :x2, :x3]
=> [:x1, :x2, :x3]
```

クラスのスーパークラスを明示的に指定しなかった場合、
`Objectクラスがスーパークラスになる`
```ruby
class aaa
end
```

Rubyのクラス階層において、Object クラスはほぼすべてのオブジェクトの基底クラスとなっています。このため、Object クラスにメソッドを追加すると、そのメソッドはほとんどすべてのオブジェクトで使用可能になります。
```ruby
class Object
  def moo
    puts "MOO!"
  end
end

"Cow".moo
```

```ruby
class Foo
  attr_reader :var
  def initialize
    @var = "apple"
  end
end

class Bar < Foo
  def initialize
    @var = "banana"
    super
  end
end

bar = Bar.new
puts bar.var
```

このコード例における動作を解説すると、Bar クラスが Foo クラスから継承されており、両方のクラスに initialize メソッドが定義されています。super キーワードは、サブクラス（この場合は Bar）のメソッドから親クラス（この場合は Foo）の同名のメソッドを呼び出すために使用されます。

コードの実行順序
Bar のインスタンスが生成されると、Bar の initialize メソッドが呼ばれます。
Bar の initialize メソッド内で、インスタンス変数 @var に "banana" が代入されます。
super が呼ばれると、Foo の initialize メソッドが実行されます。このメソッド内で、@var に "apple" が再代入されます。
Foo の initialize が完了した後、Bar の initialize も完了します。
結果
結果的に、Bar のインスタンスで @var を呼び出すと、"apple" が出力されます。これは、Foo の initialize メソッドで @var が "banana" から "apple" に上書きされたためです。

super の意味
super を使うと、現在のメソッドと同じ名前の親クラスのメソッドが呼び出されます。これにより、サブクラスで拡張または変更したい処理を行った上で、親クラスの元の処理も適用させることができます。ただし、このケースでは、super の呼び出しによってサブクラスで設定した値が親クラスのメソッドで上書きされてしまっています。

まとめ
この例では、super の使用方法と、サブクラスから親クラスのメソッドをどのように呼び出すかが示されています。super の位置によって親クラスのメソッドがいつ呼ばれるかを制御することが重要であり、変数の値がどのように影響を受けるかを理解する必要があります。

```ruby
r = "a".."e"
=> "a".."e"
irb(main):239> 
irb(main):240> p r.to_a
["a", "b", "c", "d", "e"]
=> ["a", "b", "c", "d", "e"]
irb(main):241> 
```

- `find`
ブロックの値が最初に真となる要素を返す。
detectという別名がある
```ruby
p [0,1,2,3,4,5].find {|x| x < 3}
0
=> 0
```

```ruby
p [1,16,8,4,2].sort.reverse
[16, 8, 4, 2, 1]
=> [16, 8, 4, 2, 1]
p [1,16,8,4,2].reverse.sort
[1, 2, 4, 8, 16]
=> [1, 2, 4, 8, 16]
irb(main):248> p [1,16,8,4,2].sort_by { |x| -x }
[16, 8, 4, 2, 1]
=> [16, 8, 4, 2, 1]
irb(main):249> p [1,16,8,4,2].sort_by { |x| x }
[1, 2, 4, 8, 16]
=> [1, 2, 4, 8, 16]
irb(main):250> 
```

