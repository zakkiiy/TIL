## params関連について、まとめる

#### 1. ルーティング
```ruby
Rails.application.routes.draw do
  resources :todos
  root 'todos#index'
end
```

- localhost:3000/を実行すると、root 'todos#index'のルーティングが実行される
- ルートURL(/)またはroot_path(ルーティングヘルパーメソッド)にGETリクエストが送信されると、todosコントローラのindexアクションが呼び出される。
- 要するに、アプリケーションのルートURLにアクセスされた際には、todos コントローラーの index アクションが実行され、Todoリソースの一覧ページが表示されるという意味。
```ruby
# routes
root_path	GET	/	todos#index
```

#### 2. コントローラのアクションが呼び出される
- Todo.allメソッドは、todosテーブルの全てのレコードをDBから取得して、それぞれをTodoモデルのインスタンスとして返す。
- このインスタンスはDBのレコードに対応しており対となっている。
- データベースのテーブルとRailsのモデルのオブジェクトをマッピングし、データベース操作をオブジェクト指向の方法で行う仕組み
- これらのインスタンスを@todosインスタンス変数に格納することで、viewファイルでそれらのレコードにアクセスできるようになる。
- Todo.allの場合は、複数のインスタンスが返ってくるためコレクションオブジェクト（配列）という。複数のインスタンスの集合体的な。
- 配列であるため、viewでeachメソッドで、コレクションの要素を順番に処理することができる。
```ruby
def index
  @todos = Todo.all
end
```

#### 3. viewが呼び出される
- views/todos/index.html.erb
```ruby
<% @todos.each do |todo| %>
  <li>
    <%= link_to todo.title, todo_path(todo) %>
  </li>
<% end %>
```

- todo_path(todo)で、以下のルーティングが呼び出される
```ruby
todo_path	GET	/todos/:id(.:format)	todos#show
```
- todo_path(todo)は与えられたTodoモデルのインスタンス(レコード)から、そのレコードのIDを取得してURLを生成する。(/todos/1)


#### 4. todoコントローラのshowアクションが呼び出される
- app/controllers/todos_controller.rb
- todo_path(todo)からリクエストされてきたレコードをparamsで受け取って、Todo.findでidが合致するレコードをDBから取得して、@todoインスタンス変数に格納する。

```ruby
  def show
    @todo = Todo.find(params[:id])
  end
```

```ruby
    6: def show
    7:   binding.pry
 => 8:   @todo = Todo.find(params[:id])
    9: end

[1] pry(#<TodosController>)> params
=> #<ActionController::Parameters {"controller"=>"todos", "action"=>"show", "id"=>"2"} permitted: false>
```
- レコードがidだけをparamsに、送っている。他のデータを混ぜるとセキュリティ的にも問題がある？
- 一般的にリクエストには、レコード（オブジェクト）自体を直接送ることはありません。代わりに、レコードを識別するための一意のID（主キー）をリクエストに含めることが一般的です。
- RailsのようなWebフレームワークでは、リクエストのパラメータとしてIDを含めることが一般的です。これにより、アプリケーションがどのレコードに対する操作を行うかを識別することが容易になります。IDを使用することで、データベースから特定のレコードを素早く見つけることができます。


#### 5. showのビューで表示する。

```ruby
<h1><%= @todo.title %></h1>
<p><%= @todo.id %></p>
<p><%= @todo.description %></p>
```


