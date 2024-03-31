require 'date'

# 3
# pop スタック（後入れ先出し法）最後に追加された要素が最初に取り出される。
# 再列の最後の要素を削除して、その要素をメソッドの返り値としてメソッドの返り値として得られる。
s = ["one", "tow", "three"]
a = s.pop
s.pop
s.unshift "zero" # 引数を配列の先頭に入れる
p a
p s
s.push "four"
s.unshift
p s

# キュー（先入先出法）

# 4 %Fで省略して書くことが可能。 先頭%Yであること要注意
p Date.today.strftime("%F") # 2024-02-03
p Date.today.strftime("%Y-%m-%d")

# 5
a = "Ruby"
b = " on Rails"
a.concat b # "Ruby on Rails" 結合する
a.reverse # 新しいオブジェクトを返す（元のオブジェクトは変更されない）
p a # "Ruby on Rails"
a.reverse! # 破壊的メソッド オブジェクト自体を返す（元のオブジェクトの変更がされる）
p a # "sliaR no ybuR"
p a.index("R", 1) # 

arr = [
  true.equal?(true), # Rubyにおいて true、false、および nil はそれぞれが一意の特別なオブジェクトである
  nil.eql?(NilClass), # nilはNilClassのインスタンスであり、NilClassはクラスであるため、false
  String.new.equal?(String.new), # 同じオブジェクト_idを持ってないためfalse
  1.equal?(1) # 数値も同じオブジェクトIDのためtrue
]

p arr.collect { |a| a ? 1 : 2 }.inject(:+) # 6


# 文字列の結合には concatや + を使用する。
# appendは配列に使用。

h = {a: 100, b:100}
h.each {|p|
  p p.class
  p p
}

# ハッシュにデフォルトの値を設定して、ないキーを呼び出すとデフォルトの値が出力されるようになる。
h = Hash.new("default value")
h[:a]
h[:b] = 100
p h

v1 = 1 - 1 == 0
v2 = v1 || raise(RuntimeError)
puts v2 && false


a1 = [1,2,false]
a2 = [4,2,3]
p a1 && a2
# [4,2,3]

# klassだけだと、何を参照しているかわからないため、エラーとなる。
klass = Class.new
hash = {klass => 100}
puts hash[klass] # 100

klass = Class.new
hash = {}
hash.store(klass, 100)
puts hash[klass] # 100

klass = Class.new
hash = {}
hash.store(:klass, 100)
puts hash[:klass]

klass = Class.new
hash = Hash[klass, 100]
puts hash[klass]

hoge = *"a"
puts hoge.class

def foo(n)
  n ** n
end

puts foo (2) * 2

(1..10).to_a.each.
reverse_each.
each do |i|
  puts i
end

$val = 0

p 0.upto(5).select(&:odd?)
p 0.upto(5).select{ |n| n.odd?}

str = "1;2;3;4"
p str.split(";")

a = [1]
a[5] = 10
a.compact
p a

str = "aaabbcccddd"
p str.scan("c")

a1 = "abc"
a2 = 'abc'

print a1.equal? a2
print a1.eql? a2

io = File.open('list.txt')
while not io.eof?
  io.readlines
  io.seek(0, IO::SEEK_CUR)
  p io.readlines
end

arr = [5, 3, 8, 1, 4, 2, 6, 9, 0, 7]
arr.sort!{ |a, b | b <=> a }.reverse!
p arr


a = "Ruby"
b = " on Rails"
a.concat b
a.reverse
p a.index("R", 1)