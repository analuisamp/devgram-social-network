$('#login').on('submit', login);

function login() {
    event.preventDefault();

    $.ajax({
        url: "/login",
        method: "POST",
        data: {
            email: $('#email').val(),
            password: $('#password').val(),
        }
    }).done(function() {
        window.location = "/home"
    }).fail(function(err) {
        console.log(err);
        alert("Invalid user or password")
    });
}