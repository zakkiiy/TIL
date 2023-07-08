## 基本設定
### 01.rails ganerateコマンドで不要なファイルを作成しないようにする。
- generateコマンドの設定は、cofig/application.rbを編集する。
例
```ruby
# cofig/application.rb
    config.generators do |g|
      g.skip_routes true
      g.helper false
      g.test_framework false
    end
```
