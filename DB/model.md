作間さんのmodel is 何をもとに、改めてロボらんてくんに質問してみた。

## 1
モデルについて質問します。  
@board = Board.newがあるとします。  
Board.newでBoardモデルのオブジェクトを生成する。   
これでBoardモデルのインスタンスが使用可能になる。  
このインスタンスはORマッパーでDBのレコードと1対1でマッピングされている。  
@boardに代入することをインスタンス化という。  
この認識であっていますか？間違いがあれば教えてください。  
ちなみにインスタンスの塊がコレクションなのでしょうか？  
コレクションでも実際にはコレクションの中のインスタンスがレコードとorマッパーで1対1でマッピングされているのですか？  

## 
認識はだいたい合ってるよ。ただ、一点だけ補足すると、Board.newで生成したインスタンスは、まだDBのレコードとはマッピングされてないんだ。  
DBに保存（@board.save）された時点で、そのインスタンスとDBのレコードが1対1でマッピングされるんだよ。  
そして、コレクションについてだけど、そうだね、インスタンスの塊をコレクションと呼ぶことがあるよ。  
たとえば、Board.allで取得したものは、Boardのインスタンスのコレクションだね。  
そしてその各インスタンスがDBの各レコードとマッピングされているというわけだよ。  

## 2
例えばeditとupdateの場合は、インスタンスとレコードは元々マッピングがされているということですよね？  
いまいちピンときてないのですが、modelのインスタンス側は、DBのデータの件数分インスタンスを持っているということですか？  
厳密には違うとは思いますが、DB側のレコードを同じ形になっているということでしょうか？  

## 
そうだね、editやupdateの場合は、既存のレコードとそのレコードに対応するインスタンスが既にマッピングされているんだ。  
そして、モデルのインスタンスは、DBのデータの件数分存在するわけではないよ。
モデルのインスタンスは、アプリケーション側で必要なタイミングで生成されるもので、全てのレコードに対応するインスタンスが常に存在するわけではないんだ。  
たとえば、Board.find(1)とすると、idが1のレコードに対応するBoardのインスタンスが生成されるけど、その他のレコードに対応するインスタンスは生成されないよ。  
だから、DB側のレコードをモデル側が全部持っているわけではなく、必要な時に必要なレコードに対応するインスタンスを生成して、それを使って操作を行うんだよ。  
