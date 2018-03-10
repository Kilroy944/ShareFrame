index = {
    init: function() {
        // Init
        asticode.loader.init();
        asticode.modaler.init();
        asticode.notifier.init();
    }
};


document.addEventListener('astilectron-ready', function() {
    astilectron.onMessage(function(message) {
        if (message.name === "display_picture") {
            document.getElementById("img1").setAttribute("src",message.payload);
            return {payload: "0"};
        }
    });
})