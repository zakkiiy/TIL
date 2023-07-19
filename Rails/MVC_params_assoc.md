#### MVC_paramsの続きで、テーブルのassociationについて

#### 1. Taskモデルの作成
- rails g model Task title:string completed:boolean todo:references
- rails db:migrate

#### 2. todo.rb task.rbにアソシエーションを記載する
- todo.rb
```ruby
has_many :tasks, dependent: :destroy
```
- task.rb
```ruby
belongs_to :todo
```

### TodoモデルとTaskモデルの関連性について。
#### 1. has_many 関連付け todoは複数のTaskを持てる
- １つのTodoレコードが、複数の関連付けられたTaskレコードを持つことを意味する。つまり、１つのTodoに対して、複数のタスクを持つことができる
- 例：Todoに「買い物リスト」というレコードがあり、それに関連づけられたTaskが「牛乳を買う」と「卵を買う」などのレコードが複数ある状態
#### dependent: :destroy オプションとは
- このオプションを指定することで、Todoが削除される際に、関連付けられたタスクも一緒に削除されるようになる。

#### 2. belongs_to 関連付け taskは一つのTodoに所属している
- １つのTaskレコードが、１つのTodoレコードに関連付けられている。つまりTaskは必ず一つのTodoに所属しているという関係。
- 具体的には、Taskに対して、todo_idという外部キーが保存され、それによってどのTodoに所属するか示される。
- todo_idはmodelを作成する際に、自分で作成する。todo:references

#### 3. Takskコントローラの作成
- rails generate controller Tasks

#### 4. ルーティングの設定
- 特定の todo に関連する task を操作するためのルートが作成されます。
```ruby
resources :todos do
  resources :tasks, only: [:new, :create, :edit, :update, :destroy]
end
```

#### 5. タスク作成の流れ
- index.html.erb
- 特定のIDの todo リストに新しいタスクを作成するためのリンクを生成する
```ruby
<%= link_to "タスクを追加", new_todo_task_path(todo) %>
# ルーティング
new_todo_task_path	GET	/todos/:todo_id/tasks/new(.:format)	 tasks#new
# URL 例
http://localhost:3000/todos/8/tasks/new
```

- tasks_controllerのnewアクション
```ruby
  before_action :set_todo

  def new
    @task = Task.new
  end

  private

  def set_todo
    @todo = Todo.find(params[:todo_id])
  end
```
- @todo = Todo.find(params[:todo_id])について疑問に思ったこと。
- なぜ:todo_idでわざわざ外部キーを指定している？ 外部キーのidを使って、todoの値を格納している。
- @taskで、空のインスタンスを生成
- binding.pry
```ruby
    5: def new
    6:   binding.pry
 => 7:   @task = Task.new
    8: end

[1] pry(#<TasksController>)> @todo
=> #<Todo:0x0000000108f1a000
 id: 8,
 title: "cccaa",
 description: "cccaaaaaaaa",
 created_at: Mon, 17 Jul 2023 08:26:51.481988000 UTC +00:00,
 updated_at: Mon, 17 Jul 2023 11:13:53.131013000 UTC +00:00>
[2] pry(#<TasksController>)> 
```

- 理由は、HTTPリクエストで送られてくるデータがtodo_idのため、paramsで受け取って外部キーのidで検索を行っている。
```ruby
    14: def set_todo
    15: binding.pry
 => 16:   @todo = Todo.find(params[:todo_id])
    17: end

[4] pry(#<TasksController>)> params
=> #<ActionController::Parameters {"controller"=>"tasks", "action"=>"new", "todo_id"=>"8"} permitted: false>
```

- new.html.index
- @taskモデルの空のインスタンスに入力したタイトル等を属性付与する。
- todo_tasks_path(@todo)で、以下のルーティングになりcreateアクションが実行される。
```ruby
todo_tasks_path	POST	/todos/:todo_id/tasks(.:format)	tasks#create
```

- createアクションが実行される
```ruby
<%= form_with(model: @task, url: todo_tasks_path(@todo)) do |form| %>
  <div>
    <%= form.label :title %>
    <%= form.text_field :title %>
  </div>

  <div>
    <%= form.label :completed %>
    <%= form.text_field :completed %>
  </div>

  <%= form.submit "作成" %>
```

- createアクション
- paramsで、以下を受け取っている。
```ruby
=> #<ActionController::Parameters {"authenticity_token"=>"HwF1BIdPP9AiMEgj98xr_n-n-e1Svo3XR7cj3kyO6pDFWsP6SrvJzyab34K-muC90HT5lrf3Vbg08JoOKW3otg",
 "task"=>{"title"=>"テスト", "completed"=>"テス"}, "commit"=>"作成", "controller"=>"tasks",
"action"=>"create", "todo_id"=>"8"} permitted: false>
[2] pry(#<TasksController>)> 
```

- paramsにtodo_idの8も受け取る
- @todoインスタンスに紐づいたTakskのコレクションを表している。この関連付けられたタスクのコレクションを使って、新しいタスクのインスタンスを作成しています。
```ruby
  def create
    @task = @todo.tasks.new(task_params)
    if @task.save
      redirect_to todo_path(@todo), notice: "タスクが作成されました。"
    else
      render :new
    end
  end

  def task_params
    params.require(:task).permit(:title, :completed)
  end
```
- binding.pry
```ruby
     9: def create
    10:   binding.pry
    11:   @task = @todo.tasks.new(task_params)
 => 12:   if @task.save
    13:     redirect_to todo_path(@todo), notice: "タスクが作成されました。"
    14:   else
    15:     render :new
    16:   end
    17: end

[20] pry(#<TasksController>)> @task
=> #<Task:0x000000010713f468 id: nil, title: "テスト", completed: true, todo_id: 8, created_at: nil, updated_at: nil>
[21] pry(#<TasksController>)>
```
chatgpt
はい、その通りです！@task = @todo.tasks.new(task_params) は、ToDoリスト @todo に紐づいた新しいタスクのインスタンスを生成しつつ、フォームから送信されたデータ（task_params）を使ってそのタスクの属性に値を設定しています。
具体的にどういう流れで値が入れられるかを説明します：
@todo はToDoリストのインスタンスです。このToDoリストに紐づいた新しいタスクを作成するために、@todo.tasks を使ってタスクのコレクションを取得します。
@todo.tasks は、ToDoリスト @todo に関連付けられたタスクのコレクションを表します。これに対して .new(task_params) を呼び出すことで、新しいタスクのインスタンスが作成されます。
task_params はStrong Parametersとして、フォームから送信されたタスクのデータをフィルタリングしたものです。フォームに入力されたデータは params としてコントローラーに送信されますが、この中から必要なデータだけを抽出して使うことでセキュリティを向上させます。
@todo.tasks.new(task_params) は、新しいタスクのインスタンスを作成しつつ、task_params の中から必要なデータを使ってそのタスクの属性に値を設定します。たとえば、フォームから送信されたタイトルや完了フラグのデータが task_params に含まれている場合、それらのデータを使って新しいタスクのタイトルや完了フラグが設定されます。
最終的に、@task にはToDoリスト @todo に紐づいた新しいタスクのインスタンスが格納されます。そして、@task.save を呼び出すことで、このタスクがデータベースに保存されます。これにより、ToDoリストに紐付いた新しいタスクが作成されることになります。





