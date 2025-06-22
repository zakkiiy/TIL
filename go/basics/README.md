## Go Basics
## Goの基本構文、型、構造体、ポインタ等の学習メモ

2025/05

### [001_variable](./001_variable)
- 明示的な変数定義
- 暗黙的な変数定義
  - 再定義不可　など

### [002_type](./002_type)
- int
- float
- boole
- string
- byte
- 配列、slice
- interface, nil
- 型変換
  - 文字列から数値
  - 数値から文字列

### [003_const](./003_const)
- 定数
- itoa

### [004_operators](./004_operators)
なし

### [005_function](./005_function) 
- 関数
- 無名関数
- 関数を返す関数
- 関数を引数に取る関数
- 可変長引数

### [006_closure](./006_closure) 
- クロージャー
  - 値を保持
- ジェネレーター

### [007_control_structures](./007_control_structures)
- if
  - if
  - エラーハンドリング
- for
  - for
  - 配列でfor
  - range
  - slice
  - map
- switch
  - switch
  - 型switch
- defer
  - defer
  - ファイルのクローズ処理
- panic
  - panic
  - recover()
- goroutin
  - 並列処理
- init
  - 初期化処理

### [008_reference_type](./008_reference_type)
- slice
  - 可変長引数
  - make関数でslice生成
  - sliceの操作
    - make, append, len, cap
  - slice for
- map
  - map
  - map for
  - makeを使用したmap生成
  - 関数

### [009_pointers](./009_pointers)
- ポインタ型

### [010_struct](./010_struct)
- 構造体
  - ポインタ
- 構造体のメソッド
  - レシーバ
  - レシーバポインタ
- 構造体の埋め込み
- struct型コンストラクタ
- slice
- map