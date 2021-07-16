# Todo List
Go를 이용한 Todo List을 개발 프로젝트입니다.

## ✨ TODO:
- [X] `/todo`: 새로운 TODO 리스트 추가하기 / POST
    - [X] `/todo`: `index.html`를 통한 리스트 추가하기 / POST
- [X] `/todo`: TODO 전체 리스트 출력하기 / GET
    - [X] `/todo`: `index.html`를 통한 리스트 출력하기
        - [X] `index.html`: 각 항목마다 따로따로 표시 
        - [X] `index.html`: API JSON `id`를 기준으로 정렬화
- [X] `/todo/{id:[0-9]+}`: id 해당한 TODO 리스트 출력하기 / GET
- [X] `/todo{id:[0-9]+}`: TODO 리스트 지우기 / DELETE
    - [ ] `index.html`: TODO 항목 프론트엔드에서 삭제 가능하도록 제작
- [X] `/todo`: TODO 리스트 업데이트 / PUT | PUT 메소드를 처리하는 방법에 대해서 고민 중입니다.
    - [X] `/todo/{id:[0-9]+}`: Bool형을 통해서 TODO 항목 상태 업데이트하도록 제작 / GET

## 기능 
무엇이 더 효과적으로 개발할 수 있고 가독성 좋은 API를 만들어야 하는지에 대해서 고민하고 있습니다.  

### 새로운 TODO 리스트 추가하기
```json
// Request | POST http://localhost:3000/todo 
{
    "content": "Hello" 
}
// Respone
{
    "id": 1,
    "Content": "Hello",
    "Completed": false,
    "created_at": "2021-07-11T15:47:40.890708+09:00"
}
```

###  TODO 리스트 출력하기 / GET

```json
// Reqeust / GET http://localhost:3000/todo/1

// Respone
{
    "id": 1,
    "Content": "이것은 테스트~",
    "Completed": false,
    "created_at": "2021-07-11T22:14:31.615345+09:00"
}
```

todo에 추가된 항목들을 다 보여줍니다.
```json
// Reqeust / GET http://localhost:3000/todo

// Respone
[
    {
        "id": 1,
        "Content": "이것은 테스트~",
        "Completed": false,
        "created_at": "2021-07-12T18:57:16.710946+09:00"
    },
    {
        "id": 2,
        "Content": "이것은 테스트~",
        "Completed": false,
        "created_at": "2021-07-12T18:57:17.115232+09:00"
    },
    {
        "id": 3,
        "Content": "이것은 테스트~",
        "Completed": false,
        "created_at": "2021-07-12T18:57:17.899493+09:00"
    }
]
```

## 오류를 해결하기 위한 방법
### 많은 분들께서 도와 주셨습니다
    - [코딩냄비 멤버분들](https://github.com/codingpot)
    - [@snowmerak](https://github.com/snowmerak)
    
### 오류 해결, 좋은 프로젝트를 만들기 위해서 읽은 자료들
- [Go로 TodoList 만들기(1)](https://velog.io/@soosungp33/Go%EB%A1%9C-TodoList-%EB%A7%8C%EB%93%A4%EA%B8%B01)
- [예제로 배우는 Go 프로그래밍 - Go 컬렉션 - Map](http://golang.site/go/article/14-Go-%EC%BB%AC%EB%A0%89%EC%85%98---Map)
- [GoLang에서 Map Iterate, add, update 하기](https://cpro95.tistory.com/155)