$('#register-form').on('submit', createUser);

function createUser(event) {
    event.preventDefault();
    console.log("inside function")

    if ($('#password').val() != $('#confirm-password').val()) {
        alert("passwords do not match");
        return
    }

    $.ajax({
        url: "/users",
        method: "POST",
        data: {
            name: $('#name').val(),
            nick: $('#nick').val(),
            email: $('#email').val(),
            password: $('#password').val(),
        }
    }).done(function() {
        alert("User succesfully registered")
    }).fail(function(err) {
        console.log(err);
        alert("Error when registering user")
    });
}