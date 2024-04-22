window.onload = function() {
    const form = document.querySelector("form");
    form.onsubmit = function(event) {
        // Simple client-side validation example
        if (!form.id.value || !form.name.value) {
            alert("Please fill in all fields.");
            event.preventDefault(); // Prevent form submission
        }
    }
}
