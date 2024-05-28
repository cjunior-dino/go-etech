$(document).ready(function() {
    $('#enqueueButton').click(function() {
        var val = $('#enqueueValue').val();
        $.get('/enqueue?val=' + val, function(data) {
            // Atualize a interface do usuário conforme necessário
        });
    });

    $('#dequeueButton').click(function() {
        $.get('/dequeue', function(data) {
            // Atualize a interface do usuário conforme necessário
        });
    });

    $('#frontButton').click(function() {
        $.get('/front', function(data) {
            $('#frontValue').val(data);
        });
    });

    $('#isEmptyButton').click(function() {
        $.get('/isempty', function(data) {
            $('#isEmptyValue').val(data);
        });
    });

    $('#displayButton').click(function() {
        $.get('/display', function(data) {
            $('#displayValue').val(data);
        });
    });
});