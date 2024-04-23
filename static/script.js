window.onload = function() {
    const form = document.getElementById('addItemForm');
    form.onsubmit = function(event) {
        // Prevent the default form submission
        event.preventDefault();

        // Simple client-side validation example
        if (!form.id.value || !form.name.value || !form.description.value || !form.quantity.value) {
            alert("Please fill in all fields.");
        } else {
            // Call addItem() if validation passes
            addItem();
        }
    };
};

function addItem() {
    // Get form data
    var id = document.getElementById('id').value;
    var name = document.getElementById('name').value;
    var description = document.getElementById('description').value;
    var quantity = document.getElementById('quantity').value;
    
    // Create JSON payload
    var data = JSON.stringify({
        id: id,
        name: name,
        description: description,
        quantity: parseInt(quantity, 10)  // Ensure quantity is sent as a number
    });
    
    // Send POST request with JSON payload
    fetch('/inventory', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: data
    }).then(response => {
        if (response.ok) {
            return response.json();
        }
        throw new Error('Network response was not ok.');
    }).then(data => {
        console.log(data);
        // Redirect to the inventory page to see the new item
        window.location.href = '/inventory';
    }).catch(error => {
        console.error('Error:', error);
    });
}

function deleteItem(id) {
    fetch('/inventory/' + id, {
        method: 'DELETE'
    }).then(response => {
        if (response.ok) {
            console.log("Item deleted");
            // Update the UI or reload the page
            location.reload();
        } else {
            console.error('Item not found');
        }
    }).catch(error => {
        console.error('Error:', error);
    });
}

function enableEdit(id) {
    // Change the text fields to editable input fields
    document.getElementById('name-' + id).innerHTML = '<input type="text" id="edit-name-' + id + '" value="' + document.getElementById('name-' + id).innerText + '">';
    document.getElementById('description-' + id).innerHTML = '<input type="text" id="edit-description-' + id + '" value="' + document.getElementById('description-' + id).innerText + '">';
    document.getElementById('quantity-' + id).innerHTML = '<input type="number" id="edit-quantity-' + id + '" value="' + document.getElementById('quantity-' + id).innerText + '">';
    // Change the "Edit" button to a "Save" button
    document.querySelector(`#item-${id} button[onclick^='enableEdit']`).outerHTML = '<button onclick="updateItem(\'' + id + '\')">Save</button>';
}

function updateItem(id) {
    // Get form data from inputs specific to the item we want to update
    var name = document.getElementById('edit-name-' + id).value;
    var description = document.getElementById('edit-description-' + id).value;
    var quantity = document.getElementById('edit-quantity-' + id).value;
    
    // Create JSON payload
    var data = JSON.stringify({
        id: id,
        name: name,
        description: description,
        quantity: parseInt(quantity, 10)  // Ensure quantity is sent as a number
    });
    
    // Send PUT request with JSON payload
    fetch('/inventory/' + id, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json',
        },
        body: data
    }).then(response => {
        if (response.ok) {
            return response.json();
        }
        throw new Error('Network response was not ok.');
    }).then(data => {
        console.log(data);
        // Optionally refresh the page or update the table row directly
        location.reload();  // Simplistic approach to see the changes
    }).catch(error => {
        console.error('Error:', error);
    });
}
