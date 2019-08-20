# 参考
- https://qiita.com/suin/items/a0b7eb585eeb6bda4174


# Goをインストール

## インストール
Mac, Linux, Windowsにコンパイルできるように、クロスコンパイルオプションをつける

```
brew install go --cross-compile-common
```  

なぜかエラーが・・

## 普通にインストール
brew install go

## gopathを編集
```
emacs ~/.zshrc
```
- 以下を追加

bashrc
```
...

#Goのパス設定
export GOPATH=$HOME/go
export GOROOT="$(brew --prefix golang)/libexec"
export PATH=$PATH:$HOME/go/bin
```

## ディレクトリ作成
- `$HOME/go` にsrc,bin,pkgを作成

### bin
- 実行ファイルが格納されるディレクトリ
### pkg
- ビルドしたパッケージオブジェクトが格納されるディレクトリ
### src
- パッケージごとのソースコードを配置するディレクトリ
