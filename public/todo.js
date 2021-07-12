'use strict';

$("#form").keypress(function(e) {
    if (e.keyCode === 13) {
        AddTodoListSend();
    }
})

function AddTodoListSend() {
    const todo = document.getElementById('content').value

    var data = {
        "content": todo
    }

    $.ajax({
        url: '/todo',
        type: 'POST',
        contentType: 'application/json',
        data: JSON.stringify(data),
        success: function(data) {
            console.log("전송 완료!");
            location.reload();
        },
        error: function(xhr, status, err) {
            console.log(xhr, status, err);
        }
    });
    return false;
}
function GetTodoList() {
    const parent = document.querySelector('#items')
    const fragment = document.createDocumentFragment();

    $.ajax({
        url: '/todo',
        type: 'GET',
        contentType: 'application/json',
        success: function(todo) {
            const Obj = JSON.parse(JSON.stringify(todo))
            // 코드 정렬화 / id를 기준으로 함.
            Obj.sort(function(a, b) {
                if (a.id < b.id) return -1;
                if (a.id > b.id) return 1;
                return 0;
            });
            // 출력
            Obj.forEach((item) => {
                let li = document.createElement('li');
                li.innerHTML = `<p class="item-content">${item.Content}</p>`
                fragment.append(li)                    
            });
            parent.append(fragment)

        },
        error: function(err) {
            console.error("오류 발생", err)
        }
    })
}