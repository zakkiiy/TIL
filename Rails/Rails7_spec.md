## rails7 仕様
#### status: :see_other
- link_toメソッドで、:deleteメソッドを指定した場合、以前だとHTTPのdeleteメソッドが使われていたが、Rails7だとTurboの影響で、status: :see_otherをつける必要がある。
```ruby
def destroy
  logout
  redirect_to root_path, status: :see_other
end
```

#### status: :unprocessable_entity
- Rails7系ではrenderメソッド実行時のflashメッセージやエラーメッセージはstatus: :unprocessable_entityをつけないとメッセージが表示されない
```ruby
flash.now[:danger] = t('users.create.failure')
render :new, status: :unprocessable_entity
```

