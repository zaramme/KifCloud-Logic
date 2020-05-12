# KifCloud-Logic
- 将棋の盤面をコード上で再現し、エンコーディングするためのライブラリです。

- 局面および棋譜のエンコーディングロジック
  - ある盤面の状態をbase64での文字列に可逆エンコード・デコード可能。
    - 単純計算で80文字程度のデータ量が必要なところを、出現可能性の低い局面パターンを可変長領域で対応することで、
    - 平均で４０文字程度まで削減。可逆コードなので、APIのパラメータなどで使用した場合、
    - データベースレスで局面や棋譜の再現が可能。また出力コードをハッシュとして用いることで、
    - KeyValueスキーマを用いた高速な局面情報への高速アクセスや、コード上の算術演算を用いた類似局面の検索などが可能に。
- クライアントサイドで操作可能な将棋盤インターフェース,jsBoardとの連携が可能。
  - JavaScriptのみで、駒を将棋のルールに従って操作可能。
  - またブラウザ上で操作した盤面の情報はAjaxを用いてサーバー通信され、動的に局面に関する情報を取得・表示可能。
- DynamoDBを用いた局面・棋譜情報データベースの設計
  - 前述の局面エンコーディングを活用することで、局面と、その連続である棋譜をシームレスに管理、取得可能に。