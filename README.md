# lisb

lisb는 빌드 옵션을 다루는 간단한 툴입니다.

## install

먼저 Release에서 각 환경에 맞는 바이너리를 받아서 실행 권한을 주고 환경 변수 상 실행 가능한 경로에 저장합니다.  
가급적이면 저장하면서 파일 이름을 `lis`만 남기기 바랍니다.

## setting

### init

자신이 작업하는 go 작업 폴더에 가서 `lis -init demo`을 실행합니다.  
각 질문에 성심성의껏 응답하면 해당 폴더에 `demo.json` 파일이 생성될 것입니다.

각 질문에 세세한 설정이 불가능한 항목이 있을 수 있습니다, gogc가 그렇습니다.  
이런 경우엔 `demo.json`을 열어서 직접 세세하게 조정할 수 있습니다.

```json
{
  "bin_path": "bin",
  "name": "lis",
  "target": {
    "darwin": [
      "amd64"
    ],
    "linux": [
      "amd64",
      "386"
    ],
    "windows": [
      "amd64",
      "386"
    ]
  },
  "gogc": 0,
  "to_plugin": false,
  "auto_run": false,
  "module": true
}

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

### auto_run

`AutoRun` 옵션의 경우엔 `target` 항목에서 본인의 환경에 대한 빌드가 선행되어야합니다.  
지금 빌드하는 환경의 빌드 옵션이 꺼져 있다면 실행할 수 없으므로 에러가 발생합니다.

### module

`Module` 옵션은 `go111module`의 on, off를 결정합니다.  
true라면 on이고 false라면 off로 동작하게 됩니다.

### to_plugin

`ToPlugin` 옵션은 플러그인으로 컴파일할지 결정합니다.

### gogc

`gogc` 옵션은 go의 gc 작동 메모리 기준을 나타냅니다.  
100이 기존 메모리의 1배, 150이 기존 메모리의 1.5배, 50이 기존 메모리의 0.5배 양이 추가로 힙에 할당되었을 때 gc가 동작합니다.

## compile

컴파일은 `lis -make demo`를 입력하면 지정한 옵션대로 컴파일 됩니다.

## License

MPL

## import

import : https://github.com/AlecAivazis/survey