function login(username, password) {
    $.ajax({
        type: "post",
        contentType: "application/json",
        url: "/login",
        data: {
            'username': username,
            'password': password
        },
        complete: function(data, status) {
            return data;
        }
    });
}