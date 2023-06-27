## 1.0 環境構築
__AWS Cloud9の環境を構築__  

__Ruby 3.1.2 インストール__
- __rvm__ を使用することで、様々なバージョンのRubyを指定して、インストールできる。
```ruby:rvm.rb
$ rvm get stable
$ rvm install 3.1.2
$ rvm --default use 3.1.2
$ ruby -v
```
__Railsインストール__
```ruby:rvm.rb
$ gem install rails -v 7.0.4
$ rails -v
# bundlerのバージョンを指定してインストール
$ gem install bundler -v 2.3.14
```

## 1.3 Hello Worldの表示
```ruby:rvm.rb
# 上で実行してインストールしたバージョンのbundleコマンドと互換性が保たれる。
$ rails _7.0.4_ new hello_app --skip-bundle
```

- gemで特定のバージョンを指定しない限りは、最新バージョンのgemがインストールされる。
```ruby:gem.rb
# 
gem "capybara", ">= 3.26"
# 5.0代で最新のバージョンがインストールされる。
gem "puma", "~> 5.0"
```

## 1.9 gemインストール
```ruby:gem.rb
$ bundle _2.3.14_ install
$ bundle _2.3.14_ lock --add-platform x86_64-linux
```

## 1.10 IDEへの接続許可
- config/environments/development.rb
```ruby:ide.rb
Rails.application.configure do
  # クラウドIDE への接続を許可する
  config.hosts.clear
end
```

## 1.3.3 MVC
- render html: "hello, world"
- routerとは  
  コントローラとブラウザの間に配置されブラウザからのリクエストをコントローラに振り分ける役割を持つ。
  ```ruby:routes.rb
  # ブラウザから「/」が実行されると、applicationコントローラのhelloアクションが実行される。
  root "application#hello"
  ```

## 1.4.1 Git
- cloud ideではデフォルトでGitが入っている
  ```ruby:git.rb
  # Git バージョンを上げる
  $ source <(curl -sL https://cdn.learnenough.com/upgrade_git)
  ```

