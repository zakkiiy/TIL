## each

item = ["apple", "remon", "banana"]

item.each do |i|
  p i
end

p item
#=> ["apple", "remon", "banana"]
puts item
# 綺麗な出力をしたい場合はputs
# apple
# remon
# banana

pp item 
# ["apple", "remon", "banana"]
print item
# ["apple", "remon", "banana"]

# ブロックの中から外の変数にはアクセス可能。外側からブロック内へのアクセスは不可能。
# 結果等を新しい配列にしたい場合は、定義してあげる。
item = ["apple", "remon", "banana"]
new_item = []
item.each do |i|
  new_item << i
end
p new_item

# eachメソッドはブロック内の処理を各要素に対して実行した後、元の配列（レシーバ）をそのまま返す

# map
# 新しく配列を作る。ちなみに処理毎の最後の結果で配列が作られる。
item = ["apple", "remon", "banana"]
new_item = item.map do |i|
  p i
  p "+++"
  p "lll"
end

p new_item

def foo(x:, y:, z:)
  p [x, y, z]
end

h = {x: 1, y: 2, z: 3}
foo(**h)

MSG = 42
MSG = 5
p MSG

x = true
x || exit(1)
puts("succeeded!")

ary = [ 1, 2, 3, 4, 5 ]
p ary.filter { |i| i.odd? }

begin
  puts 1 + "2"
rescue
  puts "Error ."
rescue TypeError
  puts "Type Error"
ensure
  puts " ensure"  
end

begin 
  puts 10 / 0
rescue
  p "0では割れません"
end

# rescue => e でエラー内容を eに 格納することができる
begin
  10 / 0
rescue => e
  puts e
end

begin
  puts aaa
rescue => e
  
  puts e
end

num = 10 / 2
raise
p num + 2

