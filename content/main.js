(function() {
    $('img[src*="thumb"]').parent('a').click(function(e) {
        e.preventDefault();
        $(this).ekkoLightbox();
    });
})();


if (navigator.userAgent.indexOf("Chrome") == -1) {
    console.log("psst! please see https://stackoverflow.com/a/63789826.")
}