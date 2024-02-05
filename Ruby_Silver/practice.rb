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
p a.index("R", 1)

