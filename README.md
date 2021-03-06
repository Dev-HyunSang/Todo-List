# Todo List
Go를 이용한 Todo List을 개발 프로젝트입니다.  
Go를 이용하여서 CRUD 및 기본적인 개념에 대해서 공부 겸 개발 프로젝트입니다.

> 프로젝트를 종료 하였습니다. 현재 백엔드에서의 CRUD에 해당하는 모든 프로세스는 개발 완료 하였습니다.  
프론트엔드 경우에는 아직 경험이 전무하여 추후에 업데이트 시킬 예정입니다.
## ✨ TODO:
### BackEnd
- [X] `/todo`: 새로운 TODO 리스트 추가하기 / POST
    - [X] `/todo`: `index.html`를 통한 리스트 추가하기 / POST
- [X] `/todo`: TODO 전체 리스트 출력하기 / GET    
- [X] `/todo/{id:[0-9]+}`: id 해당한 TODO 리스트 출력하기 / GET
- [X] `/todo{id:[0-9]+}`: TODO 리스트 지우기 / DELETE
- [X] `/todo`: TODO 리스트 업데이트 / PUT 
    - [X] `/todo/{id:[0-9]+}`: Bool형을 통해서 TODO 항목 상태 업데이트하도록 제작 / PUT
    - [X] `/todo/{id:[0-9]+}`: TODOLIST Content 수정할 수 있도록 제작 / PUT
### FrontEnd
- [X] `/todo`: `index.html`를 통한 리스트 출력하기
- [X] `index.html`: 각 항목마다 따로따로 표시 
- [X] `index.html`: API JSON `id`를 기준으로 정렬화
- [ ] `index.html`: TODO 항목 프론트엔드에서 삭제 가능하도록 제작
- [ ] `index.html`: 체크 박스를 이용하여서 Completed 항목에 대한 수정 /  완료되었는지 안 완료되었지
    
## 기능 
무엇이 더 효과적으로 개발할 수 있고 가독성 좋은 API를 만들어야 하는지에 대해서 고민하고 있습니다.  

|URL|Methods|설명|
|:---:|:---:|:---:|
| `/` | GET | `index.html`에 대해서 화면에 표시합니다.|
| `/todo` | GET| 추가된 TODO 항목 전체를 출력합니다.|
| `/todo` | POST | 새로운 TODO 항목을 추가할 수 있습니다. |
| `/todo/{id:[0-9]+}` | GET | TODO 항목의 id를 참고하여 해당 ID만 출력합니다.|
| `/todo/edit/{id:[0-9]+}` | GET | 기존 TODO 항목의 id를 참고하여 수정합니다.(현재 마음에 안 드는 부분이 있어서 수정 중...)|
| `/todo/{id:[0-9]+}` | DELETE | 기존 TODO 항목의 id를 참고하여서 삭제합니다. | 


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

### TODO 리스트 삭제하기 / DLETET
```json
// Reqeust / DLETET http://localhost:3000/todo/1

// Respone
DLETET TODO LIST ID:1
```

## 오류를 해결하기 위한 방법
### 많은 분들께서 도와 주셨습니다
- [코딩냄비 멤버분들](https://github.com/codingpot)
- [@snowmerak](https://github.com/snowmerak)
    
### 오류 해결, 좋은 프로젝트를 만들기 위해서 읽은 자료들
- [Go로 TodoList 만들기(1)](https://velog.io/@soosungp33/Go%EB%A1%9C-TodoList-%EB%A7%8C%EB%93%A4%EA%B8%B01)
- [예제로 배우는 Go 프로그래밍 - Go 컬렉션 - Map](http://golang.site/go/article/14-Go-%EC%BB%AC%EB%A0%89%EC%85%98---Map)
- [GoLang에서 Map Iterate, add, update 하기](https://cpro95.tistory.com/155)
