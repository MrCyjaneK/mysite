(function() {
    $('img[src*="thumb"]').parent('a').click(function(e) {
        e.preventDefault();
        $(this).ekkoLightbox();
    });
})();
