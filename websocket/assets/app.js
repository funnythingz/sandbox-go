var App = function() {

    var webSocket = new WebSocket('ws://localhost:3000/ws');

    webSocket.onopen = function() {

        webSocket.onmessage = function(response) {

            console.log(response);

            var message = $('<li class="list-group-item">').text(response.data);

            $('#messages').prepend(message);
            $('#inputMessage').val('');
        }

        $('#form').on('submit', function(e) {

            e.preventDefault();

            if(! _.isEmpty($('#inputMessage').val())) {
                webSocket.send($('#inputMessage').val());
            }

        });

    }

}

$(function() {
    new App();
});
