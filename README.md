# tucker-go-project

Tucker의 Go 언어 프로그래밍 2판의 프로젝트를 연습하기 위한 repo

## project 1. numberquiz

1부터 1000까지의 자연수 중에서 랜덤으로 하나를 선택하고 사용자에게 숫자를 입력받아 맞추는 프로그램
사용한 내용

- fmt.Scanf()를 사용하여 사용자의 입력을 받아옴
- fmt.Scanf()의 에러 처리를 통해 사용자의 입력에 대한 예외 처리 (정수가 아닌 경우, 범위를 벗어난 경우)
  - 이 경우 bufio.NewReader()를 사용하여 기존 입력을 초기화하고 사용자의 입력을 다시 받음
- 무한루프를 사용하여 사용자가 맞출 때까지 계속 입력을 받음
- 사용자가 정답을 맞추면 시도한 횟수를 출력하고 프로그램 종료
