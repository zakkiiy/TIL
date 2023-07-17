## MVC〜params関連について、まとめる

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
- todo.titleは表示するために必要。
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

#### 6. 新規入力フォーム

- 新規作成ボタンを押下する
```ruby
# index
<%= link_to "新規作成", new_todo_path %>
# ルーティング
new_todo_path	GET	/todos/new(.:format)	todos#new
```

- newにGetリクエストが送信されると newアクションが実行される。
- new アクションで Todo.new を使用してインスタンスを生成しているのは、新しいTodoレコードの作成に必要なデータをフォームにセットするためです。
- フォームを使ってユーザーに入力してもらったデータを @todo インスタンスに格納して、そのデータを使って新しいTodoレコードを作成します。
``` ruby
  def new
    @todo = Todo.new
  end
```

- 新規入力フォーム(new.html.erb)
- form_with(model: @todo, url: todos_path) のようなフォームヘルパーを使って、@todo インスタンスをフォームに関連付ける。
- 新規入力フォームで入力された値は、Todoモデルのインスタンスの属性に付与される
- @todo = Todo.new として新しい空の Todo インスタンスを作成することで、インスタンスは全ての属性がデフォルト値（通常は nil もしくは空文字）で初期化されます。
- この Todo モデルでは、title と description という属性を持っています。この場合、@todo = Todo.new によって新しいインスタンスが作成されると、@todo.title と @todo.description の値はどちらも nil となります。
- form_withに空のインスタンスを渡すと、Postリクエストが送信される。
```ruby
<%= form_with(model: @todo, url: todos_path) do |form| %>
  <div>
    <%= form.label :title %>
    <%= form.text_field :title %>
  </div>

  <div>
    <%= form.label :description %>
    <%= form.text_field :description %>
  </div>

  <%= form.submit "作成" %>
<% end %>
```

- binding.pry
```ruby
    10: def new
    11:   @todo = Todo.new
    12:   binding.pry
 => 13: end

[1] pry(#<TodosController>)> @todo
=> #<Todo:0x00000001075db198 id: nil, title: nil, description: nil, created_at: nil, updated_at: nil>
```

#### 7. 作成ボタン押下
#### ルーティング
```ruby
todos_path POST	/todos(.:format)	todos#create
```

#### コントローラのcreateアクション
- paramsで値を受け取り、保存が成功したら、DBに保存される。

```ruby
  def create
    @todo = Todo.new(todo_params)
    if @todo.save
      redirect_to todo_path(@todo)
    else
      render :new    
    end
  end
```

- binding.pry
```ruby
    15: def create
    16:   @todo = Todo.new(todo_params)
    17:   binding.pry
 => 18:   if @todo.save
    19:     redirect_to todo_path(@todo)
    20:   else
    21:     render :new    
    22:   end
    23: end

[1] pry(#<TodosController>)> @todo
=> #<Todo:0x00000001079a5be8 id: nil, title: "テスト", description: "todo_params後", created_at: nil, updated_at: nil>
[2] pry(#<TodosController>)>
```

- binding.pry
```ruby
    14: def create
    15:   @todo = Todo.new(todo_params)
    16:   if @todo.save
    17:     binding.pry
 => 18:     redirect_to todo_path(@todo)
    19:   else
    20:     render :new    
    21:   end
    22: end

[1] pry(#<TodosController>)> @todo
=> #<Todo:0x0000000115df6ea0
 id: 6,
 title: "aaa",
 description: "aaaaaaaaa",
 created_at: Mon, 17 Jul 2023 08:24:20.047079000 UTC +00:00,
 updated_at: Mon, 17 Jul 2023 08:24:20.047079000 UTC +00:00>
[2] pry(#<TodosController>)>
```

#### 編集を作成する
- こちらをクリックすると`http://localhost:3000/todos/1/edit`に遷移できる
```ruby
# index.html.erb
%= link_to "編集", edit_todo_path(todo) %>
# ルーティング
edit_todo_path	GET	/todos/:id/edit(.:format)	todos#edit
```

- コントローラのeditアクション
- paramsで受け取って合致するidをDBから取得して@todoにインスタンスを格納する
```ruby
  def edit
    @todo = Todo.find(params[:id])
  end
```

- binding.pry
```ruby
    23: def edit    
    24:   @todo = Todo.find(params[:id])
    25:   binding.pry
 => 26: end

[1] pry(#<TodosController>)> @todo
=> #<Todo:0x00000001144774e8
 id: 1,
 title: "aaaaaaa編集",
 description: "test編集",
 created_at: Sun, 16 Jul 2023 15:27:22.909039000 UTC +00:00,
 updated_at: Mon, 17 Jul 2023 08:57:34.346384000 UTC +00:00>
[2] pry(#<TodosController>)> 
```



- edit.html.erbにて表示される。

```ruby
<%= form_with(model: @todo, url: todo_path(@todo), method: :patch) do |form| %>
  <div>
    <%= form.label :title %>
    <%= form.text_field :title %>
  </div>

  <div>
    <%= form.label :description %>
    <%= form.text_field :description %>
  </div>

  <%= form.submit "更新" %>
<% end %>
```

- binding.pry
- @todo = Todo.find(params[:id])で編集前のデータを格納
```ruby

    27: def update
    28:   @todo = Todo.find(params[:id])
    29:   binding.pry
 => 30:   if @todo.update(todo_params)
    31:     redirect_to todo_path(@todo)
    32:   else
    33:     render :edit
    34:   end
    35: end
[2] pry(#<TodosController>)> @todo
=> #<Todo:0x0000000106e7aac0
 id: 1,
 title: "編集前",
 description: "編集前のデータ",
 created_at: Sun, 16 Jul 2023 15:27:22.909039000 UTC +00:00,
 updated_at: Mon, 17 Jul 2023 10:09:19.922283000 UTC +00:00>
[3] pry(#<TodosController>)>
```

- binding.pry
- if @todo.update(todo_params)で編集後のデータを格納
```ruby
    27: def update
    28:   @todo = Todo.find(params[:id])
    29:   if @todo.update(todo_params)
    30:     binding.pry
 => 31:     redirect_to todo_path(@todo)
    32:   else
    33:     render :edit
    34:   end
    35: end
[2] pry(#<TodosController>)> @todo
=> #<Todo:0x00000001069c58e0
 id: 1,
 title: "編集後の後",
 description: "編集後の後のデータ",
 created_at: Sun, 16 Jul 2023 15:27:22.909039000 UTC +00:00,
 updated_at: Mon, 17 Jul 2023 10:12:46.987879000 UTC +00:00>
[3] pry(#<TodosController>)>
```

#### 削除
`<%= link_to "削除", todo_path(todo), data: {turbo_method: :delete, turbo_confirm: "削除してもいいですか？"} %>`

```ruby
  def destroy
    @todo = Todo.find(params[:id])
    @todo.destroy
    redirect_to todos_path
  end
```

#### それぞれのアクションで、インスタンス変数に代入していることを忘れずに。



