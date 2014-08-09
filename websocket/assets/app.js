var App = function() {

    var webSocket = new WebSocket('ws://localhost:3000/ws');

    webSocket.onopen = function() {

        webSocket.onmessage = function(response) {

            console.log(response);
            var deleteObj = JSON.parse(response.data);
            console.log(deleteObj);

            if(_.isEqual(deleteObj.type, 'delete')) {
                console.log('.list-group-item[data-id="' + deleteObj.data + '"]');
                var $el = $('.list-group-item[data-id="' + deleteObj.data + '"]');
                $el.remove();
            } else {
                var message = $('<li class="list-group-item">').text(response.data);

                $('#messages').prepend(message);
                $('#inputMessage').val('');
            }
        }

        $('#form').on('submit', function(e) {

            e.preventDefault();

            if(! _.isEmpty($('#inputMessage').val())) {
                webSocket.send(JSON.stringify({type: 'new', data: $('#inputMessage').val()}));
            }

        });

        var deleteButton = $('.delete-entry');

        deleteButton.on('click', function(e) {

            e.stopPropagation();

            var data = $(this).data('id');
            var url = '/ws/' + data;

            var promise = $.ajax({
                "url" : url,
                "data" : data,
                "type" : "DELETE"
            });

            promise.done(function() {
                console.log(data);
                webSocket.send(JSON.stringify({type: 'delete', data: data}));
            });

        });

    }

}

$(function() {
    new App();
});
