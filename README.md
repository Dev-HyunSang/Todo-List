# Todo List

## TODO:
- [X] `/todo`: 새로운 TODO 리스트 추가하기 / POST
    - [X] `/todo`: `index.html`를 통한 리스트 추가하기 / POST
- [ ] `/todo`: TODO 리스트 출력하기 / GET
- [ ] `/todo`: TODO 리스트 지우기 / PUT
- [ ] `/todo`: TODO 리스트 수정 / DELETE

## 새로운 TODO 리스트 추가하기
```json
// Request | POST http://localhost:3000/todo 
{
    "content": "Hello" 
}
// Respone
{
    "id": 0,
    "Content": "Hello",
    "Completed": false,
    "created_at": "2021-07-11T15:47:40.890708+09:00"
}
```