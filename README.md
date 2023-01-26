# lis

lis는 빌드 옵션을 다루는 간단한 툴입니다.

## install

고의 install 명령어를 통해 자동으로 설치합니다.

`go install github.com/snowmerak/lis@latest`

## setting

### init

자신이 작업하는 go 작업 폴더에 가서 `lis -init demo`을 실행합니다.  
이 때, `build`를 선택하면 빌드를 위한 config 파일이 생기고, `test`나 `bench`를 선택하면 테스트를 위한 config 파일이 생깁니다.  
이 단계에서는 `build`를 선택했다고 가정합니다.  
각 질문에 입력한 대로 해당 폴더에 `demo.yaml` 파일이 생성될 것입니다.

각 질문에 세세한 설정이 불가능한 항목이 있을 수 있습니다, gogc가 그렇습니다.  
이런 경우엔 `demo.yaml`을 열어서 직접 세세하게 조정할 수 있습니다.

```yaml
bin_path: bin
name: demo
target:
  darwin:
  - amd64
  - arm64
  linux:
  - amd64
  - "386"
  - arm64
  windows:
  - amd64
  - "386"
gogc: 150
to_plugin: false
module: true
etc: []
```

### target

특히 `target` 항목에서 빌드할 OS와 아키텍처를 지정할 수 있습니다.  
이 부분은 워낙 경우의 수가 많아서 따로 CLI 입력으로 받지 않았습니다.  
이 경우의 수는 아래 리스트를 참고해주세요.

```
$GOOS	$GOARCH
aix	ppc64
android	386
android	amd64
android	arm
android	arm64
darwin	amd64
darwin	arm64
dragonfly	amd64
freebsd	386
freebsd	amd64
freebsd	arm
illumos	amd64
ios	arm64
js	wasm
linux	386
linux	amd64
linux	arm
linux	arm64
linux	ppc64
linux	ppc64le
linux	mips
linux	mipsle
linux	mips64
linux	mips64le
linux	riscv64
linux	s390x
netbsd	386
netbsd	amd64
netbsd	arm
openbsd	386
openbsd	amd64
openbsd	arm
openbsd	arm64
plan9	386
plan9	amd64
plan9	arm
solaris	amd64
windows	386
windows	amd64
```

맵을 보시면 아시겠지만 `$GOOS`에 `$GOARCH`을 넣는 방식으로 작성하시면 됩니다.  

### module

`Module` 옵션은 `go111module`의 on, off를 결정합니다.  
true라면 on이고 false라면 off로 동작하게 됩니다.

### to_plugin

`ToPlugin` 옵션은 플러그인으로 컴파일할지 결정합니다.

### gogc

`gogc` 옵션은 go의 gc 작동 메모리 기준을 나타냅니다.  
100이 기존 메모리의 1배, 150이 기존 메모리의 1.5배, 50이 기존 메모리의 0.5배 양이 추가로 힙에 할당되었을 때 gc가 동작합니다.

## compile

컴파일은 `lis -build demo`를 입력하면 지정한 옵션대로 컴파일 됩니다.

## test & benchmark

가장 위의 `init` 단계에서 `test`나 `bench`를 선택했다면, 다음과 같은 yaml 파일이 생성됩니다.

```yaml
targets: []
flags: []
```

`targets`는 테스트할 대상을 지정하고, `flags`는 테스트할 때 사용할 플래그를 지정합니다.

`targets`는 `target` 구조체의 리스트로 이루어집니다.
- `package`는 테스트할 패키지를 지정합니다. 메인 패키지로부터의 상대 경로를 입력하면 됩니다.
- `test` 혹은 `bench`는 문자열 리스트로 테스트 혹은 벤치마크할 함수들을 작성합니다.
- 양쪽 다 init 할 때 질문에 대한 응답으로 입력할 수 있습니다.

- `lis -test demo`를 입력하면 지정한 옵션대로 테스트를 실행합니다.
- `lis -bench demo`를 입력하면 지정한 옵션대로 벤치마크를 실행합니다.

## License

MPL
