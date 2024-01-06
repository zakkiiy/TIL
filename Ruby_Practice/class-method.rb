# クラスメソッドとはクラス自体に関連づけられたメソッド
# クラスを直接レシーバとして、メソッドを呼び出すことができる。
# クラス内で共通で使いたい情報をあらかじめ用意できる？
# インスタンスに依存しない値

class SampleClass
  # selfはクラス自身を指すメソッド（レシーバ）
  def self.sample_class
    puts "クラスメソッドの実行"
  end

  # selfブロック
  class << self
    def sample_class2
      puts "クラスメソッド2"
    end

    def sample_class3
      puts "クラスメソッド3"
    end

    def sample_class4
      puts "クラスメソッド4"
    end
  end

end

SampleClass.sample_class
SampleClass.sample_class2
SampleClass.sample_class3
SampleClass.sample_class4


class Exam
  def initialize(scores)
    @scores = scores
  end

  def total_score
    @scores.sum
  end

  def passing_total?
    p total_score
    if total_score >= self.class.passing_total
      "合格"
    else
      "不合格"
    end
  end

  def self.passing_total
    400
  end
end

exam = Exam.new([100,70,90,100,60])
result = exam.passing_total?
puts(result)

## クラスメソッドはクラス全体に関わる変更、参照する場合（動的に変わらない？）
## インスタンスメソッドは個別のインスタンスの情報を変更、参照する場合