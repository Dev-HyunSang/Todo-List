# Todo List
Go를 이용한 Todo List을 개발 프로젝트입니다.

## ✨ TODO:
- [X] `/todo`: 새로운 TODO 리스트 추가하기 / POST
    - [X] `/todo`: `index.html`를 통한 리스트 추가하기 / POST
- [X] `/todo`: TODO 전체 리스트 출력하기 / GET
    - [ ] `/todo`: `index.html`를 통한 리스트 출력하기
- [X] `/todo/{id:[0-9]+}`: id 해당한 TODO 리스트 출력하기 / GET
- [X] `/todo{id:[0-9]+}`: TODO 리스트 지우기 / DELETE
- [ ] `/todo`: TODO 리스트 업데이트 / PUT


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
