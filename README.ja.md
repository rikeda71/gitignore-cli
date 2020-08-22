# gitignore-cli

.gitignoreを生成できるCLIツール

## installation

```sh
brew tap s14t284/tap
brew install gitignore-cli
```

## Usage

### 標準出力にgitignoreを出力する場合

  ```sh
  gitignore-cli [OS, IDE, プログラミング言語など]
  (ex. gitignore-cli go macOS vim) # go, macOS, vimに関わるgitの監視下に必要ないファイルのリストが出力されます
  ```

### `.gitignore`に出力する場合

  ```sh
  gitignore-cli -f [OS, IDE, プログラミング言語など]
  (ex. gitignore-cli -f go macOS vim) # go, macOS, vimに関わるgitの監視下に必要ないファイルのリストが.gitignoreに保存されます
  ```

## LICENCE

MIT
