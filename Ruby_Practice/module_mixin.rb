
# モジュールはメソッドや定数をまとめるために使用される。
# モジュールは直接インスタンス化できない
# Greeterモジュールを定義

module Greeter
  def greet
    char = "Hello #{self.class}"
    p char+"module"
  end
end

# Personクラスを定義して、Greeterモジュールをミックスイン
class Person
  include Greeter
end

person = Person.new
puts person.greet