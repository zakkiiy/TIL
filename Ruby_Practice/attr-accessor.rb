# クラスの定義（設計図）
class Person
  # アクセサメソッドを使ってname, age, genderの読み書きを可能にする
  # アクセサメソッドを使用することで、外部からインスタンス変数を操作できる。
  attr_accessor :name, :age, :gender

  # インスタンス生成時に必ず1度実行される。
  def initialize(name, age, gender)
    @name = name
    @age = age
    @gender = gender
  end

  def info
    "#{@name}さんは#{@age}歳の#{@gender}です。"
  end
end

# インスタンスの生成
person = Person.new("山田", 24, "男性")

# インスタンスメソッドの実行（生成したインスタンスに対して）
info = person.info
p info  # => "山田さんは24歳の男性です。"
p person.name
p person.age
p person.gender
# アクセサメソッドを使用して属性を変更する
person.name = "田中"
person.age = 30
person.gender = "女性"

# 変更後の情報を表示
p person.info  # => "田中さんは30歳の女性です。"
